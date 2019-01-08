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
	title2 = "Title2"
	body2  = "Body2"
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
					sPK := "spk1"
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

		Describe(".RefreshWith()", func() {
			Context("when refreshing identity exists", func() {
				var (
					im  = NewMap()
					cnt *content.Content
					sPK = "spk1"
				)

				BeforeEach(func() {
					cnt = content.New(title1, body1)
					im.Add(sPK, cnt)
				})

				It("updates the identity and returns nothing", func() {
					cnt = im.Get(content.PK(title1))
					Expect(cnt.Body()).To(Equal(body1))
					cnt2 := content.New(title1, body2)
					im.RefreshWith(cnt2)
					cnt = im.Get(content.PK(title1))
					Expect(cnt.Body()).To(Equal(body2))
				})
			})
		})
	})
})
