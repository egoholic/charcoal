package viewing

import (
	"github.com/egoholic/charcoal/entities/publication"
)

type ViewingUseCase struct{}

func (v *ViewingUseCase) Play(name string, finder publication.PublicationByNameFinder, persister publication.Persister) publication.Publication {
	p := finder.FindByName(name)
	p.IncrementViews(persister)
	return p
}
