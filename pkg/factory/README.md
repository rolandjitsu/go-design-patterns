# Factory Patern

The [factory](https://en.wikipedia.org/wiki/Factory_method_pattern) pattern is a design pattern used to create different types of objects using the same interface without exposing how the object is created.

In Go, this can be achieved with the `interface` type:
```go
type MyObj interface {
    String() string
}

func NewObj1() MyObj {
    return &obj1{}
}

type obj1 struct {}

func (o *obj1) String() string {
    return "OBJ1"
}

func NewObj2() MyObj {
    return &obj2{}
}

type obj2 struct {}

func (o *obj2) String() string {
    return "OBJ2"
}
```

This can be very useful for unit testing as you can create mocks for a specific interface:
```go
type MyObj interface {
    Read() []byte
}

func NewMockMyObj(data []byte) MyObj {
    return &mockMyObj{data}
}

type mockMyObj struct {
    data []byte
}

func (o *mockMyObj) Read() []byte {
    return o.data
}
```

Check the tests for more examples.
