package command

type Action interface {
	Execute()
	Undo()
}
