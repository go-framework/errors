package errors

// Code interface.
type Code interface {
	Int() int
	Int32() int32
	Int64() int64
	Uint() uint
	Uint32() uint32
	Uint64() uint64
	String() string
	Interface() interface{}
}
