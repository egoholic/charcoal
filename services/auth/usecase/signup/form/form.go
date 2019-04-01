package form

import (
	"github.com/egoholic/validation"
)

const ROOT_TITLE = "signup"

type Form struct {
	login                string
	password             string
	passwordConfirmation string
	validationResult     *validation.Node
	uniquenessChecker    UniquenessChecker
}

type UniquenessChecker func(string) (bool, error)

func New(login, password, passwordConfirmation string, check UniquenessChecker) *Form {
	return &Form{login, password, passwordConfirmation, nil, check}
}

func (f *Form) Login() string {
	return f.login
}

func (f *Form) Password() string {
	return f.password
}

func (f *Form) PasswordConfirmation() string {
	return f.passwordConfirmation
}

func (f *Form) Validate() *validation.Node {
	if f.validationResult != nil {
		return f.validationResult
	} else {
		f.validationResult = validation.NewNode(ROOT_TITLE)
		f.validateLogin()
		f.validatePassword()
		return f.validationResult
	}
}

func (f *Form) validateLogin() {
	n := validation.NewNode("login")

	if len(f.login) < 5 {
		n.AddMessage("login must be at least 5 characters long")
	}

	if len(f.login) > 32 {
		n.AddMessage("login must be at most 32 characters long")
	}

	unique, err := f.uniquenessChecker(f.login)
	if err != nil {
		// TODO: add logging
		n.AddMessage("Sorry. Can not check uniqueness. Please try again later.")
	}
	if !unique {
		n.AddMessage("login must be unique")
	}

	f.validationResult.AddChild(n)
}

func (f *Form) validatePassword() {
	n := validation.NewNode("password")
	if len(f.login) < 8 {
		n.AddMessage("password must be at least 8 characters long")
	}

	if len(f.login) > 64 {
		n.AddMessage("password must be at most 64 characters long")
	}

	if f.password != f.passwordConfirmation {
		n.AddMessage("password does not match password confirmation")
	}

	f.validationResult.AddChild(n)
}
