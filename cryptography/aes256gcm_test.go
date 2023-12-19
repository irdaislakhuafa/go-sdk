package cryptography

import (
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
				}
			}

		})
	}
}
