package mongodb

import (
	"context"

	"github.com/egoholic/charcoal/entity/auth/account"
	"github.com/egoholic/charcoal/entity/auth/account/storage/mongodb/config"
	"github.com/egoholic/charcoal/value/auth/pwd"
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
		if err != nil {
			return nil, err
		}
		if res == nil {
			return nil, nil
		}
		return res.InsertedID, err
	}
}

func NewByPKFinder(ctx context.Context, client *mongo.Client) func(string) (interface{}, *account.Account, error) {
	return func(pk string) (interface{}, *account.Account, error) {
		db := client.Database(config.MongoDBDatabaseName())
		accounts := db.Collection(COLLECTION_NAME)
		filter := bson.D{{"title", pk}}
		res := accounts.FindOne(ctx, filter)
		sid, a, err := Deserialize(res)
		if err != nil {
			return sid, a, err
		}
		return sid, a, err
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
