package intcode

import (
	"github.com/go-framework/errors"
)

// init
func init() {
	errors.RegisterGetIntCodeTextHandler(GetCodeText)
}

// Get error code.
func GetCode(err error) errors.IntCode {
	if err == nil {
		return Succeed
	}
	if i, ok := err.(errors.Message); ok {
		return errors.IntCode(errors.ToInt64(i.GetCode()))
	} else if i, ok := err.(errors.Error); ok {
		return errors.IntCode(errors.ToInt64(i.GetErrCode()))
	}
	return ErrUndefined
}

// Get error message.
func GetMessage(err error) string {
	if err == nil {
		return GetCodeText(Succeed)
	}
	if i, ok := err.(errors.Message); ok {
		return i.GetMessage()
	} else if i, ok := err.(errors.Error); ok {
		return i.GetMessage()
	}
	return GetCodeText(ErrUndefined)
}

// Get detail from any interface.
func GetDetail(err interface{}) string {
	return errors.GetDetail(err)
}

// Get int error code text.
func GetCodeText(e errors.IntCode) string {
	return CodeTexts[e]
}

// Const int error code.
const (
	Succeed errors.IntCode = iota * -1
	Errors
	ErrUndefined
	ErrUnsupported errors.IntCode = (iota * -1) + errors.IntCodeOffset
	ErrNil
	ErrNull
	ErrMarshal
	ErrUnmarshal
	ErrCode
	ErrEncode
	ErrDecode
	ErrHttp
	ErrRedis
	ErrDatabase
	ErrTimeout
	ErrAuthorize
	ErrExist
	ErrNotExist
	ErrPermission
	ErrParameter
	ErrFormat
	ErrNotAllowed
	ErrValidation
	ErrSave
	ErrNotMatched
	ErrUnexpected
	ErrNotImplement
	ErrRequest
	ErrUnavailability
)

// Int error code text map.
var CodeTexts = map[errors.IntCode]string{
	Succeed:           "Succeed",
	Errors:            "Error list",
	ErrUndefined:      "Undefined code",
	ErrUnsupported:    "Unsupported error",
	ErrNil:            "Nil error",
	ErrNull:           "Null error",
	ErrMarshal:        "Marshal error",
	ErrUnmarshal:      "Unmarshal error",
	ErrCode:           "Code error",
	ErrEncode:         "Encode error",
	ErrDecode:         "Decode error",
	ErrHttp:           "Http error",
	ErrRedis:          "Redis error",
	ErrDatabase:       "Database error",
	ErrTimeout:        "Timeout error",
	ErrAuthorize:      "Authorize error",
	ErrExist:          "Is exist",
	ErrNotExist:       "Not exist",
	ErrPermission:     "Permission error",
	ErrParameter:      "Parameter error",
	ErrFormat:         "Format error",
	ErrNotAllowed:     "Not allowed error",
	ErrValidation:     "Validation error",
	ErrSave:           "Save error",
	ErrNotMatched:     "Not matched error",
	ErrUnexpected:     "Unexpected error",
	ErrNotImplement:   "Not implement error",
	ErrRequest:        "Request error",
	ErrUnavailability: "Unavailability error",
}
