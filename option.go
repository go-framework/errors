package errors

// An Option configures Error interface.
type Option interface {
	apply(Error)
}

// optionFunc wraps a func so it satisfies the Option interface.
type optionFunc func(Error)

func (f optionFunc) apply(e Error) {
	f(e)
}

// With level option.
// set option's level.
// DebugLevel have error source and print stack trace.
// NormalLevel have no error source and stack trace.
// WarnLevel have error source no stack trace.
// PanicLevel have error source and print stack trace, then panic error.
// FatalLevel have error source and no stack trace, then calls panic and os.Exit(1).
func WithLevel(level Level) Option {
	return optionFunc(func(e Error) {
		e.SetLevel(level)
	})
}

// With message option.
// set option's message.
func WithMessage(message string) Option {
	return optionFunc(func(e Error) {
		e.SetMessage(message)
	})
}

// With caller option.
// set option's caller.
func WithCaller(skip, deep int, stack bool) Option {
	return optionFunc(func(e Error) {
		e.SetCaller(NewCaller(skip, deep, stack))
	})
}

// With caller option.
// set option's caller.
func WithLevelCaller(skip, deep int) Option {
	return optionFunc(func(e Error) {
		e.SetCaller(e.GetLevel().GetCaller(skip, deep))
	})
}
