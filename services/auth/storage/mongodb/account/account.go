package account

import (
	"context"

	"github.com/egoholic/charcoal/services/auth/account"
	"github.com/egoholic/charcoal/services/auth/account/pwd"
	"github.com/egoholic/charcoal/services/auth/config"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

const COLLECTION_NAME = "accounts"

func NewInserter(ctx context.Context, client *mongo.Client) func(*account.Account) (interface{}, error) {
	return func(a *account.Account) (interface{}, error) {
		db := client.Database(config.MongoDBDatabaseName())
		accounts := db.Collection(COLLECTION_NAME)
		doc := bson.D{{"name", a.Name()}, {"encryptedPassword", a.EncryptedPassword()}}
		res, err := accounts.InsertOne(ctx, doc)
		return res.InsertedID, err
	}
}

func NewByPKFinder(ctx context.Context, client *mongo.Client) func(interface{}) (*account.Account, error) {
	return func(sid interface{}) (*account.Account, error) {
		db := client.Database(config.MongoDBDatabaseName())
		accounts := db.Collection(COLLECTION_NAME)
		filter := bson.D{{"_id", sid}}
		res := accounts.FindOne(ctx, filter)
		sid, a, err := Deserialize(res)
		if err != nil {
			return a, err
		}
		return a, err
	}
}

func Serialize(a *account.Account) (bson.D, error) {
	return bson.D{{"name", a.Name()}, {"encryptedPassword", a.EncryptedPassword()}}, nil
}

func SerializeWithStoreID(sid interface{}, a *account.Account) (bson.D, error) {
	return bson.D{{"_id", sid}, {"name", a.Name()}, {"encryptedPassword", a.EncryptedPassword()}}, nil
}

type payload struct {
	id                interface{}
	name              string
	encryptedPassword string
}

type Decoder interface {
	Decode(interface{}) error
}

func Deserialize(d Decoder) (interface{}, *account.Account, error) {
	var p payload
	err := d.Decode(p)
	if err != nil {
		return nil, nil, err
	}
	ep := pwd.EncryptedPassword(p.encryptedPassword)
	a := account.New(p.name, &ep)
	return p.id, a, nil
}
