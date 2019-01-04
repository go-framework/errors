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
	New(err interface{}, opts ...Option) error
	NewCode(code interface{}, detail string, opts ...Option) error
	NewError(code interface{}, message string, detail string, level Level, opts ...Option) error
}

// Message interface.
type Message interface {
	// Error is error interface.
	Error() string
	// Message is explain the error.
	Message() string
}
