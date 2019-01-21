package http

import (
	"html/template"
	"net/http"

	"github.com/egoholic/charcoal/services/auth/account/repo"

	"github.com/egoholic/charcoal/services/auth/usecase/signup/form"
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
	form := form.New("", "", "", repo.NewUniquenessChecker())
	t.Execute(w, form)
}

func performSignup(w http.ResponseWriter, r *http.Request) {

}
