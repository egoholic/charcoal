package counter

type Counter struct {
	key   string
	count int64
}

type Persister func(*Counter) error
type ByKeyFinder func(string) Counter

func (c *Counter) Count() int64 {
	return c.count
}

func (c *Counter) Increment(persist Persister) int64 {
	c.count++
	persist(c)
	return c.count
}
