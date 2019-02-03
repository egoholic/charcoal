package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/egoholic/charcoal/corelib/http/router"
	"github.com/egoholic/charcoal/corelib/http/router/params"
	signin "github.com/egoholic/charcoal/services/auth/usecase/signin/endpoint/http"
	signup "github.com/egoholic/charcoal/services/auth/usecase/signup/endpoint/http"
)

func main() {
	r := router.New()
	root := r.Root()
	signup.Extend(root)
	signin.Extend(root)
	logger := log.New(os.Stdout, "auth", 0)
	logger.Panic(http.ListenAndServe(":8080", r))
}

type h struct{}

func HandlerFunc(w http.ResponseWriter, r *http.Request, p *params.Params) {
	w.Write([]byte("hey"))
	w.WriteHeader(200)
	fmt.Println("Handler executed!")
}
