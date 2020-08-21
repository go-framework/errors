package errors

import (
	"fmt"
	"reflect"
)

type Error interface {
	error
	Is(err error) bool
	As(target interface{}) bool
	Wrap(err error) error
	Unwrap() error
}

func New(any interface{}) Error {
	switch e := any.(type) {
	case string:
		return NewMessage(e)
	case int, int8, int16, int32, int64:
		return NewCode(iToInt64(e))
	case uint, uint8, uint16, uint32, uint64:
		return NewCode(iToUint64(e))
	case *errorCode:
		return NewErrorCode(e.code, e.message, e.error)
	case ErrorCode:
		return NewErrorCode(e.GetCode(), e.GetMessage(), e)
	case error:
		return NewError(e)
	default:
		switch reflect.TypeOf(any).Kind() {
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
			return NewCode(iToInt64(any))
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
			return NewCode(iToUint64(any))
		case reflect.String:
			return NewMessage(reflect.ValueOf(any).String())
		}
	}
	return NewMessage(fmt.Sprintf("%+v", any))
}
