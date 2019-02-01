package router_test

import (
	"net/http"

	"github.com/egoholic/charcoal/corelib/http/router/params"

	. "github.com/egoholic/charcoal/corelib/http/router"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)



var _ = Describe("Router", func() {
	func TestHandlerFunc(w http.ResponseWriter, r *http.Request) {

  }

	Context("creation", func() {
		Describe("New()", func() {
			It("returns router", func() {
				Expect(New()).To(BeAssignableToTypeOf(&Router{}))
			})
		})
	})

	Context("accessors", func() {
		Describe(".Root()", func() {
			It("returns root node", func() {
				r := New()
				Expect(r.Root()).To(BeAssignableToTypeOf(&Node{}))
			})
		})
	})

	Context("routing", func() {
		Describe(".Handler()", func() {
			Context("when route exists", func() {
				It("returns handler", func() {
					r := New()
					description := "description"
					resp := NewTestResponse()
					_prms := map[string][]string{}
					_prms["header"] = []string{"1"}
					_prms["value"] = []string{"test"}
					_prms["body"] = []string{"TestBody"}
					prms := params.New("/", GET, _prms, nil)
					r.Root().GET(TestHandlerFunc, description)
					h := r.Handler(prms)
					Expect(h).To(BeAssignableToTypeOf(&Handler{}))
					handlingFunc := h.HandlingFunction()
					Expect(string(resp.Body())).To(Equal(""))
					handlingFunc(prms, resp)
					Expect(resp.Body().String()).To(Equal("TestBody"))
					Expect(resp.Status()).To(Equal(200))
					Expect(h.Description()).To(Equal(description))

					_prms2 := map[string][]string{}
					_prms2["header"] = []string{"1"}
					_prms2["value"] = []string{"test"}
					_prms2["body"] = []string{"TestBody"}
					description2 := "description2"
					prms2 := params.New("/articles", GET, _prms2, nil)
					resp2 := response.New()

					r.Root().Sub("articles").GET(TestHandlerFunc, description2)
					h = r.Handler(prms2)
					Expect(h).To(BeAssignableToTypeOf(&Handler{}))
					handlingFunc2 := h.HandlingFunction()
					Expect(resp2.Body().String()).To(Equal(""))
					handlingFunc2(prms2, resp2)
					Expect(resp2.Body().String()).To(Equal("TestBody"))
					Expect(h.Description()).To(Equal(description2))
				})
			})

			Context("when route does not exist", func() {
				It("returns nil", func() {
					r := New()
					params := params.New("/", GET, map[string][]string{}, nil)
					Expect(r.Handler(params)).To(BeNil())
				})
			})
		})
	})

	Describe("Node", func() {
		var (
			router *Router
			root   *Node
		)

		BeforeEach(func() {
			router = New()
			root = router.Root()
		})

		Context("definition", func() {

			Describe(".Sub()", func() {
				Context("when node with given path chunk exists", func() {
					It("returns existing Node", func() {
						node2 := root.Sub("articles")
						node3 := root.Sub("articles")
						Expect(node3).To(BeIdenticalTo(node2))
					})
				})

				Context("when node with given path chunk not exists", func() {
					It("creates new Node and returns it", func() {
						Expect(root.Sub("articles")).To(BeAssignableToTypeOf(&Node{}))
					})
				})
			})

			Describe(".GET()", func() {
				It("defines GET request handler", func() {
					params := params.New("/articles", "GET", map[string][]string{}, nil)
					Expect(router.Handler(params)).To(BeNil())
					root.Sub("articles").GET(TestHandlerFunc, "description1")
					Expect(router.Handler(params)).NotTo(BeNil())
				})
			})

			Describe(".POST()", func() {
				It("defines POST request handler", func() {
					params := params.New("/articles", "POST", map[string][]string{}, nil)
					Expect(router.Handler(params)).To(BeNil())
					root.Sub("articles").POST(TestHandlerFunc, "description1")
					Expect(router.Handler(params)).NotTo(BeNil())
				})
			})

			Describe(".PUT()", func() {
				It("defines PUT request handler", func() {
					params := params.New("/", "PUT", map[string][]string{}, nil)
					Expect(router.Handler(params)).To(BeNil())
					root.PUT(TestHandlerFunc, "description1")
					Expect(router.Handler(params)).NotTo(BeNil())
				})
			})

			Describe(".PATCH()", func() {
				It("defines PATCH request handler", func() {
					params := params.New("/", "PATCH", map[string][]string{}, nil)
					Expect(router.Handler(params)).To(BeNil())
					root.PATCH(TestHandlerFunc, "description1")
					Expect(router.Handler(params)).NotTo(BeNil())
				})
			})

			Describe(".DELETE()", func() {
				It("defines DELETE request handler", func() {
					params := params.New("/articles", "DELETE", map[string][]string{}, nil)
					Expect(router.Handler(params)).To(BeNil())
					root.Sub("articles").DELETE(TestHandlerFunc, "description1")
					Expect(router.Handler(params)).NotTo(BeNil())
				})
			})
		})
	})
})
