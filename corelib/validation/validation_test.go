package validation_test

import (
	. "github.com/egoholic/charcoal/corelib/validation"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	title1   = "Title1"
	title2   = "Title2"
	message1 = "message1"
)

var _ = Describe("Validation corelib", func() {
	Describe("creation", func() {
		Describe("NewNode()", func() {
			It("returns validation node", func() {
				Expect(NewNode(title1)).To(BeAssignableToTypeOf(&Node{}))
			})
		})
	})

	Describe("accessors", func() {
		var (
			node1 *Node
			node2 *Node
		)

		BeforeEach(func() {
			node1 = NewNode(title1)
			node2 = NewNode(title2)
		})

		Describe(".Title()", func() {
			It("returns title", func() {
				Expect(node1.Title()).To(Equal(title1))
			})
		})

		Describe(".Messages()", func() {
			It("returns messages", func() {
				Expect(node1.Messages()).To(BeAssignableToTypeOf([]string{}))
			})
		})

		Describe(".Children()", func() {
			It("returns children", func() {
				Expect(node1.Children()).To(BeAssignableToTypeOf(map[string]*Node{}))
			})
		})

		Describe(".AddMessage()", func() {
			It("add message", func() {
				Expect(node1.Messages()).To(BeEmpty())
				node1.AddMessage(message1)
				Expect(node1.Messages()).To(HaveLen(1))
				Expect(node1.Messages()[0]).To(Equal(message1))
			})
		})

		Describe(".AddChild()", func() {
			It("add child", func() {
				Expect(node1.Children()).To(BeEmpty())
				node1.AddChild(node2)
				Expect(node1.Children()).To(HaveLen(1))
				Expect(node1.Children()[title2]).To(Equal(node2))
			})
		})
	})
})
