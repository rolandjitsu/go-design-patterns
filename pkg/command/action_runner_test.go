package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestActionRunner(t *testing.T) {
	ar := &ActionRunner{}                  // Invoker
	counter := Counter{}                   // Receiver
	action := NewIncrementAction(&counter) // Action

	ar.Execute(action)
	assert.Equal(t, 1, counter.Value)

	ar.Undo()
	assert.Equal(t, 0, counter.Value)
}

func BenchmarkActionRunner(b *testing.B) {
	ar := &ActionRunner{}
	counter := Counter{}
	for n := 0; n < b.N; n++ {
		action := NewIncrementAction(&counter)
		ar.Execute(action)
	}
}
