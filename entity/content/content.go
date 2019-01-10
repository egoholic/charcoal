package content

import (
	"github.com/egoholic/charcoal/entity/content/chlog"
	"github.com/egoholic/charcoal/framework/tconv"
)

type Created struct {
}
type Updated struct {
}
type Reviewed struct {
}
type Commented struct {
}
type Approved struct {
}

type Content struct {
	workingTitle string
	snapshot     *chlog.Snapshot
}

type PK string

func New(wtitle, title, body string) *Content {
	s := chlog.NewSnapshot()
	p := s.Payload()
	p.AssignTitle(tconv.MakeStringable(title))
	p.AssignBody(tconv.MakeStringable(body))
	return &Content{wtitle, s}
}

func (c *Content) PK() PK {
	return PK(c.WorkingTitle())
}

func (c *Content) WorkingTitle() string {
	return c.workingTitle
}

func (c *Content) Title() string {
	return c.Payload().Title()
}

func (c *Content) AssignTitle(v string) {
	c.Payload().AssignTitle(tconv.MakeStringable(v))
}

func (c *Content) Body() string {
	return c.Payload().Body()
}

func (c *Content) AssignBody(v string) {
	c.Payload().AssignBody(tconv.MakeStringable(v))
}

func (c *Content) Payload() *chlog.Payload {
	return c.snapshot.Payload()
}
