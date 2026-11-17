# Proxy Pattern

The [proxy](https://en.wikipedia.org/wiki/Proxy_pattern) pattern is a design pattern in which an object acts as an interface to something else. The proxy could be an interface to anything: a network connection, another class, a file, etc.

A proxy can be useful in a variety of situations:
* A frontend for load balancing
* Hide private infrastructure
* Caching layer
* etc

A good example of a proxy that can be used as a load balancer (and other purposes) is [nginx](https://www.nginx.com/) or the Go standard library [net/http/httputil](https://pkg.go.dev/net/http/httputil) (`ReverseProxy`).

In this package, we implement a simple caching layer for files:
```go
package proxy

import (
	"io"
	"os"
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
	data, err = os.ReadFile(fp.Name)
	if err == nil {
		copy(b, data)
		mux.Lock()
		n = len(b)
		cache[fp.Name] = data
		mux.Unlock()
	}

	return
}
```

If we run a simple benchmark to compare how caching improves read operations we can see a significant improvement over running `os.ReadFile()` alone:
```text
BenchmarkReadFile-8        87010             13601 ns/op
BenchmarkFileProxy-8    52101362                21.9 ns/op
```

Check the code and tests for details.
