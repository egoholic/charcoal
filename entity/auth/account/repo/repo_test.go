package repo_test

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"

	storage "github.com/egoholic/charcoal/entity/auth/account/storage/mongodb"
	"github.com/egoholic/charcoal/entity/auth/account/storage/mongodb/config"

	"github.com/egoholic/charcoal/entity/auth/account"
	. "github.com/egoholic/charcoal/entity/auth/account/repo"
	"github.com/egoholic/charcoal/value/auth/pwd"
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
			_, err := insert(a)
			Expect(err).To(BeNil())
		})
	})

	Describe(".NewByPKFinder()", func() {
		It("returns finder", func() {
			a := account.New(name1, pwd.New(pwd1))
			repo := New()
			client, _ := mongo.NewClient(config.MongoDBConnectionString())
			client.Connect(context.TODO())
			insert := repo.NewInserter(storage.NewInserter(context.TODO(), client))
			insert(a)
			find := repo.NewByPKFinder(storage.NewByPKFinder(context.TODO(), client))
			_, found, _ := find(name1)
			Expect(found.PK()).To(Equal(name1))
		})
	})
})
