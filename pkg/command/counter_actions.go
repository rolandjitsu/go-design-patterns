package command

func NewIncrementAction(c *Counter) Action {
	return &incAction{c}
}

type incAction struct {
	c *Counter
}

func (a *incAction) Execute() {
	a.c.Increment()
}

func (a *incAction) Undo() {
	a.c.Decrement()
}
