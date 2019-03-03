package session

import (
	"context"

	"github.com/egoholic/charcoal/services/auth/session"
)

type Helper struct {
	store []*session.Session
}

func New() *Helper {
	return &Helper{[]*session.Session{}}
}

func (h *Helper) NewInserter() func(*session.Session) (interface{}, error) {
	return func(ssn *session.Session) (interface{}, error) {

	}
}

func (h *Helper) NewByPKFinder(ctx context.Context) {

}
