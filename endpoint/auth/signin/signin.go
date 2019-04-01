package signin

import (
	"net"
	"time"

	"github.com/egoholic/charcoal/entity/auth/account"
	"github.com/egoholic/charcoal/entity/auth/session"
	"github.com/egoholic/charcoal/value/auth/token"
)

type ByPKFinder func(string) (interface{}, *account.Account, error)
type SessionInserter func(*session.Session) (interface{}, error)

func Signin(aName, password string, ip net.IP, findAccountByPK ByPKFinder, insertSession SessionInserter) (*session.Session, error) {
	_, acc, err := findAccountByPK(aName)
	if err != nil {
		return nil, err
	}

	if acc.IsAuthenticableWith(password) {
		p := session.NewPayload(acc, token.New(), ip, time.Now())
		s := session.New(p)
		_, err := insertSession(s)
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
	_, err := insertSession(s)
	if err != nil {
		return nil, err
	} else {
		return s, nil
	}
}
