package errors

import (
	"fmt"
	"reflect"
)

// Get detail from any interface.
func GetDetail(any interface{}) string {
	// implement Error interface.
	if e, ok := any.(Error); ok {
		return e.GetDetail()
	} else if e, ok := any.(error); ok {
		return e.Error()
	} else if str, ok := any.(string); ok {
		return str
	}

	if any == nil{
		return ""
	}

	return fmt.Sprintf("%v", any)
}

// select handler Error interface.
func selectErrorHandler(code interface{}) Error {
	switch code.(type) {
	case string, StringErrCode:
		return new(StringCodeError)
	case uint8, uint16, uint32, uint64, uint, UintErrCode:
		return new(UintCodeError)
	case int8, int16, int32, int64, int, IntCode:
		return new(IntCodeError)
	}

	switch reflect.TypeOf(code).Kind() {
	case reflect.String:
		return new(StringCodeError)
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		return new(UintCodeError)
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return new(IntCodeError)
	}

	return new(StringCodeError)
}

// Return nil if err is nil.
func Nil(err error) error {
	switch t := err.(type) {
	case nil:
		return nil
	case Errors:
		if len(t) == 0 {
			return nil
		}
		return t
	default:
		return err
	}
}
