package main

import (
	httpTransport "github.com/egoholic/charcoal/services/auth/transport/http"
)

func main() {
	app := httpTransport.New()
	app.Deliver()
}
