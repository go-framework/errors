package errors

import (
	"fmt"
	"math"
	"reflect"
)

type Error interface {
	error
	Is(err error) bool
	As(target interface{}) bool
	Wrap(err error) error
	Unwrap() error
	SetError(err error) error
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
		err, _ := any.(error)
		return NewErrorCode(e.GetCode(), e.GetMessage(), err)
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

func GetIntCode(err error) int64 {
	switch e := err.(type) {
	case *errorCode:
		return iToInt64(e.code)
	case ErrorCode:
		return iToInt64(e.GetCode())
	default:
		return math.MaxInt64
	}
}

func GetUintCode(err error) uint64 {
	switch e := err.(type) {
	case *errorCode:
		return iToUint64(e.code)
	case ErrorCode:
		return iToUint64(e.GetCode())
	default:
		return math.MaxUint64
	}
}

func GetStringCode(err error) string {
	switch e := err.(type) {
	case *errorCode:
		str, ok := e.code.(string)
		if ok {
			return str
		}
	case ErrorCode:
		str := iToString(e.GetCode())
		if len(str) > 0 {
			return str
		}
	}
	return err.Error()
}

func GetMessage(err error) string {
	switch e := err.(type) {
	case *errorCode:
		return e.message
	case ErrorCode:
		return e.GetMessage()
	default:
		return e.Error()
	}
}
