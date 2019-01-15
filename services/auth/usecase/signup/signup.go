package signup

import (
	"fmt"
	"net"
	"time"

	"github.com/egoholic/charcoal/corelib/serror"
	"github.com/egoholic/charcoal/services/auth/account"
	"github.com/egoholic/charcoal/services/auth/account/pwd"
	"github.com/egoholic/charcoal/services/auth/session"
	"github.com/egoholic/charcoal/services/auth/session/token"
)

type AccountByPKFinder func(string) (*account.Account, error)
type AccountInserter func(*account.Account) error
type SessionInserter func(*session.Session) error

func Signup(name, password string, ip net.IP, findAccount AccountByPKFinder, insertAccount AccountInserter, insertSession SessionInserter) (*account.Account, *session.Session, error) {
	_a, _ := findAccount(name)
	if _a != nil {
		return nil, nil, serror.New("Can't signap.", fmt.Sprintf("Account `%s` already exists.", name))
	}

	a := account.New(name, pwd.New(password))
	err := insertAccount(a)
	if err != nil {
		return nil, nil, err
	}
	a, err = findAccount(name)
	if err != nil {
		return nil, nil, err
	}
	p := session.NewPayload(a, token.New(), ip, time.Now())
	s := session.New(p)
	err = insertSession(s)
	if err != nil {
		return nil, nil, err
	}
	return a, s, nil

}
