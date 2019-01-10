package snapshot

import (
	"time"
)

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
	return &Snapshot{nil, NO_RESON, nil, nil}
}

func (s *Snapshot) MakeNext(reason string) *Snapshot {
	return &Snapshot{s.Payload().MakeCopy(), reason, nil, s}
}

func (s *Snapshot) PK() string        { return s.commit.commitID }
func (s *Snapshot) Payload() *Payload { return s.payload }
func (s *Snapshot) Reason() string    { return s.reason }
func (s *Snapshot) IsCommited() bool  { return s.commit != nil }
func (s *Snapshot) AssignBody(body string) {
	s.Payload().AssignBody(body)
}

func (s *Snapshot) AssignTitle(title string) {
	s.Payload().AssignTitle(title)
}
