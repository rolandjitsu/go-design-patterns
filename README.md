# Go Design Patterns
> A collection of design patterns implemented in Go

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/rolandjitsu/go-design-patterns/Test?label=tests&style=flat-square)](https://github.com/rolandjitsu/go-design-patterns/actions?query=workflow%3ATest)

## Design Patterns
* Creational
    * [Factory](./pkg/factory/README.md)
    * [Prototype](./pkg/prototype/README.md)
    * [Singleton](./pkg/singleton/README.md)
* Structural
    * [Adapter](./pkg/adapter/README.md)
    * [Proxy](./pkg/proxy/README.md)
* Behavioural
    * [Command](./pkg/command/README.md)
    * [Iterator](./pkg/iterator/README.md)
    * [Observer](./pkg/observer/README.md)

## Test
To run tests for `pkg` run:
```bash
go test ./pkg/...
```

To avoid caching during tests use:
```bash
go test -count=1 ./pkg/...
```

To run benchmark tests use the `-bench` flag:
```bash
go test -bench=. ./pkg/...
```

To get coverage reports use the `-cover` flag:
```bash
go test -coverprofile=coverage.out ./pkg/...
```

And to view the profile run:
```bash
go tool cover -html=coverage.out
```
