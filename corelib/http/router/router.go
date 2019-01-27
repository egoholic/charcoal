package router

import (
	"github.com/egoholic/charcoal/corelib/http/router/params"
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

func New() *Router {
	return &Router{NewNode("")}
}

func (r *Router) Root() *Node {
	return r.root
}

func (r *Router) Handler(p *params.Params) interface{} {
	return r.Root().Handler(p, p.NewIterator())
}

type Node struct {
	pathChunk    string
	children     map[string]*Node
	verbHandlers map[string]interface{}
}

func NewNode(chunk string) *Node {
	return &Node{chunk, map[string]*Node{}, map[string]interface{}{}}
}

func (n *Node) Sub(chunk string) *Node {
	var node *Node
	node = n.children[chunk]
	if node != nil {
		return node
	}

	node = NewNode(chunk)
	n.children[chunk] = node
	return node
}

func (n *Node) Handler(p *params.Params, iter *params.PathChunksIterator) interface{} {
	if iter.HasNext() {
		chunk, _ := iter.Next()
		if child, ok := n.children[chunk]; ok {
			return child.Handler(p, iter)
		}
		return nil
	}
	return n.verbHandlers[p.Verb()]
}

func (n *Node) GET(fn interface{}, d string) {
	n.verbHandlers[GET] = newHandler(fn, d)
}

func (n *Node) POST(fn interface{}, d string) {
	n.verbHandlers[POST] = newHandler(fn, d)
}

func (n *Node) PUT(fn interface{}, d string) {
	n.verbHandlers[PUT] = newHandler(fn, d)
}

func (n *Node) PATCH(fn interface{}, d string) {
	n.verbHandlers[PATCH] = newHandler(fn, d)
}

func (n *Node) DELETE(fn interface{}, d string) {
	n.verbHandlers[DELETE] = newHandler(fn, d)
}

type Handler struct {
	handlingFunction interface{}
	desription       string
}

func newHandler(fn interface{}, description string) *Handler {
	return &Handler{fn, description}
}

func (h *Handler) HandlingFunction() interface{} {
	return h.handlingFunction
}

func (h *Handler) Description() string {
	return h.desription
}
