package main

import (
	"net/http"

	"github.com/egoholic/charcoal/corelib/http/router/net/http/adapter"

	"github.com/egoholic/charcoal/corelib/http/router"
)

func main() {
	r := router.New()
	r.Root().GET((&h{}).ServeHTTP, "desc")
	http.ListenAndServe(":8080", adapter.NewHandler(r))
}

type h struct{}

func (h *h) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hey"))
	w.WriteHeader(200)
}
