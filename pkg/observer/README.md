# Observer Patern

The [observer](https://en.wikipedia.org/wiki/Observer_pattern) pattern is a design pattern in which an object (a **subject**) keeps track of all of its dependents (**observers**) and notifies them of any state changes.

In Go, the closest example of this pattern are the builtin [channels](https://tour.golang.org/concurrency/2) and the use of [goroutines](https://gobyexample.com/goroutines):
```go
sub := make(chan interface{})

go func(c <-chan interface{}>) {
    for data := range c {
        fmt.Println(data)
    }
}(sub)

sub <- "Hey there"
```

Though, for multiple observers to be notified, you need to send the message once for each observer.

Another example is the [rxjs](https://www.learnrxjs.io) lib and accompanying libs for other languages.

And to illustrate how an implementation of this pattern looks like, check the following example:
```go
sub := NewSubject()

sub1 := sub.Subscribe(func (data interface{}) {
    fmt.Println("Sub 1 says", data)
})
sub2 := sub.Subscribe(func (data interface{}) {
    fmt.Println("Sub 2 says", data)
})

sub.Next("Hey there")

sub2.Unsubscribe()

sub.Next("Hey again")

func NewSubject() Subject {
    // ...
}

type Subject interface {
    Subscribe(func(interface{})) Subscription
    Next(interface{})
}

type Subscription interface {
    Unsubscribe()
}
```

Check the code for examples.
