package signup_test

import (
	"context"
	"net"
	"time"

	"github.com/egoholic/charcoal/services/auth/account"
	"github.com/egoholic/charcoal/services/auth/account/pwd"
	accountsRepo "github.com/egoholic/charcoal/services/auth/account/repo"
	"github.com/egoholic/charcoal/services/auth/session"
	sessionsRepo "github.com/egoholic/charcoal/services/auth/session/repo"

	"github.com/egoholic/charcoal/services/auth/config"
	accountsAdapter "github.com/egoholic/charcoal/services/auth/storage/mongodb/account"
	sessionsAdapter "github.com/egoholic/charcoal/services/auth/storage/mongodb/session"

	"github.com/mongodb/mongo-go-driver/mongo"

	. "github.com/egoholic/charcoal/services/auth/usecase/signup"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	aR        = accountsRepo.New()
	sR        = sessionsRepo.New()
	client    *mongo.Client
	name1     = "Donald Trump"
	password1 = "12345678"
	ip1       = net.ParseIP("127.0.0.1")
	now       = time.Now()
)

var _ = Describe("Signup Usecase", func() {
	BeforeEach(func() {
		client, _ = mongo.NewClient(config.MongoDBConnectionString())
		client.Connect(context.TODO())
	})

	Describe("Signup()", func() {
		Context("when valid arguments", func() {
			It("creates and returns an account", func() {
				findAccount := aR.NewByPKFinder(accountsAdapter.NewByPKFinder(context.TODO(), client))
				insertAccount := aR.NewInserter(accountsAdapter.NewInserter(context.TODO(), client))
				insertSession := sR.NewInserter(sessionsAdapter.NewInserter(context.TODO(), client))

				a, s, err := Signup(name1, password1, ip1, findAccount, insertAccount, insertSession)
				Expect(err).To(BeNil())
				Expect(a).To(BeAssignableToTypeOf(&account.Account{}))
				Expect(s).To(BeAssignableToTypeOf(&session.Session{}))
				Expect(a.PK()).To(Equal(name1))
				Expect(s.Account()).To(Equal(a))
			})
		})

		Context("when invalid arguments", func() {
			It("changes nothing and fails", func() {

			})
		})

		Context("when account already exists", func() {
			It("changes nothing and fails", func() {
				findAccount := aR.NewByPKFinder(accountsAdapter.NewByPKFinder(context.TODO(), client))
				insertAccount := aR.NewInserter(accountsAdapter.NewInserter(context.TODO(), client))
				insertSession := sR.NewInserter(sessionsAdapter.NewInserter(context.TODO(), client))
				ac := account.New(name1, pwd.New(password1))
				insertAccount(ac)
				a, s, rerr := Signup(name1, password1, ip1, findAccount, insertAccount, insertSession)
				Expect(a).To(BeNil())
				Expect(s).To(BeNil())
				Expect(rerr.Error()).To(Equal("!Err: Can't sign up.\n\tReason: Account `Donald Trump` already exists."))
			})
		})
	})
})
