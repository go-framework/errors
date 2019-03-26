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
// Debug level have error caller and print stack trace.
// Normal level have no error caller and stack trace.
// Critical level have error caller and stack trace.
// Panic level have error caller and print stack trace, then panic error.
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
		if e.GetCaller() == nil {
			e.SetCaller(NewCaller(skip, deep, stack))
		}
	})
}

// With level caller option.
// set option's caller.
func WithLevelCaller(skip, deep int) Option {
	return optionFunc(func(e Error) {
		if e.GetCaller() == nil {
			e.SetCaller(e.GetLevel().GetCaller(skip, deep))
		}
	})
}
