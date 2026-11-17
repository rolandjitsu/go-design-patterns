# Iterator Pattern

The [iterator](https://en.wikipedia.org/wiki/Iterator_pattern) pattern is a commonly used design pattern in software.
It entails that a collection must provide an iterator that can be used to iterate through its objects.

Traditionally in Go, iterators were implemented with custom structs or channels:
```go
c := NewIntRange(1, 10)

for {
    v, ok := c.Next()
    if !ok {
        break
    }
    // process v
}
```

## Go 1.23+ Range-Over-Func Iterators

Go 1.23 added native support for function iterators via `iter.Seq` and `iter.Seq2`. You can consume them with standard `for ... := range`:

```go
for v := range NewSeqIntRange(1, 10) {
    fmt.Println(v)
}
```

Check the code and tests for implementation examples and benchmarks.
