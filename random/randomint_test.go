package random

import (
	"fmt"
	"testing"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
)

func Test_Int(t *testing.T) {
	type (
		args struct {
			maxLength int
		}

		result struct {
			resultLength int
		}

		err struct {
			code codes.Code
		}

		want struct {
			result result
			err    err
		}

		test struct {
			name string
			args args
			want want
		}
	)

	tests := []test{
		{
			name: "success generate random int 4 max length",
			args: args{
				maxLength: 4,
			},
			want: want{
				result: result{
					resultLength: 4,
				},
				err: err{},
			},
		},
		{
			name: "failed invalid value generate random int 20 max length",
			args: args{
				maxLength: 20,
			},
			want: want{
				result: result{
					resultLength: 20,
				},
				err: err{
					code: codes.CodeInvalidValue,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Int(tt.args.maxLength)
			if err != nil {
				if code := errors.GetCode(err); code.IsOneOf(tt.want.err.code) {
					t.Log("wanted error code is match!")
				} else {
					t.Fatalf("wanted error code is '%v' but got '%v'", tt.want.err.code, code)
				}
			}

			t.Log(res)
			if l := len(fmt.Sprint(res)); l == tt.want.result.resultLength {
				t.Logf("wanted result '%v' length is match", l)
			} else if l != tt.want.result.resultLength && tt.want.err.code == codes.NoCode {
				t.Fatalf("wanted result length is '%v' but got '%v'", tt.want.result.resultLength, l)
			}
		})
	}
}

func Test_SInt(t *testing.T) {
	type (
		args struct {
			maxLength int
		}

		result struct {
			maxLength int
		}

		err struct {
			code codes.Code
		}

		want struct {
			result result
			err    err
		}

		test struct {
			name string
			args args
			want want
		}
	)

	tests := []test{
		{
			name: "success generate random string of int 4 max length",
			args: args{
				maxLength: 4,
			},
			want: want{
				result: result{
					maxLength: 4,
				},
				err: err{},
			},
		},
		{
			name: "success generate random string of int 30 max length",
			args: args{
				maxLength: 30,
			},
			want: want{
				result: result{
					maxLength: 30,
				},
				err: err{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := SInt(tt.args.maxLength)
			if err != nil {
				if code := errors.GetCode(err); code.IsOneOf(tt.want.err.code) {
					t.Log("wanted error code is match")
				} else {
					t.Fatalf("wanted error code is '%v' but got '%v'", tt.want.err.code, code)
				}
			}

			if l := len(res); l == tt.want.result.maxLength {
				t.Logf("wanted result '%v' length is match", l)
			} else {
				t.Fatalf("wanted result length is '%v' but got '%v'", tt.want.result.maxLength, l)
			}
		})
	}
}
