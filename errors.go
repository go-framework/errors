package errors

import (
	"strconv"
	"strings"
)

// Error list.
type Errors []error

// New Errors.
func NewErrors() Errors {
	return make([]error, 0)
}

// Implement error interface, '\n' is separator error list.
func (errs Errors) Error() string {
	if len(errs) == 0 {
		return "<nil>"
	}

	buffer := strings.Builder{}

	buffer.WriteString("error list:")
	for idx, item := range errs {
		if item == nil {
			continue
		}
		buffer.WriteString("\n\t* ")
		buffer.WriteString(strconv.Itoa(idx + 1))
		buffer.WriteString(" ")
		// replace \n \t
		str := strings.Replace(item.Error(), "\t", "\t\t", -1)
		str = strings.Replace(str, "\n", "\n\t\t", -1)
		buffer.WriteString(str)
	}

	return buffer.String()
}

// Append multiple error.
// ignore nil error.
func (errs *Errors) Append(err ...error) {
	if len(err) == 0 {
		return
	}

	for _, e := range err {
		if e == nil {
			continue
		}
		// is self type?
		if t, ok := e.(Errors); ok {
			*errs = append(*errs, t...)
			continue
		}

		*errs = append(*errs, e)
	}
}

// Errors is nil error.
func (errs Errors) Nil() error {
	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Append multiple error err to the end of error errs.
func Append(e error, err ...error) error {
	// append a empty error list, return e.
	if len(err) == 0 {
		return e
	}

	// switch error type.
	switch t := e.(type) {
	case Errors:
		t.Append(err...)
		return t
	default:
		errs := NewErrors()

		errs.Append(e)
		errs.Append(err...)

		return errs
	}
}
