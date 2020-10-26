package adapter

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUartConn(t *testing.T) {
	c := NewUartConn()

	r := make([]byte, 20)
	err := c.Tx([]byte("John"), r)
	if !assert.NoError(t, err) {
		return
	}

	r = bytes.TrimRight(r, "\x00")
	assert.Equal(t, "Hey John! from UART", string(r))
}
