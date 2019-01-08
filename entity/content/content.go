package content

import (
	"github.com/egoholic/charcoal/entity/content/chlog"
	"github.com/egoholic/charcoal/framework/tconv"
)

type Content struct {
	snapshot *chlog.Snapshot
}

func New(title, body string) *Content {
	s := chlog.NewSnapshot()
	p := s.Payload()
	p.AssignTitle(tconv.MakeStringable(title))
	p.AssignBody(tconv.MakeStringable(body))
	return &Content{s}
}

func (c *Content) PK() string {
	return c.Title()
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
