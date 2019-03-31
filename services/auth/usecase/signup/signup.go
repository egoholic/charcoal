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

	if acc != nil {
		acc = nil
		err = serror.New("Can't sign up.", fmt.Sprintf("Account `%s` already exists.", name))
		return
	}

	acc = account.New(name, pwd.New(password))
	_, err = insertAccount(acc)
	if err != nil {
		acc = nil
		err = serror.Decorate(err, fmt.Sprintf("Can't insert account: `%s`", name))
		return
	}
	_, acc, err = findAccount(name)
	if err != nil {
		acc = nil
		err = serror.Decorate(err, fmt.Sprintf("Can't find just inserted account: `%s`", name))
		return
	}
	p := session.NewPayload(acc, token.New(), ip, time.Now())
	ses = session.New(p)
	_, err = insertSession(ses)
	if err != nil {
		acc = nil
		ses = nil
		err = serror.Decorate(err, fmt.Sprintf("Can't insert session for account: `%s`", name))
		return
	}

	return
}
