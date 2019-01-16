package main

import (
	httpTransport "github.com/egoholic/charcoal/services/auth/transport/http"
)

func main() {
	transport := httpTransport.New()
	transport.Deliver()
}
