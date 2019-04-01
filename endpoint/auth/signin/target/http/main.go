package main

import (
	"net/http"

	"github.com/egoholic/router"
	"github.com/egoholic/router/params"
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