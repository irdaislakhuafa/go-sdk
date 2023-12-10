package errors

import (
	"fmt"
	"testing"

	"github.com/irdaislakhuafa/go-sdk/codes"
)

func Test_Stacktrace_Error(t *testing.T) {
	type fields struct {
		message  string
		cause    error
		code     codes.Code
		file     string
		function string
		line     int
	}

	type want struct {
		errMsg string
	}

	type test struct {
		name   string
		fields fields
		want   want
	}

	tests := []test{
		{
			name: "test success failed to start",
			fields: fields{
				message: "failed to start",
			},
			want: want{
				errMsg: "failed to start",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := stacktrace{
				message:  tt.fields.message,
				cause:    tt.fields.cause,
				code:     tt.fields.code,
				file:     tt.fields.file,
				function: tt.fields.function,
				line:     tt.fields.line,
			}

			if errMsg := st.Error(); errMsg != tt.want.errMsg {
				t.Fatalf("want result err msg is '%v' but got '%v'", tt.want.errMsg, errMsg)
			}
		})
	}
}

func Test_Stacktrace_ExitCode(t *testing.T) {
	type fields struct {
		message  string
		cause    error
		code     codes.Code
		file     string
		function string
		line     int
	}

	type want struct {
		code int
	}

	type test struct {
		name   string
		fields fields
		want   want
	}

	tests := []test{
		{
			name:   "test no code",
			fields: fields{code: codes.NoCode},
			want:   want{code: 1},
		},
		{
			name:   "test code auth failure",
			fields: fields{code: codes.CodeAuthFailure},
			want:   want{code: int(codes.CodeAuthFailure)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := stacktrace{
				message:  tt.fields.message,
				cause:    tt.fields.cause,
				code:     tt.fields.code,
				file:     tt.fields.file,
				function: tt.fields.function,
				line:     tt.fields.line,
			}

			if code := st.ExitCode(); code != int(tt.want.code) {
				t.Fatalf("want result exit code is '%v' but got '%v'", tt.want.code, code)
			}
		})
		fmt.Println("")
	}
}
