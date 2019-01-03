package errors

import "testing"

func TestNewCaller(t *testing.T) {
	caller := NewCaller(2, 64, true)
	t.Log(caller)
	for _, stack := range caller.Stacks {
		t.Log(stack)
	}
}
