package factorytest

import (
	"github.com/rolandjitsu/go-design-patterns/pkg/factory"
)

type MockAssyLine struct {
	Brand string
	Model string
}

func (p *MockAssyLine) Make(n int) <-chan factory.Vehicle {
	c := make(chan factory.Vehicle, n)

	go func() {
		for i := 1; i <= n; i++ {
			c <- &mockVehicle{
				b:  p.Brand,
				m:  p.Model,
				sn: i,
			}
		}
		close(c)
	}()

	return c
}

var _ factory.AssyLine = &MockAssyLine{}
