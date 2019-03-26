package errors

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"time"
)

// Panic function.
var PanicFunc = func(v interface{}) {
	panic(v)
}

// CapitalString returns an all-caps ASCII representation of the log level.
func (x Level) CapitalString() string {
	// Printing levels in all-caps is common enough that we should export this
	// functionality.
	switch x {
	case Level_Normal:
		return "NORMAL"
	case Level_Debug:
		return "DEBUG"
	case Level_Critical:
		return "Critical"
	case Level_Panic:
		return "PANIC"
	case Level_Fatal:
		return "FATAL"
	default:
		return fmt.Sprintf("LEVEL(%d)", x)
	}
}

// MarshalText marshals the Level to text. Note that the text representation
// drops the -Level suffix (see example).
func (x Level) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText unmarshals text to a level. Like MarshalText, UnmarshalText
// expects the text representation of a Level to drop the -Level suffix (see
// example).
//
// In particular, this makes it easy to configure logging levels using YAML,
// TOML, or JSON files.
func (x *Level) UnmarshalText(text []byte) error {
	if x == nil {
		return errors.New("can't unmarshal a nil *Level")
	}
	if !x.unmarshalText(text) && !x.unmarshalText(bytes.ToLower(text)) {
		return fmt.Errorf("unrecognized level: %q", text)
	}
	return nil
}

func (x *Level) unmarshalText(text []byte) bool {
	switch string(text) {
	case "normal", "NORMAL", "": // make the zero value useful
		*x = Level_Normal
	case "debug", "DEBUG":
		*x = Level_Debug
	case "critical", "CRITICAL":
		*x = Level_Critical
	case "panic", "PANIC":
		*x = Level_Panic
	case "fatal", "FATAL":
		*x = Level_Fatal
	default:
		return false
	}
	return true
}

// Set sets the level for the flag.Value interface.
func (x *Level) Set(s string) error {
	return x.UnmarshalText([]byte(s))
}

// Get gets the level for the flag.Getter interface.
func (x *Level) Get() interface{} {
	return *x
}

// Get the caller stack trace by error level.
// Normal level have no error caller and stack trace.
// Debug level have error caller and print stack trace.
// Critical level have error caller and stack trace.
// Panic level have error caller and print stack trace, then panic error.
// Fatal level have error caller and print stack trace, then call to os.Exit(1).
func (x Level) GetCaller(skip, deep int) *Caller {

	switch x {
	case Level_Debug, Level_Critical, Level_Panic, Level_Fatal:
		return NewCaller(skip, deep, true)
	case Level_Normal:
		return nil
	default:
		return NewCaller(skip, deep, false)
	}
}

var timer *time.Timer
// Trace level.
// Panic level have error caller and print stack trace, then panic error.
// Fatal level have error caller and print stack trace, then call to os.Exit(1).
func (x Level) Trace(err error) {
	switch x {
	case Level_Panic:
		PanicFunc(err)
	case Level_Fatal:
		log.Fatalln(err)
	}
}
