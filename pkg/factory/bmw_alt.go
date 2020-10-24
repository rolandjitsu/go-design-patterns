package factory

func NewAltBmwAssyLine(model string) func() Vehicle {
	// NOTE: Not thread safe. Use a mutex if you need thread safety.
	var c int
	return func() Vehicle {
		c++
		return &bmw{
			model: model,
			sn:    c,
		}
	}
}
