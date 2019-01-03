package account

import (
	"github.com/egoholic/charcoal/entities/account/email"
	"github.com/egoholic/charcoal/entities/account/password"
)

type Account struct {
	name     email.Email
	password *password.Password
}

type Creator func(email.Email, password.Password) *Account
type ByNameFinder func(string) *Account

func New(emailStr, pwdStr string) Account {
	eml := email.New(emailStr)
	pwd := password.New(pwdStr)
	return Account{eml, &pwd}
}

func (a *Account) Name() string {
	return string(a.name)
}

func (a *Account) IsMatch(pwd string) bool {
	return (*a).password.IsEqual(pwd)
}
