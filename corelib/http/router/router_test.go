package router_test

import (
	"github.com/egoholic/charcoal/corelib/http/router/params"

	. "github.com/egoholic/charcoal/corelib/http/router"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Router", func() {
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
					handler := "handler"
					description := "description"

					prms := params.New("/", GET, nil)
					r.Root().GET(handler, description)
					h := r.Handler(prms)
					Expect(h).To(BeAssignableToTypeOf(&Handler{}))
					hdlr, ok := h.(*Handler)
					Expect(ok).To(BeTrue())
					Expect(hdlr.HandlingFunction()).To(Equal(handler))
					Expect(hdlr.Description()).To(Equal(description))

					handler2 := "handler2"
					description2 := "description2"
					prms2 := params.New("/articles", GET, nil)

					r.Root().Sub("articles").GET(handler2, description2)
					h = r.Handler(prms2)
					Expect(h).To(BeAssignableToTypeOf(&Handler{}))
					hdlr, ok = h.(*Handler)
					Expect(ok).To(BeTrue())
					Expect(hdlr.HandlingFunction()).To(Equal(handler2))
					Expect(hdlr.Description()).To(Equal(description2))
				})
			})

			Context("when route does not exist", func() {
				It("returns nil", func() {
					r := New()
					params := params.New("/", GET, nil)
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
					params := params.New("/articles", "GET", nil)
					Expect(router.Handler(params)).To(BeNil())
					root.Sub("articles").GET("handler1", "description1")
					Expect(router.Handler(params)).NotTo(BeNil())
				})
			})

			Describe(".POST()", func() {
				It("defines POST request handler", func() {
					params := params.New("/articles", "POST", nil)
					Expect(router.Handler(params)).To(BeNil())
					root.Sub("articles").POST("handler1", "description1")
					Expect(router.Handler(params)).NotTo(BeNil())
				})
			})

			Describe(".PUT()", func() {
				It("defines PUT request handler", func() {
					params := params.New("/", "PUT", nil)
					Expect(router.Handler(params)).To(BeNil())
					root.PUT("handler1", "description1")
					Expect(router.Handler(params)).NotTo(BeNil())
				})
			})

			Describe(".PATCH()", func() {
				It("defines PATCH request handler", func() {
					params := params.New("/", "PATCH", nil)
					Expect(router.Handler(params)).To(BeNil())
					root.PATCH("handler1", "description1")
					Expect(router.Handler(params)).NotTo(BeNil())
				})
			})

			Describe(".DELETE()", func() {
				It("defines DELETE request handler", func() {
					params := params.New("/articles", "DELETE", nil)
					Expect(router.Handler(params)).To(BeNil())
					root.Sub("articles").DELETE("handler1", "description1")
					Expect(router.Handler(params)).NotTo(BeNil())
				})
			})
		})
	})
})
