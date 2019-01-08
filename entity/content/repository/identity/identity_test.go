package identity_test

import (
	"github.com/egoholic/charcoal/entity/content"
	. "github.com/egoholic/charcoal/entity/content/repository/identity"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	title1 = "Title1"
	body1  = "Body1"
)

var _ = Describe("Identity", func() {
	Context("Creation", func() {
		Describe("NewMap()", func() {
			It("returns a pointer to identity map", func() {
				im := NewMap()
				Expect(im).To(BeAssignableToTypeOf(&IdentityMap{}))
			})
		})
	})
	Context("operations", func() {
		Describe(".Add()", func() {
			Context("when there is no identities with given key", func() {
				It("adds new identity to the map", func() {
					im := NewMap()
					content := content.New(title1, body1)
					sPK := "some-store-pk"
					// expected before conditions
					c := im.Get(content.PK())
					Expect(im.GetStorePK(content.PK())).To(BeNil())
					Expect(c).To(BeNil())
					im.Add(sPK, content)
					// expected change
					c = im.Get(content.PK())
					Expect(c).To(Equal(content))
					Expect(im.GetStorePK(content.PK())).To(Equal(sPK))
				})
			})

			Context("when there is already exists identity with given key", func() {
				It("returns error", func() {

				})
			})
		})
	})
})
