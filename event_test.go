package event

import (
	"testing"
)

func Test_NewEventManager(t *testing.T) {
	e := NewEventManager()
	if e == nil {
		t.Fail()
	}
}

func Test_AddEventListener(t *testing.T) {
	e := NewEventManager()
	f1 := func(i interface{}) {}
	f2 := func(i interface{}) {}
	if err := e.AddEventListener("e1", &f1); err != nil {
		t.Fatal(err)
	}
	if err := e.AddEventListener("e2", &f2); err != nil {
		t.Fatal(err)
	}
	for i := 1; i < 10; i++ {
		f := func(i interface{}) {}
		if err := e.AddEventListener("e1", &f); err != nil {
			t.Fatal(err)
		}
	}
	if err := e.AddEventListener("e2", &f2); err != ErrListenerAlreadyAdded {
		t.Fatal("fatal on add twice")
	}
	if err := e.AddEventListener("e2", &f1); err != nil {
		t.Fatal("fatal on mult add")
	}
}

func Test_RemoveEventListener(t *testing.T) {
	e := NewEventManager()
	f1 := func(i interface{}) {}
	e.AddEventListener("e1", &f1)
	if err := e.RemoveEventListener("e1", &f1); err != nil {
		t.Fatal(err)
	}
	e.AddEventListener("e1", &f1)
	e.AddEventListener("e2", &f1)
	if err := e.RemoveEventListener("e2", &f1); err != nil {
		t.Fatal(err)
	}
	f2 := func(i interface{}) {}
	f3 := func(i interface{}) {}
	f4 := func(i interface{}) {}
	f5 := func(i interface{}) {}
	e.AddEventListener("e1", &f2)
	e.AddEventListener("e1", &f3)
	e.AddEventListener("e1", &f4)
	e.AddEventListener("e1", &f5)
	if err := e.RemoveEventListener("e1", &f3); err != nil {
		t.Fatal(err)
	}
}

func Test_RemoveAllListeners(t *testing.T) {
	e := NewEventManager()
	f1 := func(i interface{}) {}
	f2 := func(i interface{}) {}
	f3 := func(i interface{}) {}
	f4 := func(i interface{}) {}
	f5 := func(i interface{}) {}
	e.AddEventListener("e", &f1)
	e.AddEventListener("e", &f2)
	e.AddEventListener("e", &f3)
	e.AddEventListener("e", &f4)
	e.AddEventListener("e", &f5)
	if err := e.RemoveAllListeners("e"); err != nil {
		t.Fatal(err)
	}
}

func Test_DispatchEvent(t *testing.T) {
	e := NewEventManager()
	f1 := func(i interface{}) {
		if i.(bool) == false {
			t.Fail()
		}
	}
	f2 := func(i interface{}) {
		if i.(bool) == false {
			t.Fail()
		}
	}
	f3 := func(i interface{}) {
		if i.(bool) == false {
			t.Fail()
		}
	}
	f4 := func(i interface{}) {
		if i.(bool) == false {
			t.Fail()
		}
	}
	f5 := func(i interface{}) {
		if i.(bool) == false {
			t.Fail()
		}
	}
	e.AddEventListener("e", &f1)
	e.AddEventListener("e", &f2)
	e.AddEventListener("e", &f3)
	e.AddEventListener("e", &f4)
	e.AddEventListener("e", &f5)
	if err := e.DispatchEvent("e", true); err != nil {
		t.Fatal(err)
	}
}
