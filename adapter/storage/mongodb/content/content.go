package content

import (
	"context"

	"github.com/egoholic/charcoal/entity/content/chlog"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Adapter struct{}

func (a *Adapter) Insert(client mongo.Client) {
	return func(ctx context.Context, payload *chlog.Payload) (string, error) {
		db := client.Database("charcoal")
		contents := db.Collection("contents")
		doc := bson.D{{"title", payload.Title()}, {"body", payload.Body()}}
		res, err := contents.InsertOne(ctx, doc)
		if err != nil {
			return "", err
		}
		return res.InsertedID.(string), nil
	}
}

func FindByPK(ctx context.Context, pk string) {

}
