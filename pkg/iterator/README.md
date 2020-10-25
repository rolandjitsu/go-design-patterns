# Iterator Patern

The [iterator](https://en.wikipedia.org/wiki/Iterator_pattern) pattern is a commonly used design pattern in software.
It entails that a collection must provide an iterator that can be used to iterate through its objects.

To put it in simple terms:
```go
c := MakeIterator()

for c.Next() {
    v := c.Value()
    ...
}
```

One of the examples in this package illustrates the use of [channels](https://gobyexample.com/channels), which is a more common pattern in Go, but it doesn't exactly implement the iterator pattern and it's not meant for low latency algorithms (just run the benchmarks to see a comparison).

Check the tests for more implementation examples.
