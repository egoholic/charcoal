package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/egoholic/charcoal/endpoint/auth/signup"
	"github.com/egoholic/charcoal/endpoint/auth/signup/config"
	"github.com/egoholic/ruid"

	"github.com/egoholic/charcoal/endpoint/auth/signup/form"
	accountRepo "github.com/egoholic/charcoal/entity/auth/account/repo"
	accountStore "github.com/egoholic/charcoal/entity/auth/account/storage/mongodb"

	sessionRepo "github.com/egoholic/charcoal/entity/auth/session/repo"
	sessionStore "github.com/egoholic/charcoal/entity/auth/session/storage/mongodb"

	"github.com/egoholic/router"
	"github.com/egoholic/router/params"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func main() {
	logger := log.New(os.Stdout, fmt.Sprintf("[ %s ] ", config.ServiceFullName()), 0)
	logger.Print("Service initialization!\n")
	rtr := router.New()
	root := rtr.Root()
	signup := root.Sub("signup")
	signup.GET(renderSignupForm, "Renders sign up form")
	signup.POST(performSignup, "Performs sign up.")
	err := http.ListenAndServe(config.HTTPServicePort(), ruid.WrapHandlerWithRUID([]string{"signup"}, rtr))
	if err != nil {
		logger.Panicf("Panic: %s\n", err.Error())
	}
}

type SignupFormViewModel struct {
	FormAction string
	FormMethod string
}

func renderSignupForm(w http.ResponseWriter, r *http.Request, p *params.Params) {
	filePath, err := filepath.Abs("./templates/signup.gohtml")
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

	ctx, cancel := context.WithTimeout(context.Background(), config.HTTPTimeOut())
	defer cancel()
	client, err := mongo.NewClient(config.MongoDBConnectionString())
	if err != nil {
		// TODO: add logging
		panic(err)
	}

	accounts := accountRepo.New()
	accountInserter := accountStore.NewInserter(ctx, client)
	accountFinder := accountStore.NewByPKFinder(ctx, client)
	sessions := sessionRepo.New()
	sessionInserter := sessionStore.NewInserter(ctx, client)
	findAccount := accounts.NewByPKFinder(accountFinder)
	insertAccount := accounts.NewInserter(accountInserter)
	insertSession := sessions.NewInserter(sessionInserter)
	fobj := form.New(login, password, passwordConfirmation, accounts.NewUniquenessChecker(findAccount))
	vres := fobj.Validate()

	if vres.IsValid() {
		signup.Signup(login, password, ip, findAccount, insertAccount, insertSession) // signup and render success
	} else {
		// return error messages
	}
}
