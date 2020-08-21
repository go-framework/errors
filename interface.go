package errors

// ErrorCode interface
type ErrorCode interface {
	GetCode() interface{}
	GetMessage() string
}
