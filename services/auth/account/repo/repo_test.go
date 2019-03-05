package repo_test

import (
	"context"

	adapter "github.com/egoholic/charcoal/services/auth/storage/dumb/account"

	"github.com/egoholic/charcoal/services/auth/account"
	"github.com/egoholic/charcoal/services/auth/account/pwd"
	. "github.com/egoholic/charcoal/services/auth/account/repo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	name1   = "Joe Rogan"
	pwd1    = "12345678"
	storage = adapter.New()
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
			insert := repo.NewInserter(storage.NewInserter())
			Expect(insert(a)).To(BeNil())
		})
	})

	Describe(".NewByPKFinder()", func() {
		It("returns finder", func() {
			a := account.New(name1, pwd.New(pwd1))
			repo := New()
			insert := repo.NewInserter(storage.NewInserter())
			insert(a)
			find := repo.NewByPKFinder(storage.NewByPKFinder(context.TODO()))
			found, _ := find(name1)
			Expect(found.PK()).To(Equal(name1))
		})
	})
})
