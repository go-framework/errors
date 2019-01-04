package errors

// Default Error interface type.
var DefaultError Error = &StringCodeError{}
// Debug mode pointer.
var Debug *bool = new(bool)

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
func New(err interface{}, opts ...Option) error {

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
	return DefaultError.New(err, opts...)
}

// New error with code interface, detail and options.
// code interface can be Error, Message and int type.
func NewCode(code interface{}, detail string, opts ...Option) error {

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

	// default use DefaultError.
	return DefaultError.NewCode(code, detail, opts...)
}

// New error with code interface, message, detail string, level and options.
func NewError(code interface{}, message string, detail string, level Level, opts ...Option) error {

	// in debug mode.
	if *Debug {
		opts = append(opts, WithLevel(Level_Debug))
		opts = append(opts, WithCaller(6, 64, true))
	} else {
		opts = append(opts, WithLevelCaller(7, 64))
	}

	// default use DefaultError.
	return DefaultError.NewError(code, message, detail, level, opts...)
}
