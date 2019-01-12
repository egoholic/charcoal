package repo_test

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/egoholic/charcoal/services/auth/config"
	storage "github.com/egoholic/charcoal/services/auth/storage/mongodb/account"

	"github.com/egoholic/charcoal/services/auth/account"
	"github.com/egoholic/charcoal/services/auth/account/pwd"
	. "github.com/egoholic/charcoal/services/auth/account/repo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	name1 = "Joe Rogan"
	pwd1  = "12345678"
)

var _ = Describe("Accounts Repository", func() {
	Context("creation", func() {
		Describe("New()", func() {
			It("returns repository", func() {
				Expect(New()).To(BeAssignableToTypeOf(&Repo{}))
			})
		})
	})

	Describe(".Insert()", func() {
		It("persists given account", func() {
			a := account.New(name1, pwd.New(pwd1))
			repo := New()
			client, _ := mongo.NewClient(config.MongoDBConnectionString())
			client.Connect(context.TODO())
			Expect(repo.Insert(a, storage.NewInserter(context.TODO(), client))).To(BeNil())
		})
	})

	Describe(".FindByPK()", func() {
		It("finds an account by given PK and returns it", func() {
			a := account.New(name1, pwd.New(pwd1))
			repo := New()
			client, _ := mongo.NewClient(config.MongoDBConnectionString())
			client.Connect(context.TODO())
			repo.Insert(a, storage.NewInserter(context.TODO(), client))
			found, _ := repo.FindByPK(name1, storage.NewByPKFinder(context.TODO(), client))
			Expect(found.PK()).To(Equal(name1))
		})
	})
})
