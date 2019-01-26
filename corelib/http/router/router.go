package router

import (
	"github.com/egoholic/charcoal/corelib/http/router/handler"
	"github.com/egoholic/charcoal/corelib/http/router/params"
	"github.com/egoholic/charcoal/corelib/http/router/response"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

type Router struct {
	root *Node
}

type Node struct {
	pathChunk    string
	children     map[string]*Node
	verbHandlers map[string]*handler.Handler
}

func New() *Router {
	children := map[string]*Node{}
	handlers := map[string]*handler.Handler{}
	Node := &Node{"/", children, handlers}
	return &Router{Node}
}

func (r *Router) Root() *Node {
	return r.root
}

func (r *Router) Handle(resp *response.Response, p *params.Params) {
	r.Root().Handle(resp, p, p.NewIterator())
}

func (n *Node) Node(chunk string) *Node {
	var node *Node

	node = n.children[chunk]
	if node != nil {
		return node
	}

	children := map[string]*Node{}
	children[chunk] = node
	handlers := map[string]*handler.Handler{}
	node = &Node{chunk, children, handlers}
	return node
}

func (n *Node) Handle(r *response.Response, p *params.Params, iter *params.PathChunksIterator) {
	if iter.HasNext() {
		chunk, _ := iter.Next()
		child := n.children[chunk]
		child.Handle(r, p, iter)
	}

	n.verbHandlers[p.Verb()].Handle(r, p)
}

func (n *Node) GET(fn handler.HandlingFunc, d string) {
	n.verbHandlers[GET] = handler.New(fn, d)
}

func (n *Node) POST(fn handler.HandlingFunc, d string) {
	n.verbHandlers[POST] = handler.New(fn, d)
}

func (n *Node) PUT(fn handler.HandlingFunc, d string) {
	n.verbHandlers[PUT] = handler.New(fn, d)
}

func (n *Node) PATCH(fn handler.HandlingFunc, d string) {
	n.verbHandlers[PATCH] = handler.New(fn, d)
}

func (n *Node) DELETE(fn handler.HandlingFunc, d string) {
	n.verbHandlers[DELETE] = handler.New(fn, d)
}
