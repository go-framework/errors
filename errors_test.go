package errors

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestAppend(t *testing.T) {
	var errs error
	t.Log(errs)

	errs = Append(errs, fmt.Errorf("error 1"), fmt.Errorf("error 2"))
	t.Log(errs)

	errs = Append(errs, fmt.Errorf("error 3"), fmt.Errorf("error 4"))
	t.Log(errs)

	errs = fmt.Errorf("error 0")

	t.Log(errs)
	errs = Append(errs, fmt.Errorf("error 1"), fmt.Errorf("error 2"))
	t.Log(errs)

	errs = Append(errs, fmt.Errorf("error 3"), fmt.Errorf("error 4"))
	t.Log(errs)

	var errs2 error
	t.Log(errs2)

	errs2 = Append(errs2, errs)

	t.Log(errs2)

	errs2 = Append(errs2, errs)

	t.Log(errs2)
}

func TestErrors(t *testing.T) {
	var errs Errors

	t.Log("Errors", errs)

	//
	// Errors
	//

	errs.Append(nil)
	errs.Append(errors.New("go errors error"))
	errs.Append(NewTextError("text error"))
	errs.Append(NewTextError("text error \n separator"))
	errs.Append(NewTextError("text error \t separator"))
	errs.Append(NewTextError("text error \n\t separator"))
	errs.Append(NewTextError("text error \t\n separator"))

	var intCode IntCode = -1
	errs.Append(intCode.WithDetail("IntCode"))

	t.Log("Errors", errs)

	//
	// error
	//
	var err error = errs

	t.Log("error", err)
}

func TestJSON(t *testing.T) {
	var errs Errors

	//
	// Errors
	//

	errs.Append(nil)
	errs.Append(errors.New("go errors error"))
	errs.Append(NewTextError("test error"))
	errs.Append(NewTextError("test error \n separator"))
	errs.Append(NewTextError("test error \t separator"))
	errs.Append(NewTextError("test error \n\t separator"))
	errs.Append(NewTextError("test error \t\n separator"))

	var intCode IntCode = -1
	errs.Append(intCode.WithDetail("IntCode"))

	t.Log("Errors", errs)

	str, err := jsoniter.MarshalToString(errs)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(str)
}

func TestErrors_MarshalJSON(t *testing.T) {

	tests := []struct {
		name    string
		errs    Errors
		want    []byte
		wantErr bool
	}{
		{
			name:    "empty string",
			errs:    []error{nil},
			want:    []byte("[null]"),
			wantErr: false,
		},
		{
			name: "normal",
			errs: []error{
				errors.New("go error"),
				NewTextError("text error"),
			},
			want:    []byte(`[{"error":"go error"},{"error":"text error"}]`),
			wantErr: false,
		},
		{
			name: "newline tab",
			errs: []error{
				errors.New("test newline \n\t error"),
				NewTextError("test newline \n\t error"),
			},
			want:    []byte(`[{"error":"test newline \n\t error"},{"error":"test newline \n\t error"}]`),
			wantErr: false,
		},
		{
			name: "int code error",
			errs: []error{
				IntCode(-1).WithDetail("-1"),
				IntCode(-2).WithDetail("newline \n error"),
				IntCode(-3).WithDetail("tab \t error"),
				IntCode(-4).WithDetail("tab newline \t\n error"),
			},
			want:    []byte(`[{"code":-1,"detail":"-1"},{"code":-2,"detail":"newline \n error"},{"code":-3,"detail":"tab \t error"},{"code":-4,"detail":"tab newline \t\n error"}]`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := jsoniter.Marshal(tt.errs)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", string(got), string(tt.want))
			} else {
				t.Logf("MarshalJSON() got = %v", string(got))
			}
		})
	}
}
