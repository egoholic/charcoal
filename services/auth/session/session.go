package session

import (
	"net"
	"time"

	"github.com/egoholic/charcoal/services/auth/account"
)

const REMEMBER_ME_DURATION = 60 * 60 * 24 * 7 * time.Second

type Payload struct {
	Account  *account.Account
	Token    string
	IP       net.IP
	LastTime time.Time
}

type Session struct {
	payload *Payload
}

func NewPayload(ac *account.Account, tk string, ip net.IP, lt time.Time) *Payload {
	return &Payload{ac, tk, ip, lt}
}

func New(p *Payload) *Session {
	return &Session{p}
}

func (s *Session) PK() string {
	return s.Token()
}

func (s *Session) Token() string {
	return s.payload.Token
}

func (s *Session) IP() net.IP {
	return s.payload.IP
}

func (s *Session) LastTime() time.Time {
	return s.payload.LastTime
}

func (s *Session) Account() *account.Account {
	return s.payload.Account
}

func (s *Session) IsExpired() bool {
	lt := s.LastTime()
	expiration := lt.Add(REMEMBER_ME_DURATION)
	return time.Now().After(expiration)
}
