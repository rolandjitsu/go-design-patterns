package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndgame(t *testing.T) {
	err := RunEndgame()
	assert.NoError(t, err)
}
