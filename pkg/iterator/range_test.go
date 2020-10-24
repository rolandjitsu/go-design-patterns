package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntRange(t *testing.T) {
	r := NewIntRange(0, 4)

	var i int
	var done bool
	var items []int

	for {
		if i, done = r.Next(); done {
			break
		}
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
