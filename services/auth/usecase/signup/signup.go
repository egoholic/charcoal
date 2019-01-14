package signup

import (
	"fmt"

	"github.com/egoholic/charcoal/corelib/serror"
	"github.com/egoholic/charcoal/services/auth/account"
	"github.com/egoholic/charcoal/services/auth/account/pwd"
)

type Inserter func(*account.Account) error

type ByPKFinder func(string) (*account.Account, error)

func Signup(name, password string, find ByPKFinder, insert Inserter) (*account.Account, error) {
	a := account.New(name, pwd.New(password))
	_a, _ := find(name)
	if _a != nil {
		return nil, serror.New("Can't signap.", fmt.Sprintf("Account `%s` already exists.", name))
	} else {
		err := insert(a)
		if err != nil {
			return nil, err
		}
		a, err := find(name)
		return a, err
	}
}
