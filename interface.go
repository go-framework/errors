package errors

type Error interface {
	GetErrCode() interface{}
	GetLevel() Level
	GetMessage() string
	GetDetail() string
	GetCaller() *Caller
	Error() string
	SetErrCode(interface{})
	SetLevel(Level)
	SetMessage(string)
	SetDetail(string)
	SetCaller(*Caller)
}

// type Error interface {
// 	New(detail interface{}) error
// 	NewCode(code interface{}, detail interface{}) error
// 	NewError(code interface{}, message string, detail interface{}, level Level) error
// }

// Message interface.
type Message interface {
	// Error is error interface.
	Error() string
	// Message is explain the error.
	Message() string
}
