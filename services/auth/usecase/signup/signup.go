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

func Signup(name, password string, ip net.IP, findAccount AccountByPKFinder, insertAccount AccountInserter, insertSession SessionInserter) (acc *account.Account, ses *session.Session, err error) {
	_, acc, err = findAccount(name)
	if err != nil {
		return
	}
	if acc != nil {
		return nil, nil, serror.New("Can't signap.", fmt.Sprintf("Account `%s` already exists.", name))
	}

	acc = account.New(name, pwd.New(password))
	_, err = insertAccount(acc)
	if err != nil {
		err = serror.Decorate(err, "can't insert account")
		return
	}
	_, acc, err = findAccount(name)
	if err != nil {
		err = serror.Decorate(err, "can't find just inserted account")
		return
	}
	p := session.NewPayload(acc, token.New(), ip, time.Now())
	s := session.New(p)
	_, err = insertSession(s)
	if err != nil {
		err = serror.Decorate(err, "can't insert session")
		return
	}

	return
}
