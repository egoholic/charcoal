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
	return c.payload().Title()
}

func (c *Content) AssignTitle(v string) {
	c.payload().AssignTitle(tconv.MakeStringable(v))
}

func (c *Content) Body() string {
	return c.payload().Body()
}

func (c *Content) AssignBody(v string) {
	c.payload().AssignBody(tconv.MakeStringable(v))
}

func (c *Content) payload() *chlog.Payload {
	return c.snapshot.Payload()
}
