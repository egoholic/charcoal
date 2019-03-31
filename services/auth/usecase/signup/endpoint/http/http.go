package http

import (
	"context"
	"fmt"
	"html/template"
	"net"
	"net/http"
	"path/filepath"

	"github.com/egoholic/charcoal/services/auth/account/repo"
	"github.com/egoholic/charcoal/services/auth/config"
	"github.com/egoholic/charcoal/services/auth/storage/mongodb/account"
	"github.com/egoholic/charcoal/services/auth/storage/mongodb/session"
	"github.com/egoholic/charcoal/services/auth/usecase/signup"
	"github.com/egoholic/charcoal/services/auth/usecase/signup/form"
	"github.com/egoholic/router"
	"github.com/egoholic/router/params"
	"github.com/mongodb/mongo-go-driver/mongo"
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
	var (
		login                string
		password             string
		passwordConfirmation string
		ip                   net.IP
	)

	accounts := repo.New()
	ctx, cancel := context.WithTimeout(context.Background(), config.HTTPTimeOut())
	defer cancel()
	client, err := mongo.NewClient(config.MongoDBConnectionString())
	if err != nil {
		// TODO: add logging
		panic(err)
	}
	findAccount := account.NewByPKFinder(ctx, client)
	insertAccount := account.NewInserter(ctx, client)
	insertSession := session.NewInserter(ctx, client)
	fobj := form.New(login, password, passwordConfirmation, accounts.NewUniquenessChecker(findAccount))
	vres := fobj.Validate()
	if vres.IsValid() {
		signup.Signup(login, password, ip, findAccount, insertAccount, insertSession) // signup and render success
	} else {
		// return error messages
	}
}
