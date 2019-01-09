package content_test

import (
	"context"

	. "github.com/egoholic/charcoal/adapter/storage/mongodb/content"
	"github.com/egoholic/charcoal/config"
	ce "github.com/egoholic/charcoal/entity/content"
	"github.com/mongodb/mongo-go-driver/mongo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	title1 = "Title1"
	body1  = "Body1"
)

var _ = Describe("Content MongoDB Adapter", func() {
	Context("creation", func() {
		Describe("New()", func() {
			It("returns adapter", func() {
				client, _ := mongo.NewClient(config.MongoDB.ConnectionString())
				Expect(New(client)).To(BeAssignableToTypeOf(&Adapter{}))
			})
		})
	})

	Context("operations", func() {
		Describe(".Insert()", func() {
			Context("when there is no contents with the same title", func() {
				It("inserts the given content", func() {
					client, _ := mongo.NewClient(config.MongoDB.ConnectionString())
					adapter := New(client)
					cnt := ce.New(title1, body1)
					result, _ := adapter.Insert(context.TODO(), cnt)
					Expect(result).To(BeAssignableToTypeOf(mongo.InsertOneResult{}))
				})
			})

			Context("when there is already exists a content with the same name", func() {
				It("fails and inserts nothing", func() {

				})
			})
		})

		Describe(".FindByPK()", func() {

		})
	})
})
