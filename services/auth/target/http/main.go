package main

import (
	"log"
	"net/http"
	"os"

	signin "github.com/egoholic/charcoal/services/auth/usecase/signin/endpoint/http"
	signup "github.com/egoholic/charcoal/services/auth/usecase/signup/endpoint/http"
	"github.com/egoholic/router"
	"github.com/egoholic/router/params"
)

func main() {
	r := router.New()
	root := r.Root()
	root.GET(renderSignupForm, "renders sign up page")
	signup.Extend(root)
	signin.Extend(root)
	logger := log.New(os.Stdout, "Auth service started!\n", 0)
	logger.Panic(http.ListenAndServe(":8080", r))
}

func renderSignupForm(w http.ResponseWriter, r *http.Request, params *params.Params) {
	// filePath, err := filepath.Abs("./templates/signup.html")
	// if err != nil {
	// 	panic(err)
	// }
	// t, err := template.ParseFiles(filePath)
	// t.Execute(w, &SignapFormViewModel{"/signup/", http.MethodPost})
	w.Write([]byte("hello"))
	w.WriteHeader(200)
}
