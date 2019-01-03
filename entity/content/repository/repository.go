package repository

import (
	"github.com/egoholic/charcoal/entity/content"
)

type IdentityMap map[string]*content.Content

func (im *IdentityMap) Insert(c *content.Content) {
	(*im)[c.PK()] = c
}

func (im *IdentityMap) Find(k string) *content.Content {
	return (*im)[k]
}

func (im *IdentityMap) Upsert(c *content.Content) {

}
