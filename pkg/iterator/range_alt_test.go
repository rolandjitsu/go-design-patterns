package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAltIntRange(t *testing.T) {
	r := NewAltIntRange(0, 4)

	var items []int

	for r.Next() {
		items = append(items, r.Value())
	}

	assert.Len(t, items, 4)
	assert.Equal(t, []int{1, 2, 3, 4}, items)
}

func BenchmarkAltIntRange(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var items []int
		r := NewAltIntRange(0, 4)
		for r.Next() {
			// nolint: staticheck
			items = append(items, r.Value())
		}
	}
}
