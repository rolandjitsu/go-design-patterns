package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAltBmwAssyLine(t *testing.T) {
	makeBmw := NewAltBmwAssyLine("M3")

	cars := []Vehicle{
		makeBmw(),
		makeBmw(),
		makeBmw(),
	}

	for i, car := range cars {
		assert.Equal(t, "BMW", car.Brand())
		assert.Equal(t, "M3", car.Model())
		assert.Equal(t, i+1, car.SN())
	}
}

func BenchmarkAltBmwAssyLine(b *testing.B) {
	makeBmw := NewAltBmwAssyLine("M3")
	for n := 0; n < b.N; n++ {
		makeBmw()
	}
}
