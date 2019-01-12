package account_test

import (
	. "github.com/egoholic/charcoal/services/auth/account"
	"github.com/egoholic/charcoal/services/auth/account/pwd"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	name1 = "James Deen"
	pwd1  = "12345678"
	name2 = "Rob Zombie"
	pwd2  = "24681012"
)

var _ = Describe("Account Entity", func() {
	Context("creation", func() {
		Describe("New()", func() {
			It("returns an account", func() {
				Expect(New(name1, pwd.New(pwd1))).To(BeAssignableToTypeOf(&Account{}))
			})
		})
	})

	Context("accessors", func() {
		Describe(".Name()", func() {
			It("returns name", func() {
				a := New(name1, pwd.New(pwd1))
				Expect(a.Name()).To(Equal(name1))
			})
		})

		Describe(".EncryptedPassword()", func() {
			It("returns encrypted password", func() {
				a := New(name1, pwd.New(pwd1))
				Expect(a.EncryptedPassword()).To(Equal(pwd1))
			})
		})

		Describe(".PK()", func() {
			It("returns PK", func() {
				a := New(name1, pwd.New(pwd1))
				Expect(a.PK()).To(Equal(name1))
			})
		})
	})

	Context("authentication", func() {
		Describe(".IsAuthenticableWith()", func() {
			a := New(name1, pwd.New(pwd1))

			Context("when correct password", func() {
				It("returns true", func() {
					Expect(a.IsAuthenticableWith(pwd1)).To(BeTrue())
				})
			})
			Context("when incorrect password", func() {
				It("returns false", func() {
					Expect(a.IsAuthenticableWith(pwd2)).To(BeFalse())
				})
			})
		})
	})
})
