package serror_test

import (
	"fmt"

	. "github.com/egoholic/charcoal/corelib/serror"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var testErr = fmt.Errorf("<Test-Error>")

var _ = Describe("Serror - Structured errors library", func() {
	Context("creation", func() {
		Describe("New()", func() {
			It("returns structured error", func() {
				Expect(New("description", "reason")).To(BeAssignableToTypeOf(&SError{}))
			})
		})

		Describe("DumbDecorate()", func() {
			It("returns structured error", func() {
				Expect(DumbDecorate(testErr)).To(BeAssignableToTypeOf(&SError{}))
			})
		})

		Describe("Decorate()", func() {
			It("returns structured error", func() {
				Expect(Decorate(testErr, "reason")).To(BeAssignableToTypeOf(&SError{}))
			})
		})

		Describe("Wrap()", func() {
			It("returns structured error", func() {
				Expect(Wrap(testErr, "description1", "reason1")).To(BeAssignableToTypeOf(&SError{}))
			})
		})
	})

	Context("accessors", func() {
		Describe(".Error()", func() {
			Context("when wrapped", func() {
				It("returns error message", func() {
					testErr2 := Wrap(testErr, "description1", "reason1")
					Expect(testErr2.Error()).To(Equal("!Err: description1\n\tReason: reason1\n\tParent: <Test-Error>"))
					testErr3 := Wrap(testErr2, "description2", "reason2")
					Expect(testErr3.Error()).To(Equal("!Err: description2\n\tReason: reason2\n\tParent: !Err: description1\n\tReason: reason1\n\tParent: <Test-Error>"))
				})
			})

			Context("when dumb-decorated", func() {
				It("returns error message", func() {
					serr := DumbDecorate(testErr)
					Expect(serr.Error()).To(Equal("!Err: <Test-Error>\n\tReason: -NONE-\n\tDecorated: <Test-Error>"))
				})
			})

			Context("when decorated", func() {
				It("returns error message", func() {
					reason := "reason"
					serr := Decorate(testErr, reason)
					Expect(serr.Error()).To(Equal("!Err: <Test-Error>\n\tReason: reason\n\tDecorated: <Test-Error>"))
				})
			})

			Context("when created from the very beginning", func() {
				It("returns error message", func() {
					err := New("description", "reason")
					Expect(err.Error()).To(Equal("!Err: description\n\tReason: reason"))
				})
			})
		})
	})
})
