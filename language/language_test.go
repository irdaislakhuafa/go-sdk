package language

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_HTTPStatusText(t *testing.T) {
	type args struct {
		code int
		lang Language
	}

	type test struct {
		name       string
		args       args
		wantResult string
	}

	tests := []test{
		{
			name:       "test indonesian language",
			args:       args{lang: Indonesian, code: http.StatusContinue},
			wantResult: "Lanjutkan",
		},
		{
			name:       "test english language",
			args:       args{lang: English, code: http.StatusAccepted},
			wantResult: "Accepted",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HTTPStatusText(tt.args.lang, tt.args.code)
			if result != tt.wantResult {
				t.Fatalf("want result is '%v' but got '%v'", tt.wantResult, result)
			}
		})
		fmt.Printf("\n")
	}
}
