package repository

import (
	"context"

	"github.com/egoholic/charcoal/entity/content"
	"github.com/egoholic/charcoal/entity/content/chlog"
	"github.com/egoholic/charcoal/entity/content/repository/identity"
)

type Repo struct {
	identityMap *identity.IdentityMap
}

func New() *Repo {
	return &Repo{make(map[string]*content.Content)}
}

type InsertAdapter interface {
	Insert(*chlog.Payload) (interface{}, error)
}

func (r *Repo) MakeInserter(i InsertAdapter) func(context.Context, *content.Content) (string, error) {
	return func(ctx context.Context, c *content.Content) {
		id, _ := i.Insert(c.Payload())
		return id
	}
}

type FindByPKAdapter interface {
	FindByPK(context.Context, string) (*content.Content, error)
}

func (r *Repo) MakePKFinder(f FindByPKAdapter) *content.Content {
	return func(ctx context.Context, pk string) (*content.Content, error) {
		c, _ := f.FindByPK(ctx, pk)
		return nil, c
	}
}

func (im *IdentityMap) Upsert(c *content.Content) {

}
