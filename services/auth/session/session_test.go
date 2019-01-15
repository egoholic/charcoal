package session_test

import (
	"net"
	"time"

	"github.com/egoholic/charcoal/services/auth/account"
	"github.com/egoholic/charcoal/services/auth/account/pwd"
	. "github.com/egoholic/charcoal/services/auth/session"
	"github.com/egoholic/charcoal/services/auth/session/token"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	ac      = account.New("Peter Thiel", pwd.New("password"))
	ip      = net.ParseIP("127.0.0.1")
	lt      = time.Now()
	tk      = token.New()
	payload = NewPayload(ac, tk, ip, lt)
)

var _ = Describe("Session Entity", func() {
	Context("creation", func() {
		Describe("New()", func() {
			It("returns session", func() {
				s := New(payload)
				Expect(s).To(BeAssignableToTypeOf(&Session{}))
			})
		})

		Describe("NewPayload()", func() {
			It("returns payload", func() {
				s := NewPayload(ac, tk, ip, lt)
				Expect(s).To(BeAssignableToTypeOf(&Payload{}))
			})
		})
	})

	Context("accessors", func() {
		Describe(".PK()", func() {
			It("returns primary key", func() {
				s := New(payload)
				Expect(s.PK()).To(Equal(s.Token()))
			})
		})

		Describe(".Account()", func() {
			It("returns account", func() {
				s := New(payload)
				Expect(s.Account()).To(BeAssignableToTypeOf(&account.Account{}))
			})
		})

		Describe(".Token()", func() {
			It("returns token", func() {
				s := New(payload)
				Expect(s.Token()).To(BeAssignableToTypeOf(""))
				Expect(s.Token()).To(HaveLen(token.LENGTH))
			})
		})

		Describe(".LastTime()", func() {
			It("returns time of the last sign in", func() {
				s := New(payload)
				Expect(s.LastTime()).To(Equal(lt))
			})
		})

		Describe(".IP()", func() {
			It("returns IP addresss", func() {
				s := New(payload)
				Expect(s.IP()).To(Equal(ip))
			})
		})

		Describe(".IsExpired()", func() {
			Context("when expired", func() {
				It("returns true", func() {
					l, _ := time.LoadLocation("UTC")
					lt := time.Date(2010, time.January, 1, 1, 1, 1, 1, l)
					payload := NewPayload(ac, tk, ip, lt)
					s := New(payload)
					Expect(s.IsExpired()).To(BeTrue())
				})
			})

			Context("when not expired", func() {
				It("returns false", func() {
					s := New(payload)
					Expect(s.IsExpired()).To(BeFalse())
				})
			})
		})
	})
})
