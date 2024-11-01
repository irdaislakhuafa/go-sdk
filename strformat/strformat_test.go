package strformat

import (
	"fmt"
	"testing"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/irdaislakhuafa/go-sdk/files"
)

func Test_IsOnlyNumber(t *testing.T) {
	type test struct {
		name       string
		arg        string
		wantResult bool
	}

	tests := []test{
		{
			name:       "test for value xxx",
			arg:        "xxx",
			wantResult: false,
		},
		{
			name:       "test for value 001",
			arg:        "001",
			wantResult: true,
		},
		{
			name:       "test for value 0x1",
			arg:        "0x1",
			wantResult: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsOnlyNumber(tt.arg)
			if result != tt.wantResult {
				t.Fatalf("error: want result is '%v' but got '%v'", tt.wantResult, result)
			}
		})
	}
}

func Test_Tmpl(t *testing.T) {
	type (
		args struct {
			value   any
			strTmpl string
		}

		want struct {
			result string
		}

		wantErr struct {
			code codes.Code
		}

		test struct {
			name      string
			args      args
			want      want
			wantErr   wantErr
			isWantErr bool
		}
	)

	tests := []test{
		{
			name: "success",
			args: args{
				strTmpl: "Hello, my name is {{.Name}}. My age is {{.Age}}",
				value: map[string]any{
					"Name": "Irda Islakhu Afa",
					"Age":  22,
				},
			},
			want: want{
				result: "Hello, my name is Irda Islakhu Afa. My age is 22",
			},
			wantErr:   wantErr{},
			isWantErr: false,
		},
	}

	f := files.GetCurrentMethodName()
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v:%v", f, tt.name), func(t *testing.T) {
			result, err := T(tt.args.strTmpl, tt.args.value)
			if tt.isWantErr {
				if err == nil {
					t.Fatalf("want err is %#v but got err with msg %#v", tt.isWantErr, err)
				} else {
					if code := errors.GetCode(err); code != tt.wantErr.code {
						t.Fatalf("want err code is %#v but got err code %#v", tt.wantErr.code, code)
					}
				}
			}

			if result != tt.want.result {
				t.Fatalf("want result is %#v but got %#v", tt.want.result, result)
			}
		})
	}
}
