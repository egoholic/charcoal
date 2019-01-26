package handler

import (
	"github.com/egoholic/charcoal/corelib/http/router/params"
	"github.com/egoholic/charcoal/corelib/http/router/response"
)

type HandlingFunc func(*response.Response, *params.Params)
type Handler struct {
	fn         HandlingFunc
	desription string
}

func New(fn HandlingFunc, description string) *Handler {
	return &Handler{fn, description}
}

func (h *Handler) Handle(r *response.Response, p *params.Params) {
	h.fn(r, p)
}

func (h *Handler) Description() string {
	return h.desription
}
