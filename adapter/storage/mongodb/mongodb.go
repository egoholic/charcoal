package mongodb

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func MakeClient(url string, timeoutInSeconds int) *mongo.Client {
	// TODO: use connection options from config / ENV.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, "mongodb://localhost:27017")
	return client
}
