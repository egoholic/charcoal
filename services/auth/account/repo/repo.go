package repo

import (
	"github.com/egoholic/charcoal/corelib/serror"
	"github.com/egoholic/charcoal/services/auth/account"
	"github.com/egoholic/charcoal/services/auth/account/repo/idmap"
)

const (
	CACHE   = true
	NOCACHE = false
)

type Repo struct {
	identityMap *idmap.IdentityMap
}

func New() *Repo {
	return &Repo{idmap.New()}
}

type InsertAdapter func(*account.Account) (interface{}, error)

func (r *Repo) Insert(a *account.Account, insert InsertAdapter) error {
	sid, err := insert(a)
	if err != nil {
		return serror.Wrap(err)
	}

	err = r.identityMap.Add(a.PK(), sid, a)
	return err
}

type FindByPKAdapter func(interface{}) (*account.Account, error)

func (r *Repo) FindByPK(pk string, find FindByPKAdapter) (*account.Account, error) {
	sid, a, ok := r.identityMap.Get(pk)
	if ok {
		return a, nil
	} else {
		a, err := find(sid)
		if err != nil {
			return a, err
		}

		err = r.identityMap.Add(pk, sid, a)
		return a, err
	}
}
