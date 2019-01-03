package signingin

import (
	"github.com/egoholic/charcoal/entities/account"
)

type AuthenticationUseCase struct{}

func (uc *AuthenticationUseCase) Play(name, pwd string, accountByName account.ByNameFinder) {
	acc := accountByName(name)
	if acc.IsMatch(pwd) {

	} else {

	}
}
