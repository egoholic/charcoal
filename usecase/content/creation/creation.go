package creation

import "github.com/egoholic/charcoal/entity/content"

type ContentCreationUsecase struct{}
type ContentInserter func(*content.Content) error

func (cu *ContentCreationUsecase) Play(title, body string, insert ContentInserter) {
	ct := content.New(title, body)
	insert(ct)
}
