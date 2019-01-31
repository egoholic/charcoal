package adapter_test

import (
	"net/http"

	"github.com/egoholic/charcoal/corelib/http/router"
	. "github.com/egoholic/charcoal/corelib/http/router/net/http/adapter"
	"github.com/egoholic/charcoal/corelib/http/router/params"
	. "github.com/egoholic/charcoal/corelib/http/router/test/helper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("net/http server adapter", func() {
	Describe("Adapter", func() {
		Context("creation", func() {
			Describe("New", func() {
				It("returns adapter", func() {
					Expect(New(router.New())).To(BeAssignableToTypeOf(&Adapter{}))
				})
			})
		})

		Context("request handling", func() {
			Describe(".ServeHTTP()", func() {
				It("serves HTTP successfully", func() {
					r := router.New()
					root := r.Root()
					root.GET(DumbHandler, "welcomes")

					handler := New(r)
					mock := NewMock([]byte("James"))
					Expect(string(mock.ResponseData)).To(BeEmpty())
					request, _ := http.NewRequest("GET", "/", mock)
					handler.ServeHTTP(mock, request)
					Expect(string(mock.ResponseData)).To(Equal("Hello `James`!"))
				})
			})
		})
	})

	Describe("NewParams()", func() {
		It("returns params", func() {
			request, _ := http.NewRequest("GET", "", NewMock([]byte("James")))
			prms := NewParams(request, nil)
			Expect(prms).To(BeAssignableToTypeOf(&params.Params{}))
		})
	})
})
