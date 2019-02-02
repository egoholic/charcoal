package main

import (
	"net/http"

	"github.com/egoholic/charcoal/corelib/http/router"
)

func main() {
	r := router.New()
	// r.Root().GET(, "desc")
	http.ListenAndServe(":8080", r)
}

type h struct{}

func (h *h) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hey"))
	w.WriteHeader(200)
}
