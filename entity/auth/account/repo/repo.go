package repo

import (
	"fmt"

	"github.com/egoholic/charcoal/entity/auth/account"
	"github.com/egoholic/charcoal/entity/auth/account/repo/idmap"
	"github.com/egoholic/serror"
)

type Repo struct {
	identityMap *idmap.IdentityMap
}

func New() *Repo {
	return &Repo{idmap.New()}
}

type InsertAdapter func(*account.Account) (interface{}, error)

func (r *Repo) NewInserter(insert InsertAdapter) func(*account.Account) (interface{}, error) {
	return func(a *account.Account) (sid interface{}, err error) {
		sid, err = insert(a)
		if err != nil {
			err = serror.Wrap(err, "can't insert", serror.DEFAULT_REASON)
			return
		}

		err = r.identityMap.Add(a.PK(), sid, a)
		return
	}
}

type FindByPKAdapter func(string) (interface{}, *account.Account, error)

func (r *Repo) NewByPKFinder(find FindByPKAdapter) func(string) (interface{}, *account.Account, error) {
	return func(pk string) (sid interface{}, acc *account.Account, err error) {
		sid, acc, ok := r.identityMap.Get(pk)
		if ok {
			return
		}

		sid, acc, err = find(pk)
		if err != nil {
			err = serror.Decorate(err, fmt.Sprintf("can't find account by PK: %s", pk))
			return
		}

		err = r.identityMap.Add(acc.PK(), sid, acc)
		return
	}
}

func (r *Repo) NewUniquenessChecker(find FindByPKAdapter) func(string) (bool, error) {
	return func(login string) (result bool, err error) {
		result = false
		_, _, ok := r.identityMap.Get(login)
		result = !ok
		if result {
			return
		}

		_, account, err := find(login)
		if err != nil {
			return
		}
		result = account == nil
		return
	}
}
