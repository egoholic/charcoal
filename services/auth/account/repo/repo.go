package repo

import (
	"github.com/egoholic/charcoal/corelib/serror"
	"github.com/egoholic/charcoal/services/auth/account"
	"github.com/egoholic/charcoal/services/auth/account/repo/idmap"
)

type Repo struct {
	identityMap *idmap.IdentityMap
}

func New() *Repo {
	return &Repo{idmap.New()}
}

type InsertAdapter func(*account.Account) (interface{}, error)

func (r *Repo) NewInserter(insert InsertAdapter) func(*account.Account) error {
	return func(a *account.Account) error {
		sid, err := insert(a)
		if err != nil {
			return serror.Wrap(err)
		}

		err = r.identityMap.Add(a.PK(), sid, a)
		return err
	}
}

type FindByPKAdapter func(string) (interface{}, *account.Account, error)

func (r *Repo) NewByPKFinder(find FindByPKAdapter) func(string) (*account.Account, error) {
	return func(pk string) (*account.Account, error) {
		sid, a, ok := r.identityMap.Get(pk)
		if ok {
			return a, nil
		}

		sid, a, err := find(pk)
		if err != nil {
			return a, err
		}

		err = r.identityMap.Add(a.PK(), sid, a)
		return a, err
	}
}

type CheckUniquenessAdapter func(string) (bool, error)

func (r *Repo) NewUniquenessChecker(check CheckUniquenessAdapter) func(string) (bool, error) {
	return func(login string) (bool, error) {
		_, _, ok := r.identityMap.Get(login)
		if ok {
			return true, nil
		}

		return check(login)
	}
}
