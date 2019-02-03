package http

import (
	"net/http"

	"github.com/egoholic/charcoal/corelib/http/router"
	"github.com/egoholic/charcoal/corelib/http/router/params"
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

func renderSignupForm(w http.ResponseWriter, r *http.Request, p *params.Params) {
	// filePath, err := filepath.Abs("./templates/signup.html")
	// if err != nil {
	// 	panic(err)
	// }
	// t, err := template.ParseFiles(filePath)
	// t.Execute(w, &SignapFormViewModel{"/signup/", http.MethodPost})
	w.Write([]byte("hello"))
	w.WriteHeader(200)
}

func performSignup(w http.ResponseWriter, r *http.Request, p *params.Params) {
	//form := form.New(r.Body.)
}
