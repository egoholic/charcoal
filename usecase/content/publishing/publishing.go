package publishing

import (
	"context"
	"time"

	"github.com/egoholic/charcoal/entity/content"
	"github.com/egoholic/charcoal/entity/publication"
)

type ContentPublishingUsecase struct{}
type PublicationInserter func(context.Context, *publication.Publication) error

func (uc *ContentPublishingUsecase) Play(ctx context.Context, urlID string, c *content.Content, publishedAt time.Time, insert PublicationInserter) {
	p := publication.New(urlID, c, publishedAt)
	insert(ctx, p)
}
