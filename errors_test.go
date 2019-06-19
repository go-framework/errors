package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestAppend(t *testing.T) {
	var err error
	t.Log(err)

	err = Append(err, fmt.Errorf("error 1"), fmt.Errorf("error 2"))
	t.Log(err)

	err = Append(err, fmt.Errorf("error 3"), fmt.Errorf("error 4"))
	t.Log(err)

	err = fmt.Errorf("error 0")

	t.Log(err)
	err = Append(err, fmt.Errorf("error 1"), fmt.Errorf("error 2"))
	t.Log(err)

	err = Append(err, fmt.Errorf("error 3"), fmt.Errorf("error 4"))
	t.Log(err)
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
