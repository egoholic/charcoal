package account

import "github.com/egoholic/charcoal/corelib/comparable"

type Account struct {
	name   string
	secret comparable.StringComparable // prefer account/pwd
}

func New(name string, secret comparable.StringComparable) *Account {
	return &Account{name, secret}
}

func (a *Account) Name() string {
	return a.name
}

func (a *Account) PK() string {
	return a.Name()
}

func (a *Account) IsAuthenticableWith(pwd string) bool {
	return a.secret.IsEqualString(pwd)
}
