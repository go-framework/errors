package errors

type Error interface {
	// Get function.
	GetErrCode() interface{}
	GetLevel() Level
	GetMessage() string
	GetDetail() string
	GetCaller() *Caller
	// Implement error interface.
	Error() string
	// Set function.
	SetErrCode(interface{})
	SetLevel(Level)
	SetMessage(string)
	SetDetail(string)
	SetCaller(*Caller)
	// New function.
	New(err interface{}, opts ...Option) Error
	NewCode(code interface{}, detail interface{}, opts ...Option) Error
	NewError(code interface{}, message string, detail string, level Level, opts ...Option) Error
	// Equal function.
	Equal(err error) bool
	EqualCode(code interface{}) bool
}

// Message interface.
type Message interface {
	// Error is error interface.
	Error() string
	// Message is explain the error.
	Message() string
}
