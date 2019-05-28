package errors

import (
	"fmt"
)

// Debug mode pointer.
var Debug = new(bool)

// Enable debug mode.
func DebugEnable() {
	*Debug = true
}

// Disable debug mode.
func DebugDisable() {
	*Debug = false
}

// New error with interface and options.
// interface can be Error, Message, error, string type.
func New(err interface{}, opts ...Option) Error {

	if err == nil {
		return nil
	}

	// in debug mode.
	if *Debug {
		opts = append(opts, WithLevel(Level_Debug))
		opts = append(opts, WithCaller(6, 64, true))
	} else {
		opts = append(opts, WithLevelCaller(7, 64))
	}

	// implement Error interface.
	if e, ok := err.(Error); ok {
		return e.New(err, opts...)
	}

	// default use DefaultError.
	return new(StringCodeError).New(err, opts...)
}

// New Error interface with code interface, detail and options.
// code interface can be Error, Message and int type.
func NewCode(code interface{}, detail interface{}, opts ...Option) Error {

	if detail == nil {
		return nil
	}

	// in debug mode.
	if *Debug {
		opts = append(opts, WithLevel(Level_Debug))
		opts = append(opts, WithCaller(6, 64, true))
	} else {
		opts = append(opts, WithLevelCaller(7, 64))
	}

	// implement Error interface.
	if e, ok := code.(Error); ok {
		return e.NewCode(code, detail, opts...)
	}

	// select use Error handler type.
	return selectErrorHandler(code).NewCode(code, detail, opts...)
}

// New Error interface with format detail.
// code interface can be Error, Message and int type.
func NewCodeSprintf(code interface{}, format string, a ...interface{}) Error {

	var opts []Option

	// in debug mode.
	if *Debug {
		opts = append(opts, WithLevel(Level_Debug))
		opts = append(opts, WithCaller(6, 64, true))
	} else {
		opts = append(opts, WithLevelCaller(7, 64))
	}

	detail := fmt.Sprintf(format, a...)

	// implement Error interface.
	if e, ok := code.(Error); ok {
		return e.NewCode(code, detail, opts...)
	}

	// select use Error handler type.
	return selectErrorHandler(code).NewCode(code, detail, opts...)
}

// New Error interface with code interface, message, detail string, level and options.
func NewError(code interface{}, message string, detail string, level Level, opts ...Option) Error {

	// in debug mode.
	if *Debug {
		opts = append(opts, WithLevel(Level_Debug))
		opts = append(opts, WithCaller(6, 64, true))
	} else {
		opts = append(opts, WithLevelCaller(7, 64))
	}

	// select use Error handler type.
	return selectErrorHandler(code).NewError(code, message, detail, level, opts...)
}

// Equal two error, if equaled return true.
func Equal(e1 error, e2 error) bool {
	// implement Error interface.
	if e, ok := e1.(Error); ok {
		return e.Equal(e2)
	} else if e, ok := e2.(Error); ok {
		return e.Equal(e1)
	}

	// implement Message interface.
	if e, ok := e1.(Message); ok {
		if ee, ok := e2.(Message); ok {
			return e.GetMessage() == ee.GetMessage()
		}
		return false
	}

	return e1.Error() == e2.Error()
}

// Equal code with err which implement Error interface, if equaled return true.
func EqualCode(code interface{}, err error) bool {
	// implement Error interface.
	if e, ok := err.(Error); ok {
		return e.EqualCode(code)
	}

	return false
}
