package identity

import (
	"fmt"

	"github.com/egoholic/charcoal/entity/content"
)

// in that way we hide implementation of IdentityMap pattern
type Key interface{}
type iMap map[Key]*content.Content
type IdentityMap struct {
	im *iMap
}

func NewMap() *IdentityMap {
	return &IdentityMap{&iMap{}}
}

func (im *IdentityMap) Check(k Key) (bool, *content.Content) {
	c := (*im.im)[k]
	return c != nil, c
}

func (im *IdentityMap) Add(c *content.Content) error {
	k := c.PK()
	contains, old := im.Check(k)
	if contains {
		return fmt.Errorf("Can't add %s -> %#v content\n\nBecause %#v already exists.", k, *c, *old)
	} else {
		(*im.im)[k] = c
		return nil
	}
}

func (im *IdentityMap) Delete(k Key) *content.Content {
	contains, content := im.Check(k)
	if contains {
		(*im.im)[k] = nil
		return content
	} else {
		return nil
	}
}

func (im *IdentityMap) Refresh(k string, c *content.Content) (bool, error) {
	contains, old := im.Check(k)
	if contains {
		(*im.im)[k] = nil
		return true, nil
	} else {
		return false, nil
	}
}
