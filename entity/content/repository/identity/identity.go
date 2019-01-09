package identity

import (
	"fmt"

	"github.com/egoholic/charcoal/entity/content"
)

type StorePK interface{}
type keysMap map[content.PK]StorePK
type identityMap map[content.PK]*content.Content
type IdentityMap struct {
	km *keysMap
	im *identityMap
}

func NewMap() *IdentityMap {
	return &IdentityMap{&keysMap{}, &identityMap{}}
}

func (im *IdentityMap) Add(sPK StorePK, c *content.Content) error {
	bPK := c.PK()
	found := im.Get(bPK)
	if found != nil {
		return fmt.Errorf("Can't add %s -> %#v content.\n\tBecause %#v already exists.", string(bPK), *c, *found)
	} else {
		(*im.im)[bPK] = c
		(*im.km)[bPK] = sPK
		return nil
	}
}

func (im *IdentityMap) RefreshWith(c *content.Content) error {
	bPK := c.PK()
	old := im.Get(bPK)
	if old != nil {
		(*im.im)[bPK] = c // naive implementation.
		return nil
	} else {
		return fmt.Errorf("Can't refresh %s -> %#v content.\n\tBecause updating identity not exists.", string(bPK), *c)
	}
}

func (im *IdentityMap) Get(bPK content.PK) *content.Content {
	return (*im.im)[bPK]
}

func (im *IdentityMap) GetStorePK(bPK content.PK) StorePK {
	return (*im.km)[bPK]
}
