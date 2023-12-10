package errors

import (
	"testing"

	"github.com/irdaislakhuafa/go-sdk/codes"
)

func Test_StacktraceError(t *testing.T) {
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
