# Proxy Patern

The [proxy](https://en.wikipedia.org/wiki/Proxy_pattern) pattern is a design pattern in which a class (**proxy**) acts as an interface to something else.
The proxy could be an interface to anything: a network connection, another class, a file, etc.

A proxy can be useful in a variety of situations:
* A frontend for load balancing
* Hide private infrastructure
* Caching layer
* etc

A good example of a proxy that can be used as a load balancer (and other purposes) is [nginx](https://www.nginx.com/).

To illustrate how to use this pattern in Go, let's implement a simple caching layer for files to speed up multiple reads for the same files:
```go
package fileproxy

import (
	"io"
	"io/ioutil"
	"sync"
)

var mux sync.RWMutex
var cache map[string][]byte

func init() {
	cache = map[string][]byte{}
}

func NewFileProxy(name string) FileProxy {
	return &fileProxy{Name: name}
}

type FileProxy interface {
	io.Reader
}

type fileProxy struct {
	Name string
}

func (fp *fileProxy) Read(b []byte) (n int, err error) {
	mux.RLock()
	if data, ok := cache[fp.Name]; ok {
		copy(b, data)
		mux.RUnlock()
		n = len(b)
		return
	}
	mux.RUnlock()

	var data []byte
	data, err = ioutil.ReadFile(fp.Name)
	if err == nil {
		copy(b, data)
		mux.Lock()
		n = len(b)
		cache[fp.Name] = data
		mux.Unlock()
	}

	return
}

var _ FileProxy = &fileProxy{}
```

If we run a simple benchmark to compare how caching improves read operations we can see a pretty significant improvement over running `ioutil.ReadFile()` alone:
```bash
❯ go test -bench=. ./pkg/proxy/...
goos: darwin
goarch: amd64
pkg: github.com/rolandjitsu/go-design-patterns/pkg/proxy
BenchmarkReadFile-8        87010             13601 ns/op
BenchmarkFileProxy-8    52101362                21.9 ns/op
```

Check the code for examples.
