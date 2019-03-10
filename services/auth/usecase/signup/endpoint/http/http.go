package http

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/egoholic/charcoal/services/auth/usecase/signup"
	"github.com/egoholic/charcoal/services/auth/usecase/signup/form"
	"github.com/egoholic/router"
	"github.com/egoholic/router/params"
)

func Extend(node *router.Node) error {
	signup := node.Sub("signup")
	signup.GET(renderSignupForm, "Renders sign up form")
	signup.POST(performSignup, "Performs sign up.")
	return nil
}

type SignupFormViewModel struct {
	FormAction string
	FormMethod string
}

func renderSignupForm(w http.ResponseWriter, r *http.Request, p *params.Params) {
	filePath, err := filepath.Abs("./services/auth/usecase/signup/endpoint/http/templates/signup.gohtml")
	if err != nil {
		panic(err)
	}
	t, err := template.ParseFiles(filePath)
	if err != nil {
		fmt.Printf("ERROR2: %s", err.Error())
	}
	err = t.Execute(w, &SignupFormViewModel{"/signup/", http.MethodPost})
	if err != nil {
		fmt.Printf("ERROR3: %s", err.Error())
	}
	w.WriteHeader(200)
}

func performSignup(w http.ResponseWriter, r *http.Request, p *params.Params) {
	fobj := form.New()
	vres := fobj.Validate()
	if vres.IsValid() {
		signup.Signup() // signup and render success
	} else {
		vres // return error messages
	}
}
