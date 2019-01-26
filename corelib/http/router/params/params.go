package params

import (
	"strings"

	"github.com/egoholic/charcoal/corelib/validation"
)

type Form interface {
	Validate(interface{}) *validation.Node
}

type Params struct {
	path        string
	pathChunks  []string
	verb        string
	queryParams map[string]interface{}
	form        Form
}

func New(path, verb string, form Form) *Params {
	pathChunks := strings.Split(path, "/")
	queryParams := map[string]interface{}{}
	return &Params{path, pathChunks, verb, queryParams, form}
}

func (p *Params) Path() string {
	return p.path
}

func (p *Params) Verb() string {
	return p.verb
}

func (p *Params) NewIterator() *PathChunksIterator {
	return &PathChunksIterator{p.pathChunks, 0}
}

type PathChunksIterator struct {
	pathChunks []string
	cursor     int
}

func (i *PathChunksIterator) Next() (string, error) {
	return i.pathChunks[i.cursor+1], nil
}

func (i *PathChunksIterator) HasNext() bool {
	return len(i.pathChunks)-1 < i.cursor
}

func (i *PathChunksIterator) Current() string {
	return i.pathChunks[i.cursor]
}
