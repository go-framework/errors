package errors

import (
	"fmt"
)

// get detail from any interface.
func getDetail(any interface{}) string {
	// implement Error interface.
	if e, ok := any.(Error); ok {
		return e.GetDetail()
	} else if e, ok := any.(error); ok {
		return e.Error()
	} else if str, ok := any.(string); ok {
		return str
	}

	return fmt.Sprintf("%v", any)
}
