package errors

import (
	"errors"
	"testing"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/files"
)

func Test_GetCaller(t *testing.T) {
	type args struct {
		err error
	}

	type wantResult struct {
		file    string
		line    int
		message string
	}

	type wantErr struct {
		isWanted bool
		errMsg   string
	}

	type test struct {
		name       string
		args       args
		wantResult wantResult
		wantErr    wantErr
	}

	tests := []test{
		{
			name: "test error bad requets success",
			args: args{err: NewWithCode(codes.CodeBadRequest, "bad request")},
			wantResult: wantResult{
				file:    files.GetCurrentFileLocation(), // NOTE: this is line where error created with function NewWithCode
				line:    37,
				message: "bad request",
			},
			wantErr: wantErr{isWanted: false},
		},
		{
			name: "test error failed",
			args: args{err: errors.New("pure error")},
			wantResult: wantResult{
				file:    "",
				line:    0,
				message: "",
			},
			wantErr: wantErr{
				isWanted: true,
				errMsg:   "failed to cast error to stacktrace",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, line, message, err := GetCaller(tt.args.err)
			if tt.wantErr.isWanted {
				if err == nil {
					t.Fatalf("want err is '%v', got error '%v'", tt.wantErr.isWanted, err)
				} else if err.Error() != tt.wantErr.errMsg {
					t.Fatalf("want err msg is '%v' but got '%v'", tt.wantErr.errMsg, err.Error())
				}
			}

			if message != tt.wantResult.message {
				t.Fatalf("want result message is '%v' but got '%v'", tt.wantResult.message, message)
			}

			if line != tt.wantResult.line {
				t.Fatalf("want result line is '%v' but got '%v'", tt.wantResult.line, line)
			}

			if file != tt.wantResult.file {
				t.Fatalf("want result file is '%v' but got '%v'", tt.wantResult.file, file)
			}
		})
	}
}
