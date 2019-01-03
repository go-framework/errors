package errors

import (
	"fmt"
	"testing"
)

type errorMessage string

func (e errorMessage) Message() string {
	return string(e) + " message"
}

func (e errorMessage) Error() string {
	return string(e) + " detail"
}

func TestIntCodeError_New(t *testing.T) {
	e := &IntCodeError{}

	err := e.New("string")
	t.Log(err)

	e.Reset()
	err = e.New(fmt.Errorf("%s", "error"))
	t.Log(err)

	var errorMessage errorMessage = "ErrorMessage"

	e.Reset()
	err = e.New(errorMessage)
	t.Log(err)

	e2 := &IntCodeError{
		Level:   Level_Debug,
		Code:    -1,
		Message: "message",
		Detail:  "detail",
	}

	e.Reset()
	err = e.New(e2)
	t.Log(err)
}
