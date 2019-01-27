package adapter

import (
	"net/http"

	"github.com/egoholic/charcoal/corelib/http/router/params"
)

func New(r *http.Request, form params.Form) *params.Params {
	return params.New(r.RequestURI, r.Method, form)
}
