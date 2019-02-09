package serror

import "fmt"

const (
	DEFAULT_REASON           = "-NONE-"
	DEFAULT_ORIGINAL_MESSAGE = "-NONE-"
)

type SError struct {
	original    error
	description string
	reason      string
}

func New(d, r string) error {
	return &SError{nil, d, r}
}

func DumbWrap(e error) error {
	return &SError{e, e.Error(), DEFAULT_REASON}
}

func Wrap(e error, r string) error {
	return &SError{e, e.Error(), r}
}

func (e *SError) Error() string {
	var original string
	if e.original != nil {
		original = e.original.Error()
	} else {
		original = DEFAULT_ORIGINAL_MESSAGE
	}
	return fmt.Sprintf("%s\n\tReason: %s\n\t\tOriginal: %s", e.description, e.reason, original)
}
