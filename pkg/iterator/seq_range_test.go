package iterator

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeqIntRange(t *testing.T) {
	var vals []int
	for n := range NewSeqIntRange(1, 5) {
		vals = append(vals, n)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5}, vals)
}

func TestSeqIntRangeEarlyBreak(t *testing.T) {
	var vals []int
	for n := range NewSeqIntRange(1, 10) {
		if n > 3 {
			break
		}
		vals = append(vals, n)
	}
	assert.Equal(t, []int{1, 2, 3}, vals)
}

func TestSeqIntRangeSlicesCollect(t *testing.T) {
	vals := slices.Collect(NewSeqIntRange(10, 14))
	assert.Equal(t, []int{10, 11, 12, 13, 14}, vals)
}
