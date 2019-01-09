package content

import (
	"context"
	"fmt"

	"github.com/egoholic/charcoal/config"
	"github.com/egoholic/charcoal/entity/content"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Adapter struct {
	client *mongo.Client
}

func New(c *mongo.Client) *Adapter {
	return &Adapter{c}
}

func (a *Adapter) Insert(ctx context.Context, c *content.Content) (interface{}, error) {
	client, err := a.getClient()
	if err != nil {
		panic(err)
	}
	db := client.Database("charcoal")
	collection := db.Collection("contents")
	snapshot := bson.D{{"title", c.Title()}, {"body", c.Body()}, {"_prev", nil}}
	doc := bson.D{{"PK", c.PK()}, {"snapshot", snapshot}}
	res, err := collection.InsertOne(ctx, doc)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return res.InsertedID, nil
}

func FindByPK(ctx context.Context, pk string) {

}

func (a *Adapter) getClient() (*mongo.Client, error) {
	var err error

	if a.client == nil {
		a.client, err = mongo.NewClient(config.MongoDB.ConnectionString())
		if err != nil {
			return nil, err
		}
	}

	return a.client, nil
}
