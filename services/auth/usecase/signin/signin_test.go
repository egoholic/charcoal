package signin_test

import (
	"context"
	"net"

	"github.com/egoholic/charcoal/services/auth/account"
	acR "github.com/egoholic/charcoal/services/auth/account/repo"
	sR "github.com/egoholic/charcoal/services/auth/session/repo"
	sDrv "github.com/egoholic/charcoal/services/auth/storage/mongodb/session"

	acDrv "github.com/egoholic/charcoal/services/auth/storage/mongodb/account"
	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/egoholic/charcoal/services/auth/account/pwd"
	"github.com/egoholic/charcoal/services/auth/config"
	. "github.com/egoholic/charcoal/services/auth/usecase/signin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	a             *account.Account
	acRepo        *acR.Repo
	sRepo         *sR.Repo
	client        *mongo.Client
	name1         = "Chuck Norris"
	password1     = "12345678"
	wrongPassword = "wrongpassword"
)

var _ = Describe("Signin Usecase", func() {
	Describe("Signin()", func() {
		BeforeEach(func() {
			a = account.New(name1, pwd.New(password1))
			acRepo = acR.New()
			sRepo = sR.New()
			client, _ = mongo.NewClient(config.MongoDBConnectionString())
			client.Connect(context.TODO())

		})

		Context("when correct name and password", func() {
			It("signs in", func() {
				insertAccount := acRepo.NewInserter(acDrv.NewInserter(context.TODO(), client))
				insertAccount(a)
				byPKFinder := acRepo.NewByPKFinder(acDrv.NewByPKFinder(context.TODO(), client))
				sessionInserter := sRepo.NewInserter(sDrv.NewInserter(context.TODO(), client))
				session, _ := Signin(name1, password1, net.ParseIP("127.0.0.1"), byPKFinder, sessionInserter)
				Expect(session).NotTo(BeNil())
				Expect(session.Account()).To(Equal(a))
			})
		})

		Context("when wrong password", func() {
			It("fails", func() {
				insertAccount := acRepo.NewInserter(acDrv.NewInserter(context.TODO(), client))
				insertAccount(a)
				byPKFinder := acRepo.NewByPKFinder(acDrv.NewByPKFinder(context.TODO(), client))
				sessionInserter := sRepo.NewInserter(sDrv.NewInserter(context.TODO(), client))
				session, err := Signin(name1, wrongPassword, net.ParseIP("127.0.0.1"), byPKFinder, sessionInserter)
				Expect(session).To(BeNil())
				Expect(err).To(BeNil())
			})
		})

		Context("when account doesn't exist ", func() {
			It("fails", func() {
				byPKFinder := acRepo.NewByPKFinder(acDrv.NewByPKFinder(context.TODO(), client))
				sessionInserter := sRepo.NewInserter(sDrv.NewInserter(context.TODO(), client))
				session, err := Signin(name1, password1, net.ParseIP("127.0.0.1"), byPKFinder, sessionInserter)
				Expect(session).To(BeNil())
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
