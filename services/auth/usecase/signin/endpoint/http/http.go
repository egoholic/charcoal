package http

import (
	"net/http"

	"github.com/egoholic/charcoal/corelib/http/router"
)

func Extend(node *router.Node) error {
	signin := node.Sub("signin")
	signin.GET(renderSigninForm, "Renders sign in form.")
	signin.POST(performSignin, "Performs sign in.")
	return nil
}

func renderSigninForm(w http.ResponseWriter, r *http.Request) {

}

func performSignin(w http.ResponseWriter, r *http.Request) {

}
