package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChanIntRange(t *testing.T) {
	r := NewChanIntRange(0, 4)

	var items []int
	for i := range r {
		items = append(items, i)
	}

	assert.Len(t, items, 4)
	assert.Equal(t, []int{1, 2, 3, 4}, items)
}

func BenchmarkChanIntRange(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var items []int
		r := NewChanIntRange(0, 4)
		for i := range r {
			// nolint: staticheck
			items = append(items, i)
		}
	}
}
