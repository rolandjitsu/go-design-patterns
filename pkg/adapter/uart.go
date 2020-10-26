package adapter

import (
	"bytes"
	"fmt"
)

func NewUartConn() Conn {
	return &uartConn{}
}

type uartConn struct{}

func (s *uartConn) Tx(w []byte, r []byte) (err error) {
	var buf bytes.Buffer
	_, err = fmt.Fprintf(&buf, "Hey %s! from UART", string(w))
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
