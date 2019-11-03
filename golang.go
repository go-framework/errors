package errors

import (
	"errors"
	"reflect"
)

// golang error type
var errorType = errors.New("")

// Check is golang sdk error
func IsSDKError(err error) bool {
	return reflect.TypeOf(err) == reflect.TypeOf(errorType)
}
