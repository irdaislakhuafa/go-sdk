package errors

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/files"
	"github.com/irdaislakhuafa/go-sdk/language"
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
				line:    41,
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
				errMsg:   "pure error",
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

func Test_Compile(t *testing.T) {
	type args struct {
		err  error
		lang language.Language
	}

	type wantResult struct {
		statusCode int
		result     App
	}

	type test struct {
		name       string
		args       args
		wantResult wantResult
	}

	tests := []test{
		{
			name: "test auth failure compile success",
			args: args{err: NewWithCode(codes.CodeAuthFailure, "authentication failure"), lang: language.English},
			wantResult: wantResult{
				statusCode: http.StatusUnauthorized,
				result: App{
					Code:  codes.CodeAuthFailure,
					Title: language.HTTPStatusText(language.English, http.StatusUnauthorized),
					Body:  codes.GetCodeMessages(codes.CodeAuthFailure)[language.English].Body,
					sys:   NewWithCode(codes.CodeAuthAccessTokenExpired, "authentication failed"),
				},
			},
		},
		{
			name: "test failed",
			args: args{err: NewWithCode(codes.NoCode, "no code"), lang: language.English},
			wantResult: wantResult{
				statusCode: http.StatusInternalServerError,
				result: App{
					Code:  codes.NoCode,
					Title: language.HTTPStatusText(language.English, http.StatusInternalServerError),
					Body:  codes.GetCodeMessages(codes.NoCode)[language.English].Body,
					sys:   nil,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpStatusCode, app := Compile(tt.args.err, tt.args.lang)
			if httpStatusCode != tt.wantResult.statusCode {
				t.Fatalf("want result status code is '%v' but got '%v'", tt.wantResult.statusCode, httpStatusCode)
			}

			if isCodeEqual := reflect.DeepEqual(app.Code, tt.wantResult.result.Code); !isCodeEqual {
				t.Fatalf("want result app code is '%v' but got '%v'", tt.wantResult.result.Code, app.Code)
			}
		})

		fmt.Printf("result: %+v\n\n", tt.wantResult.result)
	}
}
