package errors

import (
	"errors"
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestAppend(t *testing.T) {
	var errs error
	t.Log(errs)

	errs = Append(errs, fmt.Errorf("error 1"), fmt.Errorf("error 2"))
	t.Log(errs)

	errs = Append(errs, fmt.Errorf("error 3"), fmt.Errorf("error 4"))
	t.Log(errs)

	errs = fmt.Errorf("error 0")

	t.Log(errs)
	errs = Append(errs, fmt.Errorf("error 1"), fmt.Errorf("error 2"))
	t.Log(errs)

	errs = Append(errs, fmt.Errorf("error 3"), fmt.Errorf("error 4"))
	t.Log(errs)

	var errs2 error
	t.Log(errs2)

	errs2 = Append(errs2, errs)

	t.Log(errs2)

	errs2 = Append(errs2, errs)

	t.Log(errs2)
}

func TestErrors(t *testing.T) {
	var errs Errors

	t.Log("Errors", errs)

	//
	// Errors
	//

	errs.Append(nil)
	errs.Append(errors.New("go errors error"))
	errs.Append(NewTextError("text error"))
	errs.Append(NewTextError("text error \n separator"))
	errs.Append(NewTextError("text error \t separator"))
	errs.Append(NewTextError("text error \n\t separator"))
	errs.Append(NewTextError("text error \t\n separator"))

	var intCode IntCode = -1
	errs.Append(intCode.WithDetail("IntCode"))

	t.Log("Errors", errs)

	//
	// error
	//
	var err error = errs

	t.Log("error", err)
}

func TestJSON(t *testing.T) {
	var errs Errors

	//
	// Errors
	//

	errs.Append(nil)
	errs.Append(errors.New("go errors error"))
	errs.Append(NewTextError("text error"))
	errs.Append(NewTextError("text error \n separator"))
	errs.Append(NewTextError("text error \t separator"))
	errs.Append(NewTextError("text error \n\t separator"))
	errs.Append(NewTextError("text error \t\n separator"))

	var intCode IntCode = -1
	errs.Append(intCode.WithDetail("IntCode"))

	t.Log("Errors", errs)

	str, err := jsoniter.MarshalToString(errs)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(str)
}
