package require

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
	Equal(t, paul, paul, "Paul equals paul")
}

func TestNotEqual(t *testing.T) {
	NotEqual(t, paul, peter, "Paul does not equal peter")
}

func TestNotEqualNil(t *testing.T) {
	NotEqual(t, paul, nil, "Paul is not nil")
}

func TestNilEqualNil(t *testing.T) {
	Equal(t, nil, nil, "Nil is nil")
}

type DummyT struct{}

func (t *DummyT) Fatalf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
func TestFailingAssert(t *testing.T) {
	dummyT := &DummyT{}
	require(dummyT, 1, 2, "Numbers are not equal", true)
}
