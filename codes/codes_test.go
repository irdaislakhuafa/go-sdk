package codes

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/irdaislakhuafa/go-sdk/language"
)

func Test_Compile(t *testing.T) {
	type args struct {
		code Code
		lang language.Language
	}

	type test struct {
		name       string
		args       args
		wantResult Message
	}

	tests := []test{
		{
			name:       "default success message in english language",
			args:       args{code: CodeSuccess, lang: language.English},
			wantResult: Message{StatusCode: http.StatusOK, Title: "OK", Body: "Request successful"},
		},
		{
			name:       "test success in indonesian language",
			args:       args{code: CodeSuccess, lang: language.Indonesian},
			wantResult: Message{StatusCode: http.StatusOK, Title: "OK", Body: "Request berhasil"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Compile(tt.args.code, tt.args.lang)
			if isEquals := reflect.DeepEqual(result, tt.wantResult); !isEquals {
				t.Fatalf("want result is '%+v' but got '%+v'", tt.wantResult, result)
			}
		})

		fmt.Printf("\n")
	}
}
