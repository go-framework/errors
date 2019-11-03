package errors

import (
	"reflect"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestTextError_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		want    []byte
		wantErr bool
	}{
		{
			name:    "empty string",
			err:     NewTextError(""),
			want:    []byte("null"),
			wantErr: false,
		},
		{
			name:    "normal",
			err:     NewTextError("test"),
			want:    []byte(`{"text":"test"}`),
			wantErr: false,
		},
		{
			name:    "newline",
			err:     NewTextError("test newline \n error"),
			want:    []byte(`{"text":"test newline \n error"}`),
			wantErr: false,
		},
		{
			name:    "tab",
			err:     NewTextError("test tab \t error"),
			want:    []byte(`{"text":"test tab \t error"}`),
			wantErr: false,
		},
		{
			name:    "newline tab",
			err:     NewTextError("test tab \n\t error"),
			want:    []byte(`{"text":"test tab \n\t error"}`),
			wantErr: false,
		},
		{
			name:    "tab newline",
			err:     NewTextError("test tab \t\n error"),
			want:    []byte(`{"text":"test tab \t\n error"}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := jsoniter.Marshal(tt.err)
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
