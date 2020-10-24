package command

// NOTE: Not thread safe. Use a mutex if you need thread safety.
type ActionRunner struct {
	actions []Action
}

func (ar *ActionRunner) Execute(a Action) {
	ar.actions = append(ar.actions, a)
	a.Execute()
}

func (ar *ActionRunner) Undo() {
	sz := len(ar.actions)
	if sz > 0 {
		idx := sz - 1
		action := ar.actions[idx]
		ar.actions[idx] = nil
		ar.actions = ar.actions[:idx]
		action.Undo()
	}
}
