package reflection

import (
	"reflect"
	"testing"
)

type walkSpy struct {
	fields map[string]interface{}
}

func (w *walkSpy) walked(s string) {
	w.fields[s] = nil
}

func TestWalk(t *testing.T) {
	x := struct {
		firstName  string
		age        int
		occupation string
	}{
		firstName:  "Jeffrey",
		age:        30,
		occupation: "Engineer",
	}

	spy := walkSpy{make(map[string]interface{})}

	Walk(x, spy.walked)

	want := map[string]interface{}{
		"firstName":  nil,
		"occupation": nil,
	}

	if !reflect.DeepEqual(want, spy.fields) {
		t.Errorf("wanted calls %v got %v", want, spy.fields)
	}
}
