package iterator

func NewIntRange(start, end int) RangeIterator {
	return &intRange{start, end}
}

type RangeIterator interface {
	Next() (int, bool)
}

type intRange struct {
	start int
	end   int
}

func (r *intRange) Next() (v int, done bool) {
	r.start++
	v = r.start
	if r.start > r.end {
		done = true
		return
	}
	return
}
