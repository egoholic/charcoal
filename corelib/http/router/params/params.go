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
	var chunks []string
	if path == "/" {
		chunks = []string{""}
	} else {
		chunks = strings.Split(path, "/")
	}

	queryParams := map[string]interface{}{}
	return &Params{path, chunks, verb, queryParams, form}
}

func (p *Params) Path() string {
	return p.path
}
func (p *Params) PathChunks() []string {
	return p.pathChunks
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
	i.cursor++
	return i.Current(), nil
}

func (i *PathChunksIterator) HasNext() bool {
	max := len(i.pathChunks) - 1
	return i.cursor < max
}

func (i *PathChunksIterator) Current() string {
	return i.pathChunks[i.cursor]
}
