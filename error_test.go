package errors

import "testing"

func TestNew(t *testing.T) {
	err := New("string")
	t.Log(err)
	DebugEnable()
	err = New("string")
	t.Log(err)
}

func TestNewCode(t *testing.T) {
	err := NewCode("string code", "detail")
	t.Log(err)
	DebugEnable()
	err = NewCode("string code", "detail")
	t.Log(err)
}

func TestNewError(t *testing.T) {
	err := NewError("string code", "message", "detail", Level_Debug)
	t.Log(err)
	DebugEnable()
	err = NewError("string code", "message", "detail", Level_Debug)
	t.Log(err)
}
