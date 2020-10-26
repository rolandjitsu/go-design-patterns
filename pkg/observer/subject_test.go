package observer

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSubject(t *testing.T) {
	sub := NewSubject()

	var mux sync.Mutex
	var greetings []string

	sub1 := sub.Subscribe(func(data interface{}) {
		mux.Lock()
		greetings = append(greetings, data.(string))
		mux.Unlock()
	})

	sub2 := sub.Subscribe(func(data interface{}) {
		mux.Lock()
		greetings = append(greetings, data.(string))
		mux.Unlock()
	})

	sub.Next("Hey!")
	assert.Len(t, greetings, 2)
	assert.Equal(t, "Hey!", greetings[0])
	assert.Equal(t, "Hey!", greetings[1])

	sub1.Unsubscribe()
	sub2.Unsubscribe()
	sub.Next("Hey!")

	assert.Len(t, greetings, 2)
}

func TestChanSubject(t *testing.T) {
	sub := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(2)

	var mux sync.Mutex
	var greetings []string

	go func(c <-chan interface{}) {
		for data := range c {
			time.Sleep(time.Millisecond)
			mux.Lock()
			greetings = append(greetings, data.(string))
			mux.Unlock()
			wg.Done()
		}
	}(sub)

	go func(c <-chan interface{}) {
		for data := range c {
			time.Sleep(time.Millisecond)
			mux.Lock()
			greetings = append(greetings, data.(string))
			mux.Unlock()
			wg.Done()
		}
	}(sub)

	// NOTE: We need to send the same message twice
	// so that each goroutine receives it
	sub <- "Hey!"
	sub <- "Hey!"

	close(sub)
	wg.Wait()

	assert.Len(t, greetings, 2)
	assert.Equal(t, "Hey!", greetings[0])
	assert.Equal(t, "Hey!", greetings[1])
}
