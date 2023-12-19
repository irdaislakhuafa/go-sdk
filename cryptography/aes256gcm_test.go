package cryptography

import (
	"fmt"
	"testing"
)

func Test_AES256GCM(t *testing.T) {
	type Mode uint
	const (
		MODE_ENCRYPT = Mode(iota)
		MODE_DECRYPT
	)

	type args struct {
		key  string
		text string
	}

	type wantResult struct {
		decrypted          string
		resultEncryptedNot string
	}

	type test struct {
		name       string
		mode       Mode
		args       args
		wantResult wantResult
		isWantErr  bool
	}

	tests := []test{
		{
			name:       "test encrypt success",
			mode:       MODE_ENCRYPT,
			args:       args{key: "7d5b44298bf959af149a0086d79334e6", text: "password"},
			wantResult: wantResult{resultEncryptedNot: "password"},
			isWantErr:  false,
		},
		{
			name:       "test decrypt success",
			mode:       MODE_DECRYPT,
			args:       args{key: "7d5b44298bf959af149a0086d79334e6", text: "NzM5OTNmM2YzMTk5ZTY2YmQ4Zjc1M2QxNjFkZGI0YWRiNWI4NzZiZTA4OTU0MWM5OTgwNzkzYzMwYmMyYjRlMjY1MTM2YmZm"},
			wantResult: wantResult{decrypted: "password"},
			isWantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mode == MODE_ENCRYPT {
				result, err := EncryptAES256GCM([]byte(tt.args.text), []byte(tt.args.key))
				if tt.isWantErr {
					if err == nil {
						t.Fatalf("want err is '%v' but got err '%v'", tt.isWantErr, err)
					}
				} else {
					if err != nil {
						t.Fatalf("want err is '%v' but got err '%v'", tt.isWantErr, err)
					}
				}

				if tt.wantResult.resultEncryptedNot == string(result) {
					t.Fatalf("want encrypted is not '%v' but got '%v'", tt.wantResult.resultEncryptedNot, string(result))
				} else {
					t.Logf("params: %#v, encrypted: %#v", tt.args.text, string(result))
				}
			} else if tt.mode == MODE_DECRYPT {
				result, err := DecryptAES256GCM([]byte(tt.args.text), []byte(tt.args.key))
				if tt.isWantErr {
					if err == nil {
						t.Fatalf("want err is '%v' but got err '%v'", tt.isWantErr, err)
					}
				} else {
					if err != nil {
						t.Fatalf("want err is '%v' but got err '%v'", tt.isWantErr, err)
					}
				}

				if tt.wantResult.decrypted != string(result) {
					t.Fatalf("want decrypted value is '%v' but got '%v'", tt.wantResult.decrypted, string(result))
				} else {
					t.Logf("params: %#v, decrypted: %#v", tt.args.text, string(result))
				}
			}
		})
		fmt.Println("")
	}
}
