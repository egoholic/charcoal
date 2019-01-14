package account

import (
	"github.com/egoholic/charcoal/services/auth/account/pwd"
)

type Account struct {
	name              string
	encryptedPassword *pwd.EncryptedPassword
}

func New(name string, secret *pwd.EncryptedPassword) *Account {
	return &Account{name, secret}
}

func (a *Account) Name() string {
	return a.name
}

func (a *Account) EncryptedPassword() string {
	return a.encryptedPassword.ToString()
}

func (a *Account) PK() string {
	return a.Name()
}

func (a *Account) IsAuthenticableWith(p string) bool {
	return a.encryptedPassword.ToString() == pwd.New(p).ToString()
}
