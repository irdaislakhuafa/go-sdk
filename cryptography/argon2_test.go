package cryptography

import (
	"fmt"
	"strings"
	"testing"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/irdaislakhuafa/go-sdk/files"
)

func Test_Argon2(t *testing.T) {
	type Mode int
	const (
		MODE_HASH = Mode(iota)
		MODE_COMPARE
	)

	type param struct {
		plainText  string
		hashedText string
	}

	type wantResult struct {
		isEqual  bool
		isArgon2 bool
	}

	type wantErr struct {
		errCode codes.Code
	}

	type test struct {
		name       string
		mode       Mode
		algFunc    Argon2
		param      param
		isWantErr  bool
		wantResult wantResult
		wantErr    wantErr
	}

	tests := []test{
		{
			name:       "test hash success",
			mode:       MODE_HASH,
			algFunc:    NewArgon2(),
			param:      param{plainText: "password"},
			isWantErr:  false,
			wantErr:    wantErr{},
			wantResult: wantResult{isArgon2: true},
		},
		{
			name:       "compare hash success",
			mode:       MODE_COMPARE,
			algFunc:    NewArgon2(),
			param:      param{plainText: "password", hashedText: "$argon2id$v=19$m=4096,t=3,p=1$5Jr6f1VGhrzIRCfUu8VlBQ$xhbrJs0FhNIBR+/vvCkUxcrcDZc0G2xc3ipqstj9H3A"},
			isWantErr:  false,
			wantResult: wantResult{isEqual: true},
			wantErr:    wantErr{},
		},
		{
			name:       "compare hash failed or false",
			mode:       MODE_COMPARE,
			algFunc:    NewArgon2(),
			param:      param{plainText: "x", hashedText: "$argon2id$v=19$m=4096,t=3,p=1$5Jr6f1VGhrzIRCfUu8VlBQ$xhbrJs0FhNIBR+/vvCkUxcrcDZc0G2xc3ipqstj9H3A"},
			isWantErr:  false,
			wantResult: wantResult{isEqual: false},
			wantErr:    wantErr{},
		},
		{
			name:       "compatibility test ok with hashed text from https://github.com/ranisalt/node-argon2.git",
			mode:       MODE_COMPARE,
			algFunc:    NewArgon2().SetIterations(3).SetParallelism(4).SetMemory(65536),
			param:      param{plainText: "password", hashedText: "$argon2id$v=19$m=65536,t=3,p=4$AL7pBy/YmGyqSOH4/LCWMQ$Z7Pan4USduRrdZkYIIVvRBjbYwT+pDAszV/qqYt3Y+A"},
			isWantErr:  false,
			wantResult: wantResult{isEqual: true},
			wantErr:    wantErr{},
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
						t.Fatalf("want err is %#v but got err %#v", tt.isWantErr, err)
					}

					if code := errors.GetCode(err); tt.wantErr.errCode != code {
						t.Fatalf("want err code is %#v but got %#v with msg %#v", tt.wantErr.errCode, code, err)
					}
				} else {
					if strings.Contains(result, "$argon2id$v=") == tt.wantResult.isArgon2 {
						t.Logf("result is argon2 %#v", result)
					}
				}

			case MODE_COMPARE:
				isOk, err := tt.algFunc.Compare([]byte(tt.param.plainText), []byte(tt.param.hashedText))
				if tt.isWantErr {
					if err == nil {
						t.Fatalf("want err is %#v but got err %#v", tt.isWantErr, err)
					}

					if code := errors.GetCode(err); tt.wantErr.errCode != code {
						t.Fatalf("want err code is %#v but got %#v with msg %#v", tt.wantErr.errCode, code, err)
					}
				} else {
					// if tt.wantResult.isEqual {
					if isOk != tt.wantResult.isEqual {
						t.Fatalf("want result is %#v for params %#v == %#v", tt.wantResult.isEqual, tt.param.plainText, tt.param.hashedText)
					} else {
						t.Logf("compare result for params %#v = %#v is %#v want is %#v", tt.param.hashedText, tt.param.plainText, isOk, tt.wantResult.isEqual)
					}
					// } else {
					// 	if isOk == tt.wantResult.isEqual {
					// 		t.Fatalf("want result is %#v for params %#v = %#v but got %#v", tt.wantResult.isEqual, tt.param.hashedText, tt.param.plainText, isOk)
					// 	}
					// }
				}
			}
		})
		fmt.Println("")
	}
}
