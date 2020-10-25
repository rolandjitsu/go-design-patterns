# Singleton Patern

The [singleton](https://en.wikipedia.org/wiki/Singleton_pattern) pattern is a design pattern that restricts the instantiation of a class to one object.
What that means is simply that you have to share the same object throughout the program.

The simplest example of this is a DB connection:
```go
// main.go
import (
    "mydb"
    "stuff"
)

func main() {
    c := &mydb.Conn{}
    stuff := stuff.GetTheStuff(c)
    ...
}

// mydb/conn.go
package mydb
type Conn struct {}

func (c *Conn) Query(q string) {
    ...
}

// stuff/stuff.go
package stuff

func GetTheStuff(c *Conn) {
    c.Query()
}
```

Check the code for more examples.
