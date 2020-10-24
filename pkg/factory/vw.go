package factory

func NewVwAssyLine(model string) AssyLine {
	return &vwAssyLine{model}
}

type vwAssyLine struct {
	model string
}

func (a *vwAssyLine) Make(n int) <-chan Vehicle {
	c := make(chan Vehicle, n)

	go func() {
		for i := 1; i <= n; i++ {
			c <- &vw{
				model: a.model,
				sn:    i,
			}
		}
		close(c)
	}()

	return c
}

type vw struct {
	model string
	sn    int
}

func (c *vw) Brand() string {
	return "Volkswagen"
}

func (c *vw) Model() string {
	return c.model
}

func (c *vw) SN() int {
	return c.sn
}
