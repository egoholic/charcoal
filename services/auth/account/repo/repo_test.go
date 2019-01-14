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

	Describe(".NewInserter()", func() {
		It("returns inserter", func() {
			a := account.New(name1, pwd.New(pwd1))
			repo := New()
			client, _ := mongo.NewClient(config.MongoDBConnectionString())
			client.Connect(context.TODO())
			insert := repo.NewInserter(storage.NewInserter(context.TODO(), client))
			Expect(insert(a)).To(BeNil())
		})
	})

	Describe(".NewByPKFinder()", func() {
		It("returns finder", func() {
			a := account.New(name1, pwd.New(pwd1))
			repo := New()
			client, _ := mongo.NewClient(config.MongoDBConnectionString())
			client.Connect(context.TODO())
			client.Connect(context.TODO())
			insert := repo.NewInserter(storage.NewInserter(context.TODO(), client))
			insert(a)
			find := repo.NewByPKFinder(storage.NewByPKFinder(context.TODO(), client))
			found, _ := find(name1)
			Expect(found.PK()).To(Equal(name1))
		})
	})
})
