package router

import (
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

type HandlingFunction = func(*params.Params, *response.Response)

func New() *Router {
	return &Router{NewNode("")}
}

func (r *Router) Root() *Node {
	return r.root
}

func (r *Router) Handler(p *params.Params) *Handler {
	return r.Root().Handler(p, p.NewIterator())
}

type Node struct {
	pathChunk    string
	children     map[string]*Node
	verbHandlers map[string]*Handler
}

func NewNode(chunk string) *Node {
	return &Node{chunk, map[string]*Node{}, map[string]*Handler{}}
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

func (n *Node) Handler(p *params.Params, iter *params.PathChunksIterator) *Handler {
	if iter.HasNext() {
		chunk, _ := iter.Next()
		if child, ok := n.children[chunk]; ok {
			return child.Handler(p, iter)
		}
		return nil
	}
	return n.verbHandlers[p.Verb()]
}

func (n *Node) GET(fn HandlingFunction, d string) {
	n.verbHandlers[GET] = newHandler(fn, d)
}

func (n *Node) POST(fn HandlingFunction, d string) {
	n.verbHandlers[POST] = newHandler(fn, d)
}

func (n *Node) PUT(fn HandlingFunction, d string) {
	n.verbHandlers[PUT] = newHandler(fn, d)
}

func (n *Node) PATCH(fn HandlingFunction, d string) {
	n.verbHandlers[PATCH] = newHandler(fn, d)
}

func (n *Node) DELETE(fn HandlingFunction, d string) {
	n.verbHandlers[DELETE] = newHandler(fn, d)
}

type Handler struct {
	handlingFunction HandlingFunction
	desription       string
}

func newHandler(fn HandlingFunction, description string) *Handler {
	return &Handler{fn, description}
}

func (h *Handler) HandlingFunction() HandlingFunction {
	return h.handlingFunction
}

func (h *Handler) Description() string {
	return h.desription
}
