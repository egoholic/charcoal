package comparable

type StringComparable interface {
	IsEqualString(string) bool
}
type simpleStringComparable struct {
	str string
}

func (c *simpleStringComparable) IsEqualString(s string) bool {
	return c.str == s
}

func NewStringComparable(s string) StringComparable {
	return &simpleStringComparable{s}
}
