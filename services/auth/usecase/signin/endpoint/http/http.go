package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Extend(r mux.Router) error {
	prefixed := r.PathPrefix("/signin/").Subrouter()
	prefixed.HandleFunc("/", renderSigninForm).Methods("GET")
	prefixed.HandleFunc("/", performSignin).Methods("POST")
	return nil
}

func renderSigninForm(w http.ResponseWriter, r *http.Request) {

}

func performSignin(w http.ResponseWriter, r *http.Request) {

}
