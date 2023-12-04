package convert

import (
	"reflect"
	"testing"
)

func Test_ToPointer(t *testing.T) {
	type test struct {
		name           string
		arg            string
		wantResultType reflect.Kind
	}

	tests := []test{
		{
			name:           "test tto ptr string",
			arg:            "a",
			wantResultType: reflect.Pointer,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToPointer(tt.arg)
			if kind := reflect.TypeOf(result).Kind(); kind != tt.wantResultType {
				t.Fatalf("error: want result type is '%v' but got '%v'", tt.wantResultType, kind)
			}
		})
	}
}
