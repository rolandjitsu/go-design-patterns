# Command Patern

The [command](https://en.wikipedia.org/wiki/Command_pattern) pattern is a design pattern that encapsulates an action or request as an object. And it's commonly associated with terms like receiver, command and invoker.

Usually, the invoker doesn't know anything about the implementation details of the command or receiver, it just knows the command interface and its only responsibility is to invoke the command and do bookkeeping of commands.

The command is the object that knows about the receiver and its responsibility is to execute methods on the receiver.

```go
p := &Pinger{}
ping := &Ping{}
pingCmd = &PingCmd{ping}

p.Execute(pingCmd)

type Pinger struct {}

func (p *Pinger) Execute(c Cmd) {
    c.Execute()
    ...
}

type Cmd interface {
    Execute()
}

type Ping struct {}

func (p *Ping) Send() {
    ...
}

type PingCmd struct {
    Ping *Ping
}

func (pc *PingCmd) Execute() {
    pc.Ping.Send()
}
```

A perfect example of this pattern can be illustrated with the [undo](https://en.wikipedia.org/wiki/Undo) interaction in programs.

Check the code for examples.
