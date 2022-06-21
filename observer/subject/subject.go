package subject

import (
	"fmt"
	o "observer/observer"
	"reflect"
)

type Subject interface {
	RegisterObserver(ob o.Observer)
	RemoveObserver(ob o.Observer)
	NotifyObservers()
}

type subject struct {
	Observers []o.Observer
}

func NewSubject() *subject {
	return &subject{
		Observers: make([]o.Observer, 0),
	}
}

func (s *subject) RegisterObserver(ob o.Observer) {
	s.Observers = append(s.Observers, ob)
	fmt.Printf("Observer %s registered.\n", reflect.TypeOf(ob))
}

func (s *subject) RemoveObserver(ob o.Observer) {
	for idx, observer := range s.Observers {
		if observer == ob {
			s.Observers[idx] = s.Observers[len(s.Observers)-1] // Copy last element to index idx.
			s.Observers = s.Observers[:len(s.Observers)-1]     // Erase last element (order doesn't matter)
			fmt.Printf("Observer %s removed.\n", reflect.TypeOf(ob))
			break
		}
	}
}

func (s *subject) NotifyObservers() {}

func getObservers(s *subject) []o.Observer {
	return s.Observers
}
