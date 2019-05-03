package pwd_test

import (
	. "github.com/egoholic/charcoal/value/auth/pwd"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	salt       = []byte("So much salt to create password")
	pass       = []byte("PasswordTest!@#)(*")
	hashedPass = []byte("yjWSUgh0GFBND4UtMwfpSFiuaSanChqFmu7Yw5FtE3j89vGkS+cTUroMvG4i93CvJ6X02Jhhmv9WQdStHnEz7w==")
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
				matched := MatchPassword(pass, salt, hashedPass)
				Expect(MatchPassword(pass, salt, hashedPass)).To(BeAssignableToTypeOf(matched))
			})
		})
	})
})
