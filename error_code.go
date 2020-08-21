package errors

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

// errorCode is a implementation of error.
type errorCode struct {
	code    interface{}
	message string
	error   error
	next    *errorCode
}

func NewErrorCode(code interface{}, message string, err error) Error {
	e := errorCode{
		code:    code,
		message: message,
		error:   err,
		next:    nil,
	}
	return &e
}

func NewCode(code interface{}) Error {
	e := errorCode{
		code: code,
	}
	return &e
}

func NewMessage(message string) Error {
	e := errorCode{
		message: message,
	}
	return &e
}

func NewError(err error) Error {
	e := errorCode{
		error: err,
	}
	return &e
}

func (e *errorCode) Error() string {
	data, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("error code marshal error %v", err)
	}
	return *(*string)(unsafe.Pointer(&data))
}

func (e *errorCode) Is(err error) bool {
	_e, ok := err.(*errorCode)
	if !ok {
		return false
	}
	if e.code != nil && _e.code != nil {
		return reflect.DeepEqual(e.code, _e.code)
	} else if e.code != nil && _e.code == nil {
		return false
	} else if e.code == nil && _e.code != nil {
		return false
	}
	return errors.Is(e.error, _e.error)
}

func (e *errorCode) As(target interface{}) bool {
	_, ok := target.(*errorCode)
	if !ok {
		return false
	}
	target = e
	return true
}

func (e *errorCode) Unwrap() error {
	return e.next
}

func (e *errorCode) Wrap(err error) error {
	if err == nil {
		return e
	}
	v, ok := err.(*errorCode)
	if ok {
		v.next = e
		return v
	}
	_e := &errorCode{
		error: err,
		next:  e,
	}
	return _e
}

func (e *errorCode) SetError(err error) error {
	e.error = err
	return e
}

func (e *errorCode) GetCode() interface{} {
	return e.code
}

func (e *errorCode) GetMessage() string {
	return e.message
}

func (e *errorCode) GetError() error {
	return e.error
}

func (e *errorCode) marshalJSON() ([]byte, error) {
	if e == nil {
		return nil, nil
	}

	n := 0
	if e.code != nil {
		n += 1
	}
	if len(e.message) > 0 {
		n += 2
	}
	if e.error != nil {
		n += 4
	}
	if n == 0 {
		return nil, nil
	}

	var buf = bytes.Buffer{}
	buf.WriteByte('{')

	if e.code != nil {
		buf.WriteString(`"code":`)
		switch v := e.code.(type) {
		case int:
			buf.Write(strconv.AppendInt(nil, int64(v), 10))
		case int8:
			buf.Write(strconv.AppendInt(nil, int64(v), 10))
		case int16:
			buf.Write(strconv.AppendInt(nil, int64(v), 10))
		case int32:
			buf.Write(strconv.AppendInt(nil, int64(v), 10))
		case int64:
			buf.Write(strconv.AppendInt(nil, int64(v), 10))
		case uint:
			buf.Write(strconv.AppendUint(nil, uint64(v), 10))
		case uint8:
			buf.Write(strconv.AppendUint(nil, uint64(v), 10))
		case uint16:
			buf.Write(strconv.AppendUint(nil, uint64(v), 10))
		case uint32:
			buf.Write(strconv.AppendUint(nil, uint64(v), 10))
		case uint64:
			buf.Write(strconv.AppendUint(nil, uint64(v), 10))
		case string:
			buf.WriteString(`"`)
			buf.WriteString(v)
			buf.WriteString(`"`)
		default:
			switch reflect.TypeOf(e.code).Kind() {
			case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
				buf.Write(strconv.AppendInt(nil, reflect.ValueOf(e.code).Int(), 10))
			case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
				buf.Write(strconv.AppendUint(nil, reflect.ValueOf(e.code).Uint(), 10))
			case reflect.String:
				buf.WriteString(`"`)
				buf.WriteString(reflect.ValueOf(e.code).String())
				buf.WriteString(`"`)
			default:
				fmt.Fprintf(&buf, `"%+v"`, e.code)
			}
		}
		if n > 1 {
			buf.WriteByte(',')
		}
	}

	if len(e.message) > 0 {
		buf.WriteString(`"message":"`)
		buf.WriteString(e.message)
		buf.WriteByte('"')
		if n&4 == 4 {
			buf.WriteByte(',')
		}
	}

	if e.error != nil {
		buf.WriteString(`"error":"`)
		buf.WriteString(e.error.Error())
		buf.WriteString(`"`)
	}

	buf.WriteString(`}`)

	return buf.Bytes(), nil
}

func (e *errorCode) MarshalJSON() ([]byte, error) {
	if e == nil {
		return nil, nil
	}
	if e.next == nil {
		return e.marshalJSON()
	}
	var buf = bytes.Buffer{}
	buf.WriteByte('[')
	for _e := e; _e != nil; {
		data, err := _e.marshalJSON()
		if err != nil {
			return nil, err
		}
		buf.Write(data)
		if _e.next != nil {
			buf.WriteByte(',')
		}
		_e = _e.next
	}
	buf.WriteByte(']')
	return buf.Bytes(), nil
}
