package errors

import (
	"fmt"
	"strconv"
	"strings"
)

const MaxInt64 = int64(^uint64(0) >> 1)

// convert i to int64. when i is other type then return max int64 value.
func iToInt64(i interface{}) int64 {

	switch v := i.(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return int64(v)
	}

	return MaxInt64
}

// Error code is a int64 type, implement Error interface.
func (m *IntCodeError) GetErrCode() interface{} {
	if m != nil {
		return m.Code
	}
	return MaxInt64
}

// Set error code, implement Error interface.
func (m *IntCodeError) SetErrCode(i interface{}) {
	m.Code = iToInt64(i)
}

// Set level, implement Error interface.
func (m *IntCodeError) SetLevel(l Level) {
	m.Level = l
}

// Set message, implement Error interface.
func (m *IntCodeError) SetMessage(message string) {
	m.Message = message
}

// Set detail, implement Error interface.
func (m *IntCodeError) SetDetail(detail string) {
	m.Detail = detail
}

// Set caller, implement Error interface.
func (m *IntCodeError) SetCaller(caller *Caller) {
	m.Caller = caller
}

// Implement error interface.
func (m *IntCodeError) Error() string {
	buffer := strings.Builder{}

	buffer.WriteString("level:")
	buffer.WriteString(m.Level.String())

	buffer.WriteString(" code:")
	buffer.WriteString(strconv.FormatInt(m.Code, 10))

	if len(m.Message) > 0 {
		buffer.WriteString(" message:")
		buffer.WriteString(m.Message)
	}

	if len(m.Detail) > 0 {
		buffer.WriteString(" detail:")
		buffer.WriteString(m.Detail)
	}

	if m.Caller != nil {
		buffer.WriteString(" file:")
		buffer.WriteString(m.Caller.File)
		buffer.WriteString(" function:")
		buffer.WriteString(m.Caller.Function)
		buffer.WriteString(" line:")
		buffer.WriteString(strconv.FormatInt(int64(m.Caller.Line), 10))

		if len(m.Caller.Stacks) > 0 {
			buffer.WriteByte('\n')
			buffer.WriteString(strings.Join(m.Caller.Stacks, "\n"))
		}
	}

	return buffer.String()
}

// New error with interface and options.
// interface can be Error, Message, error, string type.
func (m *IntCodeError) New(err interface{}, opts ...Option) Error {

	if e, ok := err.(Error); ok {
		// if detail is implement Error interface.
		m.SetErrCode(e.GetErrCode())
		m.Level = e.GetLevel()
		m.Message = e.GetMessage()
		m.Detail = e.GetDetail()
	} else if e, ok := err.(Message); ok {
		// if detail is implement Message interface.
		m.Message = e.Message()
		m.Detail = e.Error()
	} else if e, ok := err.(error); ok {
		// if detail is implement error interface.
		m.Detail = e.Error()
	} else if e, ok := err.(string); ok {
		// if detail is string type.
		m.Detail = e
	} else {
		// otherwise.
		m.Detail = fmt.Sprintf("%v", err)
	}

	// apply options.
	for _, opt := range opts {
		opt.apply(m)
	}

	// get caller.
	if m.Caller == nil {
		m.Caller = m.Level.GetCaller(4, 64)
	}

	// trace with level.
	m.Level.Trace(m)

	return m
}

// New error with code interface, detail and options.
// code interface can be Error, Message and int type.
func (m *IntCodeError) NewCode(code interface{}, detail string, opts ...Option) Error {

	if c, ok := code.(Error); ok {
		// if detail is implement Error interface.
		m.SetErrCode(c.GetErrCode())
		m.Level = c.GetLevel()
		m.Message = c.GetMessage()
	} else if c, ok := code.(Message); ok {
		// if detail is implement Message interface.
		m.Message = c.Message()
	} else {
		// code is int type.
		m.SetErrCode(code)
	}

	// set detail value.
	m.Detail = detail

	// apply options.
	for _, opt := range opts {
		opt.apply(m)
	}

	// get caller.
	if m.Caller == nil {
		m.Caller = m.Level.GetCaller(4, 64)
	}

	// trace with level.
	m.Level.Trace(m)

	return m
}

// New error with code interface, message, detail string, level and options.
// Code can be int type.
func (m *IntCodeError) NewError(code interface{}, message string, detail string, level Level, opts ...Option) Error {

	m.SetErrCode(code)
	m.Message = message
	m.Detail = detail
	m.Level = level

	// apply options.
	for _, opt := range opts {
		opt.apply(m)
	}

	// get caller.
	if m.Caller == nil {
		m.Caller = m.Level.GetCaller(4, 64)
	}

	// trace with level.
	m.Level.Trace(m)

	return m
}
