package http

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

func Extend(r *mux.Router) error {
	prefixed := r.PathPrefix("/signup/").Subrouter()
	prefixed.HandleFunc("/", renderSignupForm).Methods("GET")
	prefixed.HandleFunc("/", performSignup).Methods("POST")
	return nil
}

type SignapFormViewModel struct {
	FormAction string
	FormMethod string
}

func renderSignupForm(w http.ResponseWriter, r *http.Request) {
	filePath, err := filepath.Abs("./templates/signup.html")
	if err != nil {
		panic(err)
	}
	t, err := template.ParseFiles(filePath)
	t.Execute(w, &SignapFormViewModel{"/signup/", http.MethodPost})
	w.WriteHeader(200)
}

func performSignup(w http.ResponseWriter, r *http.Request) {
	//form := form.New(r.Body.)
}
