package comparable

type Comparison int

func (c *Comparison) IsGreater() bool { return *c == 1 }
func (c *Comparison) IsLess() bool    { return *c == -1 }
func (c *Comparison) IsEqual() bool   { return *c == 0 }
func (c *Comparison) Key() int        { return int(*c) }

type Comparable interface {
	Compare(Comparable) *Comparison
}
