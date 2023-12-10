package errors

import (
	"testing"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/files"
)

func Test_GetCaller(t *testing.T) {
	type args struct {
		err error
	}

	type want struct {
		file    string
		line    int
		message string
	}

	type test struct {
		name      string
		args      args
		want      want
		isWantErr bool
	}

	tests := []test{
		{
			name: "test error bad requets success",
			args: args{err: NewWithCode(codes.CodeBadRequest, "bad request")},
			want: want{
				file:    files.GetCurrentFileLocation(), // NOTE: this is line where error created with function NewWithCode
				line:    31,
				message: "bad request",
			},
			isWantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, line, message, err := GetCaller(tt.args.err)
			if isErr := (err != nil); isErr != tt.isWantErr {
				t.Fatalf("want err is '%v', got error '%v'", tt.isWantErr, err)
			}

			if message != tt.want.message {
				t.Fatalf("want result message is '%v' but got '%v'", tt.want.message, message)
			}

			if line != tt.want.line {
				t.Fatalf("want result line is '%v' but got '%v'", tt.want.line, line)
			}

			if file != tt.want.file {
				t.Fatalf("want result file is '%v' but got '%v'", tt.want.file, file)
			}
		})
	}
}
