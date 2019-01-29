package http

import (
	"log"
	"net/http"

	"github.com/egoholic/charcoal/corelib/http/router"
	"github.com/egoholic/charcoal/corelib/http/router/net/http/adapter"
)

type Transport struct {
	router *router.Router
}

func New() *Transport {
	r := router.New()
	r.Root().GET(renderSignupForm, "description")
	//signin.Extend(r.Root())
	//signup.Extend(r.Root())
	return &Transport{r}
}

func (t *Transport) Deliver() {
	log.Fatal(http.ListenAndServe(":8080", adapter.NewHandler(t.router)))
}

func renderSignupForm(w http.ResponseWriter, r *http.Request) {
	// filePath, err := filepath.Abs("./templates/signup.html")
	// if err != nil {
	// 	panic(err)
	// }
	// t, err := template.ParseFiles(filePath)
	// t.Execute(w, &SignapFormViewModel{"/signup/", http.MethodPost})
	w.Write([]byte("hello"))
	w.WriteHeader(200)
}
