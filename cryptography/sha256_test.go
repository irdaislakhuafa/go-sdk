package cryptography

import (
	"fmt"
	"testing"

	"github.com/irdaislakhuafa/go-sdk/files"
)

func Test_NewSHA256(t *testing.T) {
	type args struct {
		text []byte
		key  []byte
	}

	type want struct {
		result string
	}

	type test struct {
		name string
		args args
		want want
	}

	tests := []test{
		{
			name: "test hash without key success",
			args: args{text: []byte("password")},
			want: want{result: "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"},
		},
		{
			name: "test hash with key success",
			args: args{text: []byte("password"), key: []byte("key")},
			want: want{result: "4d42fb9ffc8d7d0a245429438b4bc73db1007a167026a0a0c6a74fa58e8e86ca"},
		},
	}

	f := files.GetCurrentMethodName()
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v: %v", f, tt.name), func(t *testing.T) {
			sha256 := NewSHA256([]byte(tt.args.text))
			if tt.args.key != nil {
				sha256 = sha256.WithKey(tt.args.key)
			}

			result, err := sha256.Build()
			if err != nil {
				t.Fatalf("failed generate hash, %v", err)
			}

			if result != tt.want.result {
				t.Fatalf("want result is %#v but got %#v", tt.want.result, result)
			}
		})
		fmt.Println("")
	}
}
