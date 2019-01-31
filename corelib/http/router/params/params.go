package params

import (
	"strings"

	"github.com/egoholic/charcoal/corelib/validation"
)

type Form interface {
	Validate(interface{}) *validation.Node
}

type Params struct {
	path       string
	pathChunks []string
	verb       string
	params     map[string][]string
	formObj    Form
}

func New(path, verb string, params map[string][]string, form Form) *Params {
	var chunks []string
	if path == "/" {
		chunks = []string{""}
	} else {
		chunks = strings.Split(path, "/")
	}

	return &Params{path, chunks, verb, params, form}
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

func (p *Params) Param(pname string) []string {
	if value, ok := p.params[pname]; ok {
		return value
	}

	return []string{}
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
