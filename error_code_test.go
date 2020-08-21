package errors

import (
	"errors"
	"math"
	"reflect"
	"testing"
)

func Test_errorCode_marshalJSON(t *testing.T) {
	type fields struct {
		code    interface{}
		message string
		error   error
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "MaxInt64",
			fields: fields{
				code:    math.MaxInt64,
				message: "failed",
				error:   errors.New("unexpected"),
			},
			want:    []byte(`{"code":9223372036854775807,"message":"failed","error":"unexpected"}`),
			wantErr: false,
		},
		{
			name: "MinInt64",
			fields: fields{
				code:    math.MinInt64,
				message: "failed",
				error:   errors.New("unexpected"),
			},
			want:    []byte(`{"code":-9223372036854775808,"message":"failed","error":"unexpected"}`),
			wantErr: false,
		},
		{
			name: "0",
			fields: fields{
				code:    0,
				message: "failed",
				error:   errors.New("unexpected"),
			},
			want:    []byte(`{"code":0,"message":"failed","error":"unexpected"}`),
			wantErr: false,
		},
		{
			name: "MaxUint64",
			fields: fields{
				code:    uint64(math.MaxUint64),
				message: "failed",
				error:   errors.New("unexpected"),
			},
			want:    []byte(`{"code":18446744073709551615,"message":"failed","error":"unexpected"}`),
			wantErr: false,
		},
		{
			name: "String",
			fields: fields{
				code:    "undefined",
				message: "failed",
				error:   errors.New("unexpected"),
			},
			want:    []byte(`{"code":"undefined","message":"failed","error":"unexpected"}`),
			wantErr: false,
		},
		{
			name: "Nil",
			fields: fields{
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Just error",
			fields: fields{
				error: errors.New("unexpected"),
			},
			want:    []byte(`{"error":"unexpected"}`),
			wantErr: false,
		},
		{
			name: "code and Error",
			fields: fields{
				code:  1,
				error: errors.New("unexpected"),
			},
			want:    []byte(`{"code":1,"error":"unexpected"}`),
			wantErr: false,
		},
		{
			name: "message and Error",
			fields: fields{
				message: "test",
				error:   errors.New("unexpected"),
			},
			want:    []byte(`{"message":"test","error":"unexpected"}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := errorCode{
				code:    tt.fields.code,
				message: tt.fields.message,
				error:   tt.fields.error,
			}
			got, err := e.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("marshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("marshalJSON() got = %s, want %s", got, tt.want)
			}
		})
	}
}

func Test_errorCode_MarshalJSON(t *testing.T) {
	type fields struct {
		code    interface{}
		message string
		error   error
		next    *errorCode
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "MaxUint64",
			fields: fields{
				code:    uint64(math.MaxUint64),
				message: "failed",
				error:   errors.New("unexpected"),
			},
			want:    []byte(`{"code":18446744073709551615,"message":"failed","error":"unexpected"}`),
			wantErr: false,
		},
		{
			name: "String",
			fields: fields{
				code:    "undefined",
				message: "failed",
				error:   errors.New("unexpected"),
			},
			want:    []byte(`{"code":"undefined","message":"failed","error":"unexpected"}`),
			wantErr: false,
		},
		{
			name: "Wrap",
			fields: fields{
				code:    "1",
				message: "failed",
				error:   errors.New("unexpected"),
				next: &errorCode{
					code:    2,
					message: "failed",
					error:   errors.New("unexpected"),
					next: &errorCode{
						code:    math.MaxInt64,
						message: "failed",
						error:   errors.New("unexpected"),
					},
				},
			},
			want:    []byte(`[{"code":"1","message":"failed","error":"unexpected"},{"code":2,"message":"failed","error":"unexpected"},{"code":9223372036854775807,"message":"failed","error":"unexpected"}]`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorCode{
				code:    tt.fields.code,
				message: tt.fields.message,
				error:   tt.fields.error,
				next:    tt.fields.next,
			}
			got, err := e.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %s, want %s", got, tt.want)
			}
		})
	}
}

func Test_errorCode_Wrap(t *testing.T) {
	type fields struct {
		code    interface{}
		message string
		error   error
		next    *errorCode
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
	}{
		{
			name: "nil",
			fields: fields{
				code:    math.MaxInt64,
				message: "failed",
				error:   errors.New("unexpected"),
			},
			args: args{
				err: nil,
			},
			want: &errorCode{
				code:    math.MaxInt64,
				message: "failed",
				error:   errors.New("unexpected"),
			},
		},
		{
			name: "errors",
			fields: fields{
				code:    math.MaxInt64,
				message: "failed",
				error:   errors.New("unexpected"),
			},
			args: args{
				err: errors.New("errors"),
			},
			want: nil,
		},
		//{
		//	name: "Errorf",
		//	fields: fields{
		//		code:    math.MaxInt64,
		//		message: "failed",
		//		error:   errors.New("unexpected"),
		//	},
		//	args: args{
		//		err: fmt.Errorf("%s", "Errorf"),
		//	},
		//	want: nil,
		//},
		//{
		//	name: "Errorf Wrap",
		//	fields: fields{
		//		code:    math.MaxInt64,
		//		message: "failed",
		//		error:   errors.New("unexpected"),
		//	},
		//	args: args{
		//		err: fmt.Errorf("wrap %w", errors.New("errors")),
		//	},
		//	want: nil,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorCode{
				code:    tt.fields.code,
				message: tt.fields.message,
				error:   tt.fields.error,
				next:    tt.fields.next,
			}
			got := e.Wrap(tt.args.err)
			if !errors.Is(tt.want, got) {
				t.Errorf("Wrap() failed got = %v, want %v", got, tt.want)
			}
		})
	}
}
