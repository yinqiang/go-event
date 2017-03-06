package event

import (
	"container/list"
	"errors"
	"sync"
)

var (
	ErrEventNameEmpty       = errors.New("event name is empty")
	ErrEventNotFound        = errors.New("event not found")
	ErrListenerIsNil        = errors.New("listener is nil")
	ErrListenerNotFound     = errors.New("listener not found")
	ErrListenerAlreadyAdded = errors.New("listener already added")
)

type Event struct {
	sync.Mutex
	events map[string]*list.List
}

type EventListener *func(interface{})

func NewEventManager() *Event {
	return &Event{
		events: make(map[string]*list.List),
	}
}

func (e *Event) AddEventListener(name string, listener EventListener) error {
	if len(name) == 0 {
		return ErrEventNameEmpty
	}
	e.Lock()
	defer e.Unlock()
	lst, exist := e.events[name]
	if !exist {
		lst = list.New()
		e.events[name] = lst
	} else {
		for lt := lst.Front(); lt != nil; lt = lt.Next() {
			if lt.Value == listener {
				return ErrListenerAlreadyAdded
			}
		}
	}
	lst.PushBack(listener)
	return nil
}

func (e *Event) RemoveEventListener(name string, listener EventListener) error {
	if len(name) == 0 {
		return ErrEventNameEmpty
	} else if listener == nil {
		return ErrListenerIsNil
	}
	e.Lock()
	defer e.Unlock()
	lst, exist := e.events[name]
	if !exist {
		return ErrEventNotFound
	}
	for lt := lst.Front(); lt != nil; lt = lt.Next() {
		if lt.Value == listener {
			lst.Remove(lt)
			return nil
		}
	}
	return ErrListenerNotFound
}

func (e *Event) RemoveAllListeners(name string) error {
	if len(name) == 0 {
		return ErrEventNameEmpty
	}
	e.Lock()
	defer e.Unlock()
	_, exist := e.events[name]
	if !exist {
		return ErrEventNotFound
	}
	delete(e.events, name)
	return nil
}

func (e *Event) DispatchEvent(name string, data interface{}) error {
	if len(name) == 0 {
		return ErrEventNameEmpty
	}
	e.Lock()
	defer e.Unlock()
	lst, exist := e.events[name]
	if !exist {
		return ErrEventNotFound
	}
	for c := lst.Front(); c != nil; c = c.Next() {
		f := c.Value.(EventListener)
		go (*f)(data)
	}
	return nil
}
