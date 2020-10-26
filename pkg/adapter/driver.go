package adapter

type Driver struct {
	Conn Conn
}

func (d *Driver) Tx(w []byte, r []byte) error {
	return d.Conn.Tx(w, r)
}

type Conn interface {
	Tx([]byte, []byte) error
}
