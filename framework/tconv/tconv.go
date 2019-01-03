package tconv

const emptyString = ""

type Stringable interface {
	ToString() string
}

type stringableNull struct{}

func (sn *stringableNull) ToString() string { return emptyString }

func StringableNull() Stringable {
	return &stringableNull{}
}

type simpleStringable string

func (s *simpleStringable) ToString() string { return string(*s) }
func MakeStringable(str string) Stringable {
	s := simpleStringable(str)
	return &s
}
