package adapter_test

import (
	"fmt"
	"net/http"

	"github.com/egoholic/charcoal/corelib/http/router"
	. "github.com/egoholic/charcoal/corelib/http/router/net/http/adapter"
	"github.com/egoholic/charcoal/corelib/http/router/params"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Mock struct {
	ResponseData []byte
	RequestData  []byte
	header       http.Header
	StatusCode   int
}

func NewMock(requestBody []byte) *Mock {
	header := http.Header{}
	return &Mock{[]byte{}, requestBody, header, 0}
}

func (m *Mock) Read(p []byte) (int, error) {
	maxLen := len(p)
	dataLen := len(m.RequestData)
	if maxLen > dataLen {
		maxLen = dataLen
	}

	for i := 0; i < maxLen; i++ {
		p[i] = m.RequestData[i]
	}

	return maxLen, nil
}

func (m *Mock) Header() http.Header {
	return m.header
}

func (m *Mock) Write(v []byte) (int, error) {
	var i = 0

	for i, b := range v {
		m.ResponseData[i] = b
	}

	return i, nil
}

func (m *Mock) WriteHeader(statusCode int) {
	m.StatusCode = statusCode
}

var _ = Describe("net/http server adapter", func() {
	Describe("Handler", func() {
		Context("creation", func() {
			Describe("NewHandler", func() {
				It("returns handler", func() {
					Expect(NewHandler(router.New())).To(BeAssignableToTypeOf(&Handler{}))
				})
			})
		})

		Context("request handling", func() {
			Describe(".ServeHTTP()", func() {
				It("serves HTTP successfully", func() {
					r := router.New()
					root := r.Root()
					root.GET(func(w http.ResponseWriter, r *http.Request) {
						body := make([]byte, 32)
						r.Body.Read(body)
						r.Body.Close()
						str := fmt.Sprintf("Hello `%s`!", body)
						w.Write([]byte(str))
					}, "welcomes")

					handler := NewHandler(r)
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
