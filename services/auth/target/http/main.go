package main

import (
	"net/http"

	httpTransport "github.com/egoholic/charcoal/services/auth/transport/http"
)

func main() {
	handler := httpTransport.New()
	http.Handle("/", handler)
}
