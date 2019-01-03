package repository

import (
	"github.com/egoholic/charcoal/entity/publication"
)

type
type IdentityMap map[string]*publication.Publication

func (im *IdentityMap) Insert(c *publication.Publication) {
	(*im)[c.PK()] = c
}

func (im *IdentityMap) Find(k string) *publication.Publication {
	return (*im)[k]
}

func (im *IdentityMap) Upsert(c *publication.Publication) {

}
