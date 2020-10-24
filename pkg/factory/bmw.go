package factory

func NewBmwAssyLine(model string) AssyLine {
	return &bmwAssyLine{model}
}

type bmwAssyLine struct {
	model string
}

func (a *bmwAssyLine) Make(n int) <-chan Vehicle {
	c := make(chan Vehicle, n)

	go func() {
		for i := 1; i <= n; i++ {
			c <- &bmw{
				model: a.model,
				sn:    i,
			}
		}
		close(c)
	}()

	return c
}

type bmw struct {
	model string
	sn    int
}

func (c *bmw) Brand() string {
	return "BMW"
}

func (c *bmw) Model() string {
	return c.model
}

func (c *bmw) SN() int {
	return c.sn
}
