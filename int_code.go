package errors

// Defined GetIntCodeTextHandler type.
type GetIntCodeTextHandler func(IntCode) string

// Global variable.
var getIntCodeTextHandlers []GetIntCodeTextHandler

// Register GetIntCodeTextHandler.
func RegisterGetIntCodeTextHandler(f GetIntCodeTextHandler) {
	getIntCodeTextHandlers = append(getIntCodeTextHandlers, f)
}

// get code text.
func getIntCodeText(e IntCode) string {
	for _, f := range getIntCodeTextHandlers {
		str := f(e)
		if len(str) > 0 {
			return str
		}
	}
	return ""
}

// Int code offset.
const IntCodeOffset = -100
const UserIntCodeOffset = -10000

// Defined IntCode type.
type IntCode int64

// Get code raw value.
func (e IntCode) Value() int64 {
	return int64(e)
}

// Get code raw value.
func (e IntCode) Int() int {
	return int(e)
}

// Implement error interface.
func (e IntCode) Error() string {
	return getIntCodeText(e)
}

// Implement Message interface.
func (e IntCode) GetCode() interface{} {
	return e
}

// Implement Message interface.
func (e IntCode) GetMessage() string {
	return getIntCodeText(e)
}

// Return a new error with detail.
func (e IntCode) WithDetail(detail interface{}, opts ...Option) error {
	if detail == nil {
		return nil
	}
	opts = append(opts, WithLevelCaller(8, 64))
	return NewCode(e, detail, opts...)
}

// Return a new error with normal level detail.
func (e IntCode) Normal(detail interface{}, opts ...Option) error {
	if detail == nil {
		return nil
	}
	opts = append(opts, WithSetLevelCaller(Level_Normal, 8, 64))
	return NewCode(e, detail, opts...)
}

// Return a new error with debug level detail.
func (e IntCode) Debug(detail interface{}, opts ...Option) error {
	if detail == nil {
		return nil
	}
	opts = append(opts, WithSetLevelCaller(Level_Debug, 8, 64))
	return NewCode(e, detail, opts...)
}

// Return a new error with critical level detail.
func (e IntCode) Warn(detail interface{}, opts ...Option) error {
	if detail == nil {
		return nil
	}
	opts = append(opts, WithSetLevelCaller(Level_Warn, 8, 64))
	return NewCode(e, detail, opts...)
}

// Return a new error with critical level detail.
func (e IntCode) Critical(detail interface{}, opts ...Option) error {
	if detail == nil {
		return nil
	}
	opts = append(opts, WithSetLevelCaller(Level_Critical, 8, 64))
	return NewCode(e, detail, opts...)
}

// Return a new error with panic level detail.
func (e IntCode) Panic(detail interface{}, opts ...Option) error {
	if detail == nil {
		return nil
	}
	opts = append(opts, WithSetLevelCaller(Level_Panic, 8, 64))
	return NewCode(e, detail, opts...)
}

// Return a new error with fatal level detail.
func (e IntCode) Fatal(detail interface{}, opts ...Option) error {
	if detail == nil {
		return nil
	}
	opts = append(opts, WithSetLevelCaller(Level_Fatal, 8, 64))
	return NewCode(e, detail, opts...)
}
