package idmap

import (
	"fmt"

	"github.com/egoholic/charcoal/services/auth/account"
	"github.com/egoholic/serror"
)

type IdentityMap struct {
	keyMap      map[string]interface{}
	identityMap map[string]*account.Account
}

func New() *IdentityMap {
	return &IdentityMap{map[string]interface{}{}, map[string]*account.Account{}}
}

func (im *IdentityMap) Get(k string) (interface{}, *account.Account, bool) {
	a := im.identityMap[k]
	sid := im.keyMap[k]
	return sid, a, (a != nil && sid != nil)
}

func (im *IdentityMap) Add(k string, sid interface{}, a *account.Account) error {
	_sid, _a, ok := im.Get(k)
	if ok {
		desc := fmt.Sprintf("Can't add account with key: %s", k)
		reason := fmt.Sprintf("Key `%s` already exists and there are %#v value and %#v SID under it.", k, _a, _sid)
		return serror.New(desc, reason)
	}

	im.keyMap[k] = sid
	im.identityMap[k] = a
	return nil
}
