# Iterator Patern

The [iterator](https://www.geeksforgeeks.org/iterator-pattern/) pattern is a commonly used design pattern in software.
It entails that a collection must provide an iterator that can be used to iterate through its objects.

To put it in simple terms:
```go
c := MakeIterator()

for c.Next() {
    v := c.Value()
    ...
}

```

Check the tests for a few implementation examples.

One of the examples illustrates using [channels](https://gobyexample.com/channels), which is a more common pattern in Go, but doesn't exactly implement the iterator pattern and it's not meant for low latency algorithms (just run the benchmarks to see a comparison).
