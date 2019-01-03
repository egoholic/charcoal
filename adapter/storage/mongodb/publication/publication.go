package publication

import (
	"context"
	"time"

	"github.com/egoholic/charcoal/entities/publication"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Repository struct {
	client *mongo.Client
	ctx    context.Context
}

type P struct {
	name        string
	content     string
	publishedAt time.Time
}

func New(client *mongo.Client, ctx context.Context) *Repository {
	return &Repository{client, ctx}
}

func (r *Repository) Latest(n int) []*publication.Publication {
	sortCond := bson.M{"publishedAt": 1}
	var rec P
	ctx, cancel := context.WithCancel(r.ctx)
	defer cancel()
	cur, err := r.collection().Find(ctx, nil)
	if err != nil {

	}
	defer cur.Close(ctx)

	cur.Order(sortCond)
	if err != nil {

	}
	for cur.Next(ctx) {
		cur.Decode(rec)
	}
}

func (r *Repository) MostPopular(n int) []*publication.Publication {
	r.collection()
}

func (r *Repository) FindByName(name string) *publication.Publication {
	r.collection()
}

func (r *Repository) Persist(p *publication.Publication) bool {

}

func (r *Repository) collection() *mongo.Collection {
	r.client.Database("charcoalTest").Collection("publications")
}
