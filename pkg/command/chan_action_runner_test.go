package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChanActionRunner(t *testing.T) {
	ar := NewChanActionRunner()            // Invoker
	counter := Counter{}                   // Receiver
	action := NewIncrementAction(&counter) // Action

	ar.A <- action
	assert.Equal(t, 1, counter.Value)

	ar.Undo()
	assert.Equal(t, 0, counter.Value)
}

func BenchmarkChanActionRunner(b *testing.B) {
	ar := NewChanActionRunner()
	counter := Counter{}
	for n := 0; n < b.N; n++ {
		action := NewIncrementAction(&counter)
		ar.A <- action
	}
}
