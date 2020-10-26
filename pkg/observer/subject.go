package observer

import (
	"sync"
)

func NewSubject() Subject {
	return &subject{
		subs: map[int]Cb{},
	}
}

type Subject interface {
	Subscribe(Cb) Subscription
	Next(interface{})
}

type Cb func(interface{})

type Subscription interface {
	Unsubscribe()
}

type subject struct {
	mux  sync.RWMutex
	subs map[int]Cb
}

func (s *subject) Subscribe(cb Cb) Subscription {
	s.mux.Lock()
	defer s.mux.Unlock()
	id := len(s.subs) + 1
	s.subs[id] = cb
	return &subscription{
		id:      id,
		subject: s,
	}
}

func (s *subject) Next(data interface{}) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	for _, cb := range s.subs {
		cb(data)
	}
}

func (s *subject) removeSub(id int) {
	s.mux.Lock()
	defer s.mux.Unlock()
	delete(s.subs, id)
}

type subscription struct {
	id      int
	subject *subject
}

func (s *subscription) Unsubscribe() {
	s.subject.removeSub(s.id)
}

var _ Subject = &subject{}
var _ Subscription = &subscription{}
