package session_test

import (
	"net"

	"github.com/egoholic/charcoal/services/auth/account"
	"github.com/egoholic/charcoal/services/auth/account/pwd"
	. "github.com/egoholic/charcoal/services/auth/session"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	ip  = net.ParseIP("127.0.0.1")
	acc = account.New("Peter Thiel", pwd.New("password"))
)

var _ = Describe("Session Entity", func() {
	Context("creation", func() {
		Describe("New()", func() {
			It("returns session", func() {
				s := New(acc, ip)
				Expect(s).To(BeAssignableToTypeOf(&Session{}))
			})
		})
	})

	Context("accessors", func() {
		Describe(".PK()", func() {
			It("returns primary key", func() {
				s := New(acc, ip)
				Expect(s.PK()).To(Equal(s.Token()))
			})
		})

		Describe(".Token()", func() {
			It("returns token", func() {
				s := New(acc, ip)
				Expect(s.Token()).NotTo(BeNil())
			})
		})

		Describe(".IP()", func() {
			It("returns IP addresss", func() {
				s := New(acc, ip)
				Expect(s.IP()).To(Equal(ip))
			})
		})

		Describe(".IsExpired()", func() {
			Context("when expired", func() {
				It("returns true", func() {
					// TODO
				})
			})

			Context("when not expired", func() {
				It("returns false", func() {
					s := New(acc, ip)
					Expect(s.IsExpired()).To(BeFalse())
				})
			})
		})
	})
})
