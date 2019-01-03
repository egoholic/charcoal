package content

import (
	"github.com/egoholic/charcoal/entity/content"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func MakeInserter(conn mongo.Connection) content.Inserter {
  return func (ctx Context)
}
