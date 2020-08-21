package errors

import (
	"math"
	"reflect"
)

// convert i to int64. when i is other type then return max int64 value.
func iToInt64(i interface{}) int64 {
	switch v := i.(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return int64(v)
	}

	switch reflect.TypeOf(i).Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return reflect.ValueOf(i).Int()
	}

	return math.MaxInt64
}

// convert ui to uint64. when i is other type then return max uint64 value.
func iToUint64(ui interface{}) uint64 {
	switch v := ui.(type) {
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return uint64(v)
	}

	switch reflect.TypeOf(ui).Kind() {
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		return reflect.ValueOf(ui).Uint()
	}

	return math.MaxUint64
}

// convert i to string. when i is other type then return empty string.
func iToString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	}

	switch reflect.TypeOf(i).Kind() {
	case reflect.String:
		return reflect.ValueOf(i).String()
	}

	return ""
}
