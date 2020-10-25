package prototype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloneLymphocyte(t *testing.T) {
	tLymphocyte := NewLymphocyte(T, []Protein{
		{Name: "CD4"},
		{Name: "CD28"},
	})

	tc := tLymphocyte.Clone()

	assert.Equal(t, T, tc.Type())
	assert.Equal(t, []Protein{
		{Name: "CD4"},
		{Name: "CD28"},
	}, tc.Proteins())
}
