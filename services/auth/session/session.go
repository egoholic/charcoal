package session

import (
	"net"
	"time"

	"github.com/egoholic/charcoal/services/auth/account"
	"github.com/egoholic/charcoal/services/auth/session/token"
)

const REMEMBER_ME_IN_SECONDS = 60 * 60 * 24 * 7 * time.Second

type Session struct {
	account      *account.Account
	token        string
	ip           net.IP
	lastSigninAt time.Time
}

func New(a *account.Account, ip net.IP) *Session {
	token := token.New()
	lastSigninAt := time.Now()
	return &Session{a, token, ip, lastSigninAt}
}

func (s *Session) Token() string {
	return s.token
}

func (s *Session) IP() net.IP {
	return s.ip
}

func (s *Session) IsExpired() bool {
	expiration := s.lastSigninAt.Add(REMEMBER_ME_IN_SECONDS)
	return s.lastSigninAt.After(expiration)
}
