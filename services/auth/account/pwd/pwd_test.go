package pwd_test

import (
	. "github.com/egoholic/charcoal/services/auth/account/pwd"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	pwd1 = "12345678"
	pwd2 = "24681012"
)

var _ = Describe("Passwords domain", func() {
	Context("creation", func() {
		Describe("New()", func() {
			It("returns encrypted password", func() {
				ep := EncryptedPassword("")
				Expect(New(pwd1)).To(BeAssignableToTypeOf(&ep))
			})
		})
	})

	Context("type conversion", func() {
		Describe(".ToString()", func() {
			It("returns string", func() {
				ep := EncryptedPassword("12345678")
				Expect(ep.ToString()).To(BeAssignableToTypeOf(pwd1))
			})
		})
	})
})
