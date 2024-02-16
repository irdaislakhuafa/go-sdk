package codes

import (
	"fmt"
	"testing"

	"github.com/irdaislakhuafa/go-sdk/files"
)

func Test_CodeInterface(t *testing.T) {
	type Func uint64
	const (
		IsBetween = Func(iota + 1)
		IsOneOf
		IsNotOneOf
	)

	type (
		params struct {
			code      Code
			listCodes []Code
		}

		want struct {
			isOk bool
		}

		test struct {
			name     string
			funcName Func
			params   params
			want     want
		}
	)

	tests := []test{
		{
			name:     "IsBetween() true",
			funcName: IsBetween,
			params:   params{listCodes: []Code{CodeSQLStart, CodeSQLEnd}, code: CodeSQLRead},
			want:     want{isOk: true},
		},
		{
			name:     "IsBetween() false",
			funcName: IsBetween,
			params:   params{listCodes: []Code{CodeSMTPStart, CodeSMTPEnd}, code: CodeSQLRead},
			want:     want{isOk: false},
		},
		{
			name:     "IsOneOf() true",
			funcName: IsOneOf,
			params:   params{listCodes: []Code{CodeSMTPStart, CodeSMTPEnd, CodeSQLRead}, code: CodeSQLRead},
			want:     want{isOk: true},
		},
		{
			name:     "IsOneOf() false",
			funcName: IsOneOf,
			params:   params{listCodes: []Code{CodeSMTPStart, CodeSMTPEnd}, code: CodeSQLRead},
			want:     want{isOk: false},
		},
		{
			name:     "IsNotOneOf() true",
			funcName: IsNotOneOf,
			params:   params{listCodes: []Code{CodeSMTPStart, CodeSMTPEnd}, code: CodeSQLRead},
			want:     want{isOk: true},
		},
		{
			name:     "IsNotOneOf() false",
			funcName: IsNotOneOf,
			params:   params{listCodes: []Code{CodeSMTPStart, CodeSMTPEnd, CodeSQLRead}, code: CodeSQLRead},
			want:     want{isOk: false},
		},
	}

	f := files.GetCurrentMethodName()
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v:%v", f, tt.name), func(t *testing.T) {
			switch tt.funcName {
			case IsBetween:
				if result := tt.params.code.IsBetween(tt.params.listCodes[0], tt.params.listCodes[1]); tt.want.isOk != result {
					t.Fatalf("want result is %#v but got %#v", tt.want.isOk, result)
				}
			case IsOneOf:
				if result := tt.params.code.IsOneOf(tt.params.listCodes...); tt.want.isOk != result {
					t.Fatalf("want result is %#v but got %#v", tt.want.isOk, result)
				}
			case IsNotOneOf:
				if result := tt.params.code.IsNotOneOf(tt.params.listCodes...); tt.want.isOk != result {
					t.Fatalf("want result is %#v but got %#v", tt.want.isOk, result)
				}
			}
		})
	}
}
