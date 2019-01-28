package http

import (
	"log"
	"net/http"

	"github.com/egoholic/charcoal/corelib/http/router"
	"github.com/egoholic/charcoal/corelib/http/router/net/http/adapter"

	signin "github.com/egoholic/charcoal/services/auth/usecase/signin/endpoint/http"
	signup "github.com/egoholic/charcoal/services/auth/usecase/signup/endpoint/http"
)

type Transport struct {
	router *router.Router
}

func New() *Transport {
	r := router.New()
	signin.Extend(r.Root())
	signup.Extend(r.Root())
	return &Transport{r}
}

func (t *Transport) Deliver() {
	log.Fatal(http.ListenAndServe(":8080", adapter.NewHandler(t.router)))
}
