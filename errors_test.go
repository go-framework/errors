package errors

import (
	"errors"
	"fmt"
	"testing"
)

type intCode int
type uintCode uint
type stringCode string
type structCode struct {
	Name string
}
type intErrCode int

const (
	Succeed intErrCode = iota
)

func (i intErrCode) String() string {
	switch i {
	case Succeed:
		return "succeed"
	default:
		return fmt.Sprintf("%d", i)
	}
}

func (i intErrCode) GetCode() interface{} {
	return i
}

func (i intErrCode) GetMessage() string {
	return i.String()
}
func (i intErrCode) Error() string {
	switch i {
	case Succeed:
		return "ok"
	default:
		return fmt.Sprintf("intErrCode(%d)", i)
	}
}

type stringErrCode string

const (
	Failed stringErrCode = "failed"
)

func (i stringErrCode) String() string {
	return string(i)
}

func (i stringErrCode) GetCode() interface{} {
	return i
}

func (i stringErrCode) GetMessage() string {
	return i.String()
}
func (i stringErrCode) Error() string {
	switch i {
	case Failed:
		return "failed error"
	default:
		return fmt.Sprintf("stringErr(%s)", i)
	}
}

func TestNew(t *testing.T) {

	type args struct {
		any interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "string",
			args: args{
				any: "string",
			},
			want: `{"message":"string"}`,
		},
		{
			name: "int",
			args: args{
				any: int(0),
			},
			want: `{"code":0}`,
		},
		{
			name: "int8",
			args: args{
				any: int8(0),
			},
			want: `{"code":0}`,
		},
		{
			name: "int16",
			args: args{
				any: int16(0),
			},
			want: `{"code":0}`,
		},
		{
			name: "int32",
			args: args{
				any: int32(0),
			},
			want: `{"code":0}`,
		},
		{
			name: "int64",
			args: args{
				any: int64(0),
			},
			want: `{"code":0}`,
		},
		{
			name: "uint",
			args: args{
				any: uint(0),
			},
			want: `{"code":0}`,
		},
		{
			name: "uint8",
			args: args{
				any: uint8(0),
			},
			want: `{"code":0}`,
		},
		{
			name: "uint16",
			args: args{
				any: uint16(0),
			},
			want: `{"code":0}`,
		},
		{
			name: "uint32",
			args: args{
				any: uint32(0),
			},
			want: `{"code":0}`,
		},
		{
			name: "uint64",
			args: args{
				any: uint64(0),
			},
			want: `{"code":0}`,
		},
		{
			name: "intErrCode",
			args: args{
				any: Succeed,
			},
			want: `{"code":0,"message":"succeed"}`,
		},
		{
			name: "stringErrCode",
			args: args{
				any: Failed,
			},
			want: `{"code":"failed","message":"failed"}`,
		},
		{
			name: "errorCode",
			args: args{
				any: &errorCode{
					code:    1,
					message: "2",
					error:   nil,
					next:    nil,
				},
			},
			want: `{"code":1,"message":"2"}`,
		},
		{
			name: "errorCode2",
			args: args{
				any: &errorCode{
					code:    1,
					message: "2",
					error:   errors.New("3"),
					next:    nil,
				},
			},
			want: `{"code":1,"message":"2","error":"3"}`,
		},
		{
			name: "error",
			args: args{
				any: errors.New("errors"),
			},
			want: `{"error":"errors"}`,
		},
		{
			name: "intCode",
			args: args{
				any: intCode(-1),
			},
			want: `{"code":-1}`,
		},
		{
			name: "uintCode",
			args: args{
				any: uintCode(1),
			},
			want: `{"code":1}`,
		},
		{
			name: "stringCode",
			args: args{
				any: stringCode("stringCode"),
			},
			want: `{"message":"stringCode"}`,
		},
		{
			name: "structCode",
			args: args{
				any: structCode{
					Name: "test",
				},
			},
			want: `{"message":"{Name:test}"}`,
		},
		{
			name: "structCode2",
			args: args{
				any: &structCode{
					Name: "test",
				},
			},
			want: `{"message":"{Name:test}"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.any); got.Error() != tt.want {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}
