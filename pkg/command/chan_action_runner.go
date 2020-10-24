package command

func NewChanActionRunner() *ChanActionRunner {
	c := make(chan Action)
	ar := &ChanActionRunner{A: c}

	go func() {
		for a := range c {
			ar.execute(a)
		}
	}()

	return ar
}

type ChanActionRunner struct {
	A       chan<- Action
	actions []Action
}

func (ar *ChanActionRunner) execute(a Action) {
	ar.actions = append(ar.actions, a)
	a.Execute()
}

func (ar *ChanActionRunner) Close() {
	close(ar.A)
}

// NOTE: Not thread safe. Use a mutex if you need thread safety.
func (ar *ChanActionRunner) Undo() {
	sz := len(ar.actions)
	if sz > 0 {
		idx := sz - 1
		action := ar.actions[idx]
		ar.actions[idx] = nil
		ar.actions = ar.actions[:idx]
		action.Undo()
	}
}
