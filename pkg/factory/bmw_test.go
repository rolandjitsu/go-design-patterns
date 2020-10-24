package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBmwAssyLine(t *testing.T) {
	bmwAssyLine := NewBmwAssyLine("M3")

	var cars []Vehicle
	for v := range bmwAssyLine.Make(4) {
		cars = append(cars, v)
	}

	assert.Len(t, cars, 4)
	for i, car := range cars {
		assert.Equal(t, "BMW", car.Brand())
		assert.Equal(t, "M3", car.Model())
		assert.Equal(t, i+1, car.SN())
	}
}

func BenchmarkBmwAssyLine(b *testing.B) {
	bmwAssyLine := NewBmwAssyLine("M3")
	for n := 0; n < b.N; n++ {
		<-bmwAssyLine.Make(1)
	}
}
