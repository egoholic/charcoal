package idmap

import (
	"fmt"

	"github.com/egoholic/charcoal/services/auth/session"
	"github.com/egoholic/serror"
)

type IdentityMap struct {
	keyMap      map[string]interface{}
	identityMap map[string]*session.Session
}

func New() *IdentityMap {
	return &IdentityMap{map[string]interface{}{}, map[string]*session.Session{}}
}

func (im *IdentityMap) Get(k string) (interface{}, *session.Session, bool) {
	s := im.identityMap[k]
	sid := im.keyMap[k]
	return sid, s, (s != nil && sid != nil)
}

func (im *IdentityMap) Add(k string, sid interface{}, s *session.Session) error {
	_sid, _s, ok := im.Get(k)
	if ok {
		desc := fmt.Sprintf("Can't add session with key: %s", k)
		reason := fmt.Sprintf("Key `%s` already exists and there are %#v value and %#v SID under it.", k, _s, _sid)
		return serror.New(desc, reason)
	}

	im.keyMap[k] = sid
	im.identityMap[k] = s
	return nil
}
