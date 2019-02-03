package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/egoholic/charcoal/corelib/http/router"
	"github.com/egoholic/charcoal/corelib/http/router/params"
)

func main() {
	r := router.New()
	r.Root().GET(HandlerFunc, "desc")
	r.Root().Sub("test").GET(HandlerFunc, "ololo")
	logger := log.New(os.Stdout, "auth", 0)
	logger.Panic(http.ListenAndServe(":8080", r))
}

type h struct{}

func HandlerFunc(w http.ResponseWriter, r *http.Request, p *params.Params) {
	w.Write([]byte("hey"))
	w.WriteHeader(200)
	fmt.Println("Handler executed!")
}
