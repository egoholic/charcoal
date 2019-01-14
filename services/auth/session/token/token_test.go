package token_test

import (
	. "github.com/egoholic/charcoal/services/auth/session/token"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Token domain", func() {
	Context("creation", func() {
		Describe("New()", func() {
			It("returns token", func() {
				t := New()
				Expect(t).To(BeAssignableToTypeOf(""))
				Expect(t).To(HaveLen(LENGTH))
			})
		})
	})
})
