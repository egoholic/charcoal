package repository

import (
	"context"

	"github.com/egoholic/charcoal/entity/content"
	"github.com/egoholic/charcoal/entity/content/repository/identity"
)

type Repo struct {
	identityMap *identity.IdentityMap
}

func New() *Repo {
	return &Repo{identity.NewMap()}
}

type InsertAdapter interface {
	Insert(*content.Content) (interface{}, error)
}

type FindByPKAdapter interface {
	FindByPK(context.Context, interface{}) (*bson.D, error)
}

func (r *Repo) Insert(ctx context.Context, i InsertAdapter, f FindByPKAdapter, c *content.Content) error {
	id, err := i.Insert(c)
	if err != nil {
		// todo: return custom error
	}

	content, err := f.FindByPK(ctx, id)
	if err != nil {
		// todo: return custom error
	}
	r.identityMap.Add(id, content)
	return nil
}

func (r *Repo) FindByPK(ctx context.Context, f FindByPKAdapter, pk *content.PK) (*content.Content, error) {
		doc, err := f.FindByPK(ctx, pk)
		if err != nil {
			// todo: return custom error
		}

    
		return nil, c
	}
}

func (im *IdentityMap) Upsert(c *content.Content) {

}
