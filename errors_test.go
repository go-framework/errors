package errors

import (
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
