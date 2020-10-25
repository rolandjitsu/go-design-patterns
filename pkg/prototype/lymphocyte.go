package prototype

func NewLymphocyte(t LymphocyteType, p []Protein) Lymphocyte {
	return &lymphocyte{
		t: t,
		p: p,
	}
}

type LymphocyteType int

const (
	B LymphocyteType = iota
	T
)

type Lymphocyte interface {
	Proteins() []Protein
	Type() LymphocyteType
	Clone() Lymphocyte
}

type lymphocyte struct {
	p []Protein
	t LymphocyteType
}

func (c *lymphocyte) Proteins() []Protein {
	return c.p
}

func (c *lymphocyte) Type() LymphocyteType {
	return c.t
}

func (c *lymphocyte) Clone() Lymphocyte {
	return &lymphocyte{
		t: c.t,
		p: cloneProteins(c.p),
	}
}
