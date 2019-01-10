package content_test

import (
	. "github.com/egoholic/charcoal/entity/content"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	wtitle1 = "Title 1 Draft"
	title1  = "Title1"
	body1   = "Body 1"
)

var _ = Describe("Content Entity", func() {
	Context("creation", func() {
		Describe("New()", func() {
			It("returns pointer on Content", func() {
				Expect(New(wtitle1, title1, body1)).To(BeAssignableToTypeOf(&Content{}))
			})
		})
	})
	Context("accessors", func() {
		Describe(".PK()", func() {
			It("returns primary key", func() {
				content := New(wtitle1, title1, body1)
				Expect(content.PK()).To(Equal(PK(wtitle1)))
			})
		})
		Describe(".WorkingTitle()", func() {
			It("returns working title", func() {
				content := New(wtitle1, title1, body1)
				Expect(content.WorkingTitle()).To(Equal(wtitle1))
			})
		})
		Describe(".Title()", func() {
			It("returns title", func() {
				content := New(wtitle1, title1, body1)
				Expect(content.Title()).To(Equal(title1))
			})
		})
		Describe(".AssignTitle()", func() {
			Context("when persisted", func() {
				It("returns title", func() {
					content := New(wtitle1, title1, body1)
					Expect(content.Title()).To(Equal(title1))
					newTitle := "New Title"
					content.AssignTitle(newTitle)
					Expect(content.Title()).To(Equal(newTitle))
				})
			})
			Context("when not persisted", func() {
				It("returns title", func() {
					content := New(wtitle1, title1, body1)
					Expect(content.Title()).To(Equal(title1))
				})
			})
		})
		Describe(".Body()", func() {
			It("returns body", func() {
				content := New(wtitle1, title1, body1)
				Expect(content.Body()).To(Equal(body1))
			})
		})
		Describe(".AssignBody()", func() {
			Context("when persisted", func() {
				It("returns body", func() {
					content := New(wtitle1, title1, body1)
					Expect(content.Body()).To(Equal(body1))
				})
			})
			Context("when not persisted", func() {
				It("returns body", func() {
					content := New(wtitle1, title1, body1)
					Expect(content.Body()).To(Equal(body1))
				})
			})
		})
		Describe(".Payload()", func() {
			It("returns payload", func() {
				content := New(wtitle1, title1, body1)
				payload := content.Payload()
				Expect(payload.Title()).To(Equal(title1))
				Expect(payload.Body()).To(Equal(body1))
			})
		})
	})
})
