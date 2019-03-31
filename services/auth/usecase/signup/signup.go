package signup

import (
	"fmt"
	"net"
	"time"

	"github.com/egoholic/charcoal/services/auth/account"
	"github.com/egoholic/charcoal/services/auth/account/pwd"
	"github.com/egoholic/charcoal/services/auth/session"
	"github.com/egoholic/charcoal/services/auth/session/token"
	"github.com/egoholic/serror"
)

type AccountByPKFinder func(string) (interface{}, *account.Account, error)
type AccountInserter func(*account.Account) (interface{}, error)
type SessionInserter func(*session.Session) (interface{}, error)

func Signup(name, password string, ip net.IP, findAccount AccountByPKFinder, insertAccount AccountInserter, insertSession SessionInserter) (*account.Account, *session.Session, error) {
	_, _a, _ := findAccount(name)
	if _a != nil {
		return nil, nil, serror.New("Can't signap.", fmt.Sprintf("Account `%s` already exists.", name))
	}

	a := account.New(name, pwd.New(password))
	_, err := insertAccount(a)
	if err != nil {
		return nil, nil, err
	}
	_, a, err = findAccount(name)
	if err != nil {
		return nil, nil, err
	}
	p := session.NewPayload(a, token.New(), ip, time.Now())
	s := session.New(p)
	_, err = insertSession(s)
	if err != nil {
		return nil, nil, err
	}
	return a, s, nil
}
