package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntRange(t *testing.T) {
	r := NewIntRange(0, 4)

	var items []int

	for i, done := r.Next(); !done; i, done = r.Next() {
		items = append(items, i)
	}

	assert.Len(t, items, 4)
	assert.Equal(t, []int{1, 2, 3, 4}, items)
}

func BenchmarkIntRange(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var items []int
		r := NewIntRange(0, 4)
		for i, done := r.Next(); !done; i, done = r.Next() {
			// nolint: staticheck
			items = append(items, i)
		}
	}
}
