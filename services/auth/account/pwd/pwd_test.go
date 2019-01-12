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

	Context("comparison", func() {
		Describe(".IsEqualString()", func() {
			p := New(pwd1)

			Context("when matches string", func() {
				It("returns true", func() {
					Expect(p.IsEqualString(pwd1)).To(BeTrue())
				})
			})
			Context("when does not match string", func() {
				It("returns true", func() {
					Expect(p.IsEqualString(pwd2)).To(BeFalse())
				})
			})
		})
	})
})
