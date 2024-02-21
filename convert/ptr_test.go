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

func Test_ToSafeValue(t *testing.T) {
	type DataType int
	const (
		INT = DataType(iota + 1)
		FLOAT
		STRING
		STRUCT
	)

	type (
		exampleStruct struct {
			Anything string
		}

		params struct {
			dataType DataType
			value    any
		}

		want struct {
			value any
		}

		test struct {
			name   string
			params params
			want   want
		}
	)
	tests := []test{
		{
			name:   "int nil ptr",
			params: params{dataType: INT, value: nil},
			want:   want{value: 0},
		},
		{
			name:   "int ptr value 10",
			params: params{dataType: INT, value: ToPointer(10)},
			want:   want{value: 10},
		},
		{
			name:   "string ptr",
			params: params{dataType: STRING, value: nil},
			want:   want{value: ""},
		},
		{
			name:   "string ptr value 'str'",
			params: params{dataType: STRING, value: ToPointer("str")},
			want:   want{value: "str"},
		},
		{
			name:   "struct ptr",
			params: params{dataType: STRUCT, value: nil},
			want:   want{value: exampleStruct{}},
		},
		{
			name:   "struct ptr value 'str'",
			params: params{dataType: STRUCT, value: ToPointer(exampleStruct{Anything: "haha"})},
			want:   want{value: exampleStruct{Anything: "haha"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result any
			switch tt.params.dataType {
			case INT:
				result = ToSafeValue[int](tt.params.value)
			case STRING:
				result = ToSafeValue[string](tt.params.value)
			case FLOAT:
				result = ToSafeValue[float64](tt.params.value)
			case STRUCT:
				result = ToSafeValue[exampleStruct](tt.params.value)
			default:
				t.Fatalf("data type not supported!")
			}

			if result != tt.want.value {
				t.Fatalf("want result is '%v' but got '%v'", tt.want.value, result)
			}
		})
	}
}
