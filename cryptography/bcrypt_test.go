package cryptography

import (
	"fmt"
	"testing"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/irdaislakhuafa/go-sdk/files"
)

func Test_Bcrypt(t *testing.T) {
	type Mode int

	const (
		MODE_HASH = iota
		MODE_COMPARE
	)

	type params struct {
		plainText  string
		hashedText string
	}

	type wantResult struct {
		isEqual bool
	}

	type wantErr struct {
		errCode codes.Code
	}

	type test struct {
		name       string
		algFunc    Bcrypt
		mode       Mode
		param      params
		isWantErr  bool
		wantResult wantResult
		wantErr    wantErr
	}

	tests := []test{
		{
			name:       "test bcrypt success",
			algFunc:    NewBcrypt(),
			mode:       MODE_HASH,
			param:      params{plainText: "password"},
			isWantErr:  false,
			wantResult: wantResult{},
			wantErr:    wantErr{},
		},
		{
			name:       "test bcrypt success with cost 10",
			algFunc:    NewBcrypt().SetCost(10),
			mode:       MODE_HASH,
			param:      params{plainText: "password"},
			isWantErr:  false,
			wantResult: wantResult{},
			wantErr:    wantErr{},
		},
		{
			name:       "test compare success",
			algFunc:    NewBcrypt(),
			mode:       MODE_COMPARE,
			param:      params{plainText: "password", hashedText: "$2a$10$3aElHMD12UI9BAhXSSwCqO0QhpUmy9koayt6EEz.N4SV6CeSJcJMu"},
			isWantErr:  false,
			wantResult: wantResult{isEqual: true},
			wantErr:    wantErr{},
		},
		{
			name:       "test compare success compatible with nodejs lib github.com/kelektiv/node.bcrypt.js",
			algFunc:    NewBcrypt(),
			mode:       MODE_COMPARE,
			param:      params{plainText: "password", hashedText: "$2b$10$IBbySgCR.aXyx3ddeY4LNuoFP1QeqeJA36Y6RAz3gBz1pIpDyAayS"},
			isWantErr:  false,
			wantResult: wantResult{isEqual: true},
			wantErr:    wantErr{},
		},
		{
			name:       "test compare not match",
			algFunc:    NewBcrypt(),
			mode:       MODE_COMPARE,
			param:      params{plainText: "passwords", hashedText: "$2a$10$3aElHMD12UI9BAhXSSwCqO0QhpUmy9koayt6EEz.N4SV6CeSJcJMu"},
			isWantErr:  true,
			wantResult: wantResult{isEqual: false},
			wantErr:    wantErr{errCode: codes.CodeBcrypt},
		},
	}

	f := files.GetCurrentMethodName()
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v:%v:%v", f, tt.mode, tt.name), func(t *testing.T) {
			switch tt.mode {
			case MODE_HASH:
				result, err := tt.algFunc.Hash([]byte(tt.param.plainText))
				if tt.isWantErr {
					if err == nil {
						t.Fatalf("is want err %#v but got err msg %#v", tt.isWantErr, err)
					}

					if code := errors.GetCode(err); code != tt.wantErr.errCode {
						t.Fatalf("want err code is %#v but get err code %#v with msg %#v", tt.wantErr.errCode, code, err)
					}
				} else {
					if result := string(result); result != tt.param.plainText || result != "" {
						t.Logf("result bcrypt is %#v using plain text %#v", result, tt.param.plainText)
					} else {
						t.Fatalf("result bcrypt %#v for plain text %v", result, tt.param.plainText)
					}
				}
			case MODE_COMPARE:
				err := tt.algFunc.Compare([]byte(tt.param.plainText), []byte(tt.param.hashedText))
				if tt.isWantErr {
					if err == nil {
						t.Fatalf("is want err %#v but got err msg %#v", tt.isWantErr, err)
					}

					if code := errors.GetCode(err); code != tt.wantErr.errCode {
						t.Fatalf("want err code is %#v but get err code %#v with msg %#v", tt.wantErr.errCode, code, err)
					}
				} else {
					if isEqual := (err == nil); isEqual != tt.wantResult.isEqual {
						t.Fatalf("hashed text %#v is not equals with %#v", tt.param.hashedText, tt.param.plainText)
					} else {
						t.Logf("hashed text %#v is equals with %#v", tt.param.hashedText, tt.param.plainText)
					}
				}
			}
		})
		fmt.Println("")
	}
}
