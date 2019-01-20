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
		Describe("New()", func() {
			It("returns validation node", func() {
				Expect(New(title1)).To(BeAssignableToTypeOf(&Node{}))
			})
		})
	})

	Describe("accessors", func() {
		Describe("Title()", func() {
			It("returns title", func() {
				node := New(title1)
				Expect(node.Title()).To(Equal(title1))
			})
		})

		Describe("Messages()", func() {
			It("returns messages", func() {
				node := New(title1)
				Expect(node.Messages()).To(BeAssignableToTypeOf([]string{}))
			})
		})

		Describe("Children()", func() {
			It("returns children", func() {
				node := New(title1)
				Expect(node.Children()).To(BeAssignableToTypeOf(map[string]*Node{}))
			})
		})

		Describe("AddMessage()", func() {
			It("add message", func() {
				node := New(title1)
				Expect(node.Messages()).To(BeEmpty())
				node.AddMessage(message1)
				Expect(node.Messages()).To(HaveLen(1))
				Expect(node.Messages()[0]).To(Equal(message1))
			})
		})

		Describe("AddChild()", func() {
			It("add child", func() {
				node := New(title1)
				childNode := New(title2)
				Expect(node.Children()).To(BeEmpty())
				node.AddChild(childNode)
				Expect(node.Children()).To(HaveLen(1))
				Expect(node.Children()[title2]).To(Equal(childNode))
			})
		})
	})
})
