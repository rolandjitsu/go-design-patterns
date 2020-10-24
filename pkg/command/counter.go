package command

type Counter struct {
	Value int
}

func (c *Counter) Increment() {
	c.Value++
}

func (c *Counter) Decrement() {
	c.Value--
}
