package repo_test

import (
	"context"
	"net"
	"time"

	"github.com/egoholic/charcoal/services/auth/account"
	"github.com/egoholic/charcoal/services/auth/account/pwd"
	"github.com/egoholic/charcoal/services/auth/config"
	"github.com/egoholic/charcoal/services/auth/session"
	adapter "github.com/egoholic/charcoal/services/auth/storage/mongodb/session"

	. "github.com/egoholic/charcoal/services/auth/session/repo"
	"github.com/egoholic/charcoal/services/auth/session/token"
	"github.com/mongodb/mongo-go-driver/mongo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	repo     *Repo
	client   *mongo.Client
	account1 *account.Account
	payload1 *session.Payload
	session1 *session.Session
)

var _ = Describe("Sessions Repository", func() {
	BeforeEach(func() {
		repo = New()
		client, _ = mongo.NewClient(config.MongoDBConnectionString())
		client.Connect(context.TODO())
		account1 = account.New("Joe Rogan", pwd.New("12345678"))
		payload1 = session.NewPayload(account1, token.New(), net.ParseIP("127.0.0.1"), time.Now().Add(time.Second*1000))
		session1 = session.New(payload1)
	})

	Context("creation", func() {
		Describe("New()", func() {
			It("returns repository", func() {
				Expect(New()).To(BeAssignableToTypeOf(&Repo{}))
			})
		})
	})

	Describe(".NewInserter()", func() {
		It("returns correct inserter", func() {
			insert := repo.NewInserter(adapter.NewInserter(context.TODO(), client))

			Expect(insert(session1)).To(BeNil())
		})
	})

	Describe(".NewByTokenFinder()", func() {
		It("returns correct finder", func() {
			find := repo.NewByTokenFinder(adapter.NewByPKFinder(context.TODO(), client))
			insert := repo.NewInserter(adapter.NewInserter(context.TODO(), client))
			insert(session1)
			r, _ := find(session1.Token())
			Expect(r).To(BeAssignableToTypeOf(&session.Session{}))
			Expect(r).NotTo(BeNil())
			Expect(r.Token()).To(Equal(session1.Token()))
		})
	})

	Describe(".NewByAccountFinder()", func() {
		It("returns correct finder", func() {
			find := repo.NewByAccountFinder(adapter.NewByAccountFinder(context.TODO(), client))
			insert := repo.NewInserter(adapter.NewInserter(context.TODO(), client))
			insert(session1)
			r, _ := find(session1.Token())
			Expect(r).To(BeAssignableToTypeOf(&session.Session{}))
			Expect(r).NotTo(BeNil())
			Expect(r.Account()).To(Equal(account1))
		})
	})
})
