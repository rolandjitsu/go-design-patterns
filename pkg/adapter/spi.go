package adapter

import (
	"bytes"
	"fmt"
)

func NewSpiConn() Conn {
	return &spiConn{}
}

type spiConn struct{}

func (s *spiConn) Tx(w []byte, r []byte) (err error) {
	var buf bytes.Buffer
	_, err = fmt.Fprintf(&buf, "Hey %s! from SPI", string(w))
	if err != nil {
		return
	}

	data := buf.Bytes()
	sz := len(data)

	for i := range r {
		if i < sz {
			r[i] = data[i]
		} else {
			break
		}
	}

	return
}
