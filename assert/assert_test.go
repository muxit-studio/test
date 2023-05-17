package assert

import (
	"fmt"
	"testing"
)

type Person struct {
	name string
	age  int
}

var paul = Person{
	name: "Paul",
	age:  32,
}

var peter = Person{
	name: "Peter",
	age:  21,
}

func TestEqual(t *testing.T) {
	ok := Equal(t, paul, paul, "Paul equals paul")
	if !ok {
		t.Errorf("Equal should return true")
	}
}

func TestNotEqual(t *testing.T) {
	ok := NotEqual(t, paul, peter, "Paul does not equal peter")
	if !ok {
		t.Errorf("NotEqual should return true")
	}
}

func TestNotEqualNil(t *testing.T) {
	ok := NotEqual(t, paul, nil, "Paul is not nil")
	if !ok {
		t.Errorf("NotEqual should return true")
	}
}

func TestNilEqualNil(t *testing.T) {
	ok := Equal(t, nil, nil, "Nil is nil")
	if !ok {
		t.Errorf("Equal should return true")
	}
}

type DummyT struct {
}
func (t *DummyT) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
func TestFailingAssert(t *testing.T) {
	dummyT := &DummyT{}
	_ = assert(dummyT, 1, 2, "Numbers are not equal", true)
}
