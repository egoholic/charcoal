package pwd_test

import (
	. "github.com/egoholic/charcoal/value/auth/pwd"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	salt = []byte("So much salt to create password")
	pass = []byte("PasswordTest!@#)(*")
)

var _ = Describe("Passwords domain", func() {
	Context("creation", func() {
		Describe("EncryptPassword()", func() {
			It("returns encrypted password", func() {
				ep := EncryptPassword(pass, salt)
				Expect(EncryptPassword(pass, salt)).To(BeAssignableToTypeOf(ep))
			})
		})
	})

	Context("mathcing", func() {
		Describe("MatchPassword()", func() {
			It("returns logical result when match passwords", func() {
				matched := MatchPassword(pass, salt)
				Expect(MatchPassword(pass, salt)).To(BeAssignableToTypeOf(matched))
			})
		})
	})
})
