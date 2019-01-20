package http

import (
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Extend(r *mux.Router) error {
	prefixed := r.PathPrefix("/signup/").Subrouter()
	prefixed.HandleFunc("/", renderSignupForm).Methods("GET")
	prefixed.HandleFunc("/", performSignup).Methods("POST")
	return nil
}

func renderSignupForm(w http.ResponseWriter, r *http.Request) {
	tFile := "./templates/signup.html.template"
	t := template.Must(template.ParseGlob(tFile))
	t.Execute(os.Stdout, DATA)
}

func performSignup(w http.ResponseWriter, r *http.Request) {

}
