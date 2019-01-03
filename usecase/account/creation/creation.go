package creation

import (
	"github.com/egoholic/charcoal/entities/account"
	"github.com/egoholic/charcoal/entities/account/email"
	"github.com/egoholic/charcoal/entities/account/password"
)

type CreationUseCase struct{}

func (uc *CreationUseCase) Play(name, pwd string, create account.Creator) account.Account {
	name := email.New(name)
	encryptedPwd := password.New(pwd)
	return create(name, encryptedPwd)
}
