# Singleton Pattern

The [singleton](https://en.wikipedia.org/wiki/Singleton_pattern) pattern is a design pattern that restricts the instantiation of a class (struct in Go) to a single instance.

This can be useful when you want to share an object across your program, e.g database connections, configs, etc.

```go
type Thanos struct {}

var thanos *Thanos
var once sync.Once

func GetThanos() *Thanos {
    once.Do(func() {
        thanos = &Thanos{}
    })
    return thanos
}
```

Check the code and tests for examples.
