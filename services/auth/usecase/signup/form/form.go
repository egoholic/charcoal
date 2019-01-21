package form

import (
	"github.com/egoholic/charcoal/corelib/validation"
)

const ROOT_TITLE = "signup"

type Form struct {
	formTitle            string
	login                string
	password             string
	passwordConfirmation string
	validationResult     *validation.Node
	uniquenessChecker    uniquenessChecker
}

type uniquenessChecker interface {
	IsUnique(string) bool
}

func New(title, login, password, passwordConfirmation string, checker uniquenessChecker) *Form {
	return &Form{title, login, password, passwordConfirmation, nil, checker}
}

func (f *Form) FormTitle() string {
	return f.formTitle
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

	if !f.uniquenessChecker.IsUnique(f.login) {
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
