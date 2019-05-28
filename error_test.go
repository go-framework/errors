package errors

import (
	"reflect"
	"testing"
)

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

func TestNewCode_int(t *testing.T) {
	DebugEnable()
	err := NewCode(1, "detail")
	t.Log(reflect.TypeOf(err))
	t.Log(err)
}

func TestNewCode_uint(t *testing.T) {
	DebugEnable()
	err := NewCode(uint(1), "detail")
	t.Log(reflect.TypeOf(err))
	t.Log(err)
}

func TestNewCode_defined(t *testing.T) {
	type Code int
	var code Code = 1
	DebugEnable()
	err := NewCode(code, "detail")
	t.Log(reflect.TypeOf(err))
	t.Log(err)
}

func TestNewError(t *testing.T) {
	err := NewError("string code", "message", "detail", Level_Debug)
	t.Log(err)
	DebugEnable()
	err = NewError("string code", "message", "detail", Level_Debug)
	t.Log(err)
}
