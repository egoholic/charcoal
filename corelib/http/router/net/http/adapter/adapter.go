package adapter

import (
	"net/http"

	"github.com/egoholic/charcoal/corelib/http/router"
	"github.com/egoholic/charcoal/corelib/http/router/params"
)

func NewParams(r *http.Request, form params.Form) *params.Params {
	return params.New(r.RequestURI, r.Method, form)
}

type Handler struct {
	router *router.Router
}

type HandlingFunction = func(w http.ResponseWriter, p *params.Params)

func NewHandler(r *router.Router) http.Handler {
	return &Handler{r}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	prms := NewParams(r, nil)
	hfn := h.router.Handler(prms)

	fn, ok := hfn.(HandlingFunction)
	if ok {
		fn(w, prms)
	}
}
