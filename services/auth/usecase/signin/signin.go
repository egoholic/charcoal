package signin

import (
	"net"
	"time"

	"github.com/egoholic/charcoal/services/auth/account"
	"github.com/egoholic/charcoal/services/auth/session"
	"github.com/egoholic/charcoal/services/auth/session/token"
)

type ByPKFinder func(string) (*account.Account, error)
type SessionInserter func(*session.Session) error

func Signin(aName, password string, ip net.IP, findAccountByPK ByPKFinder, insertSession SessionInserter) (*session.Session, error) {
	a, err := findAccountByPK(aName)
	if err != nil {
		return nil, err
	}

	if a.IsAuthenticableWith(password) {
		p := session.NewPayload(a, token.New(), ip, time.Now())
		s := session.New(p)
		err := insertSession(s)
		if err != nil {
			return nil, err
		} else {
			return s, nil
		}
		return s, nil
	} else {
		return nil, nil
	}
}

func SigninWithoutCheck(a *account.Account, ip net.IP, insertSession SessionInserter) (*session.Session, error) {
	p := session.NewPayload(a, token.New(), ip, time.Now())
	s := session.New(p)
	err := insertSession(s)
	if err != nil {
		return nil, err
	} else {
		return s, nil
	}
}
