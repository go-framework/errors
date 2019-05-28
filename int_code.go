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
