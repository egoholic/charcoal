package http

import (
	"net/http"

	signin "github.com/egoholic/charcoal/services/auth/usecase/signin/endpoint/http"
	signup "github.com/egoholic/charcoal/services/auth/usecase/signup/endpoint/http"
	"github.com/gorilla/mux"
)

type Transport struct {
	router *mux.Router
}

func New() *Transport {
	r := mux.NewRouter()
	signup.Extend(r)
	signin.Extend(r)
	return &Transport{r}
}

func (t *Transport) Run() {
	http.Handle("/", t.router)
}
