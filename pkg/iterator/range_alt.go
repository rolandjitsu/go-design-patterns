package iterator

func NewAltIntRange(start, end int) AltRangeIterator {
	return &altIntRange{
		end: end,
		v:   start,
	}
}

type AltRangeIterator interface {
	Next() bool
	Value() int
}

type altIntRange struct {
	end int
	v   int
}

func (r *altIntRange) Next() (ok bool) {
	if r.v < r.end {
		ok = true
		r.v++
	}
	return
}

func (r *altIntRange) Value() int {
	return r.v
}
