package factorytest

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rolandjitsu/go-design-patterns/pkg/factory"
)

func TestMockAssyLine(t *testing.T) {
	assyLine := &MockAssyLine{
		Brand: "MyBrand",
		Model: "MyModel",
	}

	var cars []factory.Vehicle
	for c := range assyLine.Make(4) {
		cars = append(cars, c)
	}

	assert.Len(t, cars, 4)

	for i, c := range cars {
		assert.Equal(t, "MyBrand", c.Brand())
		assert.Equal(t, "MyModel", c.Model())
		assert.Equal(t, i+1, c.SN())
	}
}
