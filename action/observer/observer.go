package observer

import "fmt"

type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify(msg string)
}

type IObserver interface {
	Update(msg string)
}

type Subject struct {
	observers []IObserver
}

func (s *Subject) Register(observer IObserver) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Remove(observer IObserver) {
	for key, value := range s.observers {
		if observer == value {
			s.observers = append(s.observers[:key], s.observers[key+1:]...)
		}
	}
}

func (s *Subject) Notify(msg string) {
	for _, o := range s.observers {
		o.Update(msg)
	}
}

type Observer struct {
	data string
}

func (o *Observer) Update(msg string) {
	fmt.Println(o.data + ":" + msg)
}
