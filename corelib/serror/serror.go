package serror

import "fmt"

const (
	DEFAULT_REASON           = "-NONE-"
	DEFAULT_ORIGINAL_MESSAGE = "-NONE-"
)

type SError struct {
	decorated   error
	description string
	reason      string
	parent      error
}

func New(d, r string) error {
	return &SError{nil, d, r, nil}
}

func Wrap(err error, d, r string) error {
	return &SError{nil, d, r, err}
}

func DumbDecorate(e error) error {
	return &SError{e, e.Error(), DEFAULT_REASON, nil}
}

func Decorate(e error, r string) error {
	return &SError{e, e.Error(), r, nil}
}

func (e *SError) Error() string {
	if e.decorated != nil {
		return fmt.Sprintf("!Err: %s\n\tReason: %s\n\tDecorated: %s", e.description, e.reason, e.decorated.Error())
	} else if e.parent != nil {
		return fmt.Sprintf("!Err: %s\n\tReason: %s\n\tParent: %s", e.description, e.reason, e.parent.Error())
	} else {
		return fmt.Sprintf("!Err: %s\n\tReason: %s", e.description, e.reason)
	}
}
