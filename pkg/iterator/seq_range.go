package iterator

import "iter"

// NewSeqIntRange returns an iter.Seq[int] range-over-func iterator (Go 1.23+).
func NewSeqIntRange(start, end int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; i <= end; i++ {
			if !yield(i) {
				return
			}
		}
	}
}
