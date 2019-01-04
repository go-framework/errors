package errors

import (
	"fmt"
	"strconv"
	"strings"
)

const MaxUint64 = uint64(1<<64 - 1)

// convert ui to uint64. when i is other type then return max uint64 value.
func uiToUint64(ui interface{}) uint64 {
	switch v := ui.(type) {
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return uint64(v)
	}

	return MaxUint64
}

// Error code is a uint64 type, implement Error interface.
func (m *UintCodeError) GetErrCode() interface{} {
	if m != nil {
		return m.Code
	}
	return MaxUint64
}

// Set error code, implement Error interface.
func (m *UintCodeError) SetErrCode(i interface{}) {
	m.Code = uiToUint64(i)
}

// Set level, implement Error interface.
func (m *UintCodeError) SetLevel(l Level) {
	m.Level = l
}

// Set message, implement Error interface.
func (m *UintCodeError) SetMessage(message string) {
	m.Message = message
}

// Set detail, implement Error interface.
func (m *UintCodeError) SetDetail(detail string) {
	m.Detail = detail
}

// Set caller, implement Error interface.
func (m *UintCodeError) SetCaller(caller *Caller) {
	m.Caller = caller
}

// Implement error interface.
func (m *UintCodeError) Error() string {
	buffer := strings.Builder{}

	buffer.WriteString("level:")
	buffer.WriteString(m.Level.String())

	buffer.WriteString(" code:")
	buffer.WriteString(strconv.FormatUint(m.Code, 10))

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
func (m *UintCodeError) New(err interface{}, opts ...Option) error {

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
func (m *UintCodeError) NewCode(code interface{}, detail string, opts ...Option) error {

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
// Code can be uint type.
func (m *UintCodeError) NewError(code interface{}, message string, detail string, level Level, opts ...Option) error {

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
