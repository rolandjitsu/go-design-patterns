# Adapter Patern

The [adapter](https://en.wikipedia.org/wiki/Adapter_pattern) pattern is a design pattern that allows the interface of a class to be used as another interface.

This is a frequently used pattern in ORM libraries ([ActiveRecord](https://guides.rubyonrails.org/active_record_basics.html), [GORM](https://gorm.io/index.html), etc), because it allows for connections and queries to different data backends (databases) while keeping the same client interface.

```go
bkd := &SQLBackend{}
db := &Conn{Backend: bkd}

db.Query()

...

type Conn struct {
    Backend Backend
}

func (c *Conn) Query() ([]byte, error) {
    return c.Backend.Query()
}

type Backend interface {
    Query() ([]byte, error)
}

type SQLBackend struct {}

func (db *SQLBackend) Query() ([]byte, error) {
    ...
}

type NoSQLBackend struct {}

func (db *NoSQLBackend) Query() ([]byte, error) {
    ...
}
```

Another common occurrence of this pattern is in hardware drivers. Take printers for instance.
Most printers can be used either via USB or serial connections or over the network.

Check the code for examples.
