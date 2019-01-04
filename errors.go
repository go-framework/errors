package errors

import (
	"strings"
)

// multiple error.
type multipleError struct {
	list []error
}

// implement error interface, '\n' is separator error list.
func (e *multipleError) Error() string {
	buffer := strings.Builder{}

	for _, item := range e.list {
		buffer.WriteString(item.Error())
		buffer.WriteByte('\n')
	}

	return buffer.String()
}

// Append multiple error err to the end of error e.
func Append(e error, err ...error) error {
	// append nil error, then return nil.
	if e == nil && len(err) == 0 {
		return nil
	}
	// e is nil then new multipleError.
	if e == nil {
		e = &multipleError{
			list: err,
		}

		return e
	}
	// switch error type.
	switch m := e.(type) {
	case *multipleError:
		m.list = append(m.list, err...)
		return m
	default:
		e2 := &multipleError{}
		e2.list = append(e2.list, e)
		e2.list = append(e2.list, err...)
		return e2
	}
}
