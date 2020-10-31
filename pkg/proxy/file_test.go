package proxy

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testdata = "testdata/avengers.txt"

func TestFileProxy(t *testing.T) {
	fp := NewFileProxy(testdata)

	data := make([]byte, 8)
	n, err := fp.Read(data)
	assert.NoError(t, err)
	assert.Equal(t, 8, n)
	assert.Equal(t, "Iron Man", string(data))

	data = make([]byte, 24)
	n, err = fp.Read(data)
	assert.NoError(t, err)
	assert.Equal(t, 24, n)
	assert.Equal(t, "Iron Man, Ant-Man, Hulk\n", string(data))
}

func BenchmarkReadFile(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := ioutil.ReadFile(testdata)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFileProxy(b *testing.B) {
	data := make([]byte, 24)
	fp := NewFileProxy(testdata)

	for n := 0; n < b.N; n++ {
		_, err := fp.Read(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}
