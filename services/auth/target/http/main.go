package main

import (
	"log"
	"net/http"
	"os"

	signin "github.com/egoholic/charcoal/services/auth/usecase/signin/endpoint/http"
	signup "github.com/egoholic/charcoal/services/auth/usecase/signup/endpoint/http"
	"github.com/egoholic/router"
)

func main() {
	r := router.New()
	root := r.Root()
	signup.Extend(root)
	signin.Extend(root)
	logger := log.New(os.Stdout, "auth", 0)
	logger.Panic(http.ListenAndServe(":8080", r))
}
