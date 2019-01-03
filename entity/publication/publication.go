package publication

import (
	"time"

	"github.com/egoholic/charcoal/entity/content"
)

type Publication struct {
	urlID       string
	content     *content.Content
	publishedAt time.Time
}

func New(urlID string, content *content.Content, publishedAt time.Time) *Publication {
	return &Publication{urlID, content, publishedAt}
}

func (p *Publication) PK() string {
	return p.URLID()
}

func (p *Publication) URLID() string {
	return p.urlID
}

func (p *Publication) Content() content.Content {
	return *p.content
}

func (p *Publication) PublishedAt() time.Time {
	return p.publishedAt
}
