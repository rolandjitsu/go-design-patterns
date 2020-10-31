package proxy

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
