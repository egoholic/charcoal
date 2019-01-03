package publishing

import (
	"github.com/egoholic/charcoal/entities/publication"
)

type PublishingUseCase struct{}

func (uc *PublishingUseCase) Play(persister publication.PublicationCreator) {
	p := publication.New()
	persister.Persist(p)

	return
}
