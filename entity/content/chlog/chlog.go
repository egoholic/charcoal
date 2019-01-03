package chlog

import (
	"time"

	"github.com/egoholic/charcoal/framework/tconv"
)

const (
	NO_RESON        = "no-reason"
	NOT_VALID_TITLE = "not-valid-title"
	NOT_VALID_BODY  = "not-valid-body"
)

type Payload struct {
	title tconv.Stringable
	body  tconv.Stringable
}

func (p *Payload) MakeCopy() *Payload             { return &Payload{p.title, p.body} }
func (p *Payload) Title() string                  { return p.title.ToString() }
func (p *Payload) Body() string                   { return p.body.ToString() }
func (p *Payload) AssignTitle(v tconv.Stringable) { p.title = v }
func (p *Payload) AssignBody(v tconv.Stringable)  { p.body = v }

type Commit struct {
	persistedAt time.Time
	commitID    string
}

func (c *Commit) PersistedAt() time.Time {
	return c.persistedAt
}

type Snapshot struct {
	payload *Payload
	reason  string
	commit  *Commit
	prev    *Snapshot
}

func NewSnapshot() *Snapshot {
	pl := &Payload{tconv.StringableNull(), tconv.StringableNull()}
	return &Snapshot{pl, NO_RESON, nil, nil}
}
func (s *Snapshot) MakeNext(reason string) *Snapshot {
	return &Snapshot{s.Payload().MakeCopy(), reason, nil, s}
}

func (s *Snapshot) PK() string        { return s.commit.commitID }
func (s *Snapshot) Payload() *Payload { return s.payload }
func (s *Snapshot) Reason() string    { return s.reason }
func (s *Snapshot) IsCommited() bool  { return s.commit != nil }
func (s *Snapshot) AssignBody(body string) {
	sb := tconv.MakeStringable(body)
	s.Payload().AssignBody(sb)
}

func (s *Snapshot) AssignTitle(title string) {
	st := tconv.MakeStringable(title)
	s.Payload().AssignTitle(st)
}
