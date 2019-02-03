package serror_test

import (
	"fmt"

	. "github.com/egoholic/charcoal/corelib/serror"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Serror - Structured errors library", func() {
	Context("creation", func() {
		Describe("New()", func() {
			It("returns structured error", func() {
				Expect(New("description", "reason")).To(BeAssignableToTypeOf(&SError{}))
			})
		})

		Describe("Wrap()", func() {
			It("returns structured error", func() {
				err := fmt.Errorf("testError")
				Expect(Wrap(err)).To(BeAssignableToTypeOf(&SError{}))
			})
		})
	})

	Context("accessors", func() {
		Describe(".Error()", func() {
			Context("when wrapped", func() {
				It("returns error message", func() {
					err := fmt.Errorf("testError")
					serr := Wrap(err)
					Expect(serr.Error()).To(Equal("testError\n\tReason: -NONE-\n\t\tOriginal: testError"))
				})
			})

			Context("when created from the very beginning", func() {
				It("returns error message", func() {
					Expect(New("description", "reason").Error()).To(Equal("description\n\tReason: reason\n\t\tOriginal: -NONE-"))
				})
			})
		})
	})
})
