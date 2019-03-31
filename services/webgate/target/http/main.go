package main

import (
	"net/http"

	"github.com/egoholic/charcoal/corelib/http/router/net/http/adapter"
	"github.com/egoholic/router"
)

var h http.Handler

func main() {
	r := router.New()
	root := r.Root()

	signup := root.Sub("signup")
	signup.GET(nil, "Renders Signup form.")
	signup.POST(nil, "Signs user up.")

	signin := root.Sub("signin")
	signin.GET(nil, "Renders Signin form.")
	signin.POST(nil, "Signs user in.")

	handler := adapter.NewHandler(r)
	http.ListenAndServe("/", handler)
}
