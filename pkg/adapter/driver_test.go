package adapter

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDriver(t *testing.T) {
	c := NewSpiConn()
	d := &Driver{Conn: c}

	r := make([]byte, 20)
	err := d.Tx([]byte("John"), r)
	if !assert.NoError(t, err) {
		return
	}

	r = bytes.TrimRight(r, "\x00")
	assert.Equal(t, "Hey John! from SPI", string(r))
}
