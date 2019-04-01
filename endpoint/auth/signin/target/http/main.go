package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/egoholic/charcoal/endpoint/auth/signin/config"
	"github.com/egoholic/router"
	"github.com/egoholic/router/params"
)

func main() {
	logger := log.New(os.Stdout, fmt.Sprintf("[ %s ] ", config.ServiceFullName()), 0)
	logger.Print("Service initialization!\n")
	rtr := router.New()
	root := rtr.Root()
	signin := root.Sub("signin")
	signin.GET(renderSigninForm, "Renders sign in form.")
	signin.POST(performSignin, "Performs sign in.")
	err := http.ListenAndServe(config.HTTPServicePort(), rtr)
	if err != nil {
		logger.Panicf("Panic: %s\n", err.Error())
	}
	logger.Printf("Service `%s` runned well!\n", config.ServiceInstanceName())
}

func renderSigninForm(w http.ResponseWriter, r *http.Request, p *params.Params) {

}

func performSignin(w http.ResponseWriter, r *http.Request, p *params.Params) {

}
