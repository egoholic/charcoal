package http

import (
	"net/http"

	"github.com/egoholic/charcoal/corelib/http/router"
	"github.com/egoholic/charcoal/corelib/http/router/params"
)

func Extend(node *router.Node) error {
	signin := node.Sub("signin")
	signin.GET(renderSigninForm, "Renders sign in form.")
	signin.POST(performSignin, "Performs sign in.")
	return nil
}

func renderSigninForm(w http.ResponseWriter, r *http.Request, p *params.Params) {

}

func performSignin(w http.ResponseWriter, r *http.Request, p *params.Params) {

}
