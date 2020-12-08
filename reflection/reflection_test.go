package reflection

import (
	"reflect"
	"testing"
)

type person struct {
	firstName  string
	lastName   string
	age        int
	occupation string
}

type walkSpy struct {
	fields map[string]interface{}
}

func (w *walkSpy) walked(s string) {
	w.fields[s] = nil
}

func TestWalk(t *testing.T) {
	person := person{
		firstName:  "Jeffrey",
		lastName:   "Hershals",
		age:        30,
		occupation: "Engineer",
	}
	spy := walkSpy{make(map[string]interface{})}

	Walk(person, spy.walked)

	want := map[string]interface{}{
		"firstName":  nil,
		"lastName":   nil,
		"occupation": nil,
	}

	if !reflect.DeepEqual(want, spy.fields) {
		t.Errorf("wanted calls %v got %v", want, spy.fields)
	}
}
