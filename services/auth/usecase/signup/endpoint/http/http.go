package http

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/egoholic/charcoal/corelib/http/router"
)

func Extend(node *router.Node) error {
	signup := node.Sub("signup")
	signup.GET(renderSignupForm, "Renders sign up form")
	signup.POST(performSignup, "Performs sign up.")
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
