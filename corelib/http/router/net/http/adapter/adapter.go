package adapter

import (
	"net/http"

	"github.com/egoholic/charcoal/corelib/http/router"
	"github.com/egoholic/charcoal/corelib/http/router/params"
	"github.com/egoholic/charcoal/corelib/http/router/response"
)

func NewParams(r *http.Request, form params.Form) *params.Params {
	return params.New(r.RequestURI, r.Method, map[string][]string{}, form)
}

func NewResponse(w http.ResponseWriter) *response.Response {
	return &response.Response{}
}

type Adapter struct {
	router *router.Router
}

func New(r *router.Router) http.Handler {
	return &Adapter{r}
}

func (a *Adapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	prms := NewParams(r, nil)
	resp := response.New()
	h := a.router.Handler(prms)
	if h != nil {
		h.HandlingFunction()(prms, resp)
	} else {
		w.Write([]byte("Not found!"))
		w.WriteHeader(404)
	}
}
