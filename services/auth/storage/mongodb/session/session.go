package session

import (
	"context"
	"time"

	"github.com/egoholic/charcoal/services/auth/config"
	"github.com/egoholic/charcoal/services/auth/session"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

const COLLECTION_NAME = "sessions"

func NewInserter(ctx context.Context, client *mongo.Client) func(*session.Session) (interface{}, error) {
	return func(s *session.Session) (interface{}, error) {
		db := client.Database(config.MongoDBDatabaseName())
		sessions := db.Collection(COLLECTION_NAME)
		doc, err := Serialize(s)
		if err != nil {
			return nil, err
		}
		res, err := sessions.InsertOne(ctx, doc)
		return res.InsertedID, err
	}
}

func NewByPKFinder(ctx context.Context, client *mongo.Client) func(string) (interface{}, *session.Session, error) {
	return func(pk string) (interface{}, *session.Session, error) {
		db := client.Database(config.MongoDBDatabaseName())
		sessions := db.Collection(COLLECTION_NAME)
		filter := bson.D{{"token", pk}}
		res := sessions.FindOne(ctx, filter)
		sid, s, err := Deserialize(res)
		if err != nil {
			return sid, s, err
		}
		return sid, s, err
	}
}

func Serialize(s *session.Session) (bson.D, error) {
	return bson.D{{"token", s.Token()}, {"accountName", s.Account().Name()}, {"ip", s.IP().String()}, {"lastSigninAt", s.LastSigninAt().String()}}, nil
}

func SerializeWithStoreID(s *session.Session, sid interface{}) (bson.D, error) {
	return bson.D{{"_id", sid}, {"token", s.Token()}, {"accountName", s.Account().Name()}, {"ip", s.IP().String()}, {"lastSigninAt", s.LastSigninAt().String()}}, nil
}

type payload struct {
	id           interface{}
	token        string
	accountName  string
	ip           string
	lastSignInAt time.Time
}

type Decoder interface {
	Decode(interface{}) error
}

func Deserialize(d Decoder) (interface{}, *session.Session, error) {
	var p payload
	err := d.Decode(p)
	if err != nil {
		return nil, nil, err
	}
	s := session.New()
	return p.id, a, nil

}
