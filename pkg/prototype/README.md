# Prototype Patern

The [prototype](https://en.wikipedia.org/wiki/Prototype_pattern) pattern is a design pattern that implies creating clones of objects instead of creating new instances.

This pattern can be used when:
* creating instances of an object is time and resource intensive
* making copies of an object is complex (deep copies, private members, etc)
* the object is exposed as an interface

For example:
```go
obj := NewObj(1)
copy := obj.Clone()

func NewObj(prop int) MyObj {
    return &myObj{PubProp: prop}
}

type MyObj interface {
    Clone() MyObj
}

type myObj struct {
    PubProp int
    privProp int
}

func (o *myObj) Clone() *MyObj {
    return &myObj{
        PubProp: o.Prop,
        privProp: o.privProp,
    }
}
```

Check the code for more examples.
