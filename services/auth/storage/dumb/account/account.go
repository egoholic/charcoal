package account

import (
	"context"

	"github.com/egoholic/charcoal/services/auth/account"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Adapter struct {
	store []*account.Account
}

func New() *Adapter {
	return &Adapter{[]*account.Account{}}
}

func (a *Adapter) NewInserter() func(*account.Account) (interface{}, error) {
	return func(acc *account.Account) (interface{}, error) {
		id := len(a.store)
		a.store = append(a.store, acc)
		return id, nil
	}
}

func (a *Adapter) NewByPKFinder(ctx context.Context, client *mongo.Client) func(string) (interface{}, *account.Account, error) {
	return func(pk string) (interface{}, *account.Account, error) {
		for i, acc := range a.store {
			if acc.PK() == pk {
				return i, acc, nil
			}
		}

		return nil, nil, nil
	}
}
