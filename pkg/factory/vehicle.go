package factory

type AssyLine interface {
	Make(n int) <-chan Vehicle
}

type Vehicle interface {
	Brand() string
	Model() string
	SN() int
}
