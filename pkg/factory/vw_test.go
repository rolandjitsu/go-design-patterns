package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVwAssyLine(t *testing.T) {
	vwAssyLine := NewVwAssyLine("Beetle")

	var cars []Vehicle
	for v := range vwAssyLine.Make(4) {
		cars = append(cars, v)
	}

	assert.Len(t, cars, 4)
	for i, car := range cars {
		assert.Equal(t, "Volkswagen", car.Brand())
		assert.Equal(t, "Beetle", car.Model())
		assert.Equal(t, i+1, car.SN())
	}
}
