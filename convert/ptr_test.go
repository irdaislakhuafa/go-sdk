package convert

import (
	"database/sql"
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
			name:   "int ptr ptr value 10",
			params: params{dataType: INT, value: ToPointer(ToPointer(10))},
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
			name:   "struct ptr value",
			params: params{dataType: STRUCT, value: ToPointer(exampleStruct{Anything: "haha"})},
			want:   want{value: exampleStruct{Anything: "haha"}},
		},
		{
			name:   "struct ptr value empty str",
			params: params{dataType: STRUCT, value: ToPointer("")},
			want:   want{value: exampleStruct{}},
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

func Test_SQLNullToNil(t *testing.T) {
	type (
		args struct {
			typeData reflect.Kind
			snull    sql.NullString
			inull    sql.NullInt16
			bnull    sql.NullBool
			fnull    sql.NullFloat64
		}
		want struct {
			valid  bool
			result any
		}
		test struct {
			name string
			args args
			want want
		}
	)

	tests := []test{
		{
			name: "success test int",
			args: args{
				typeData: reflect.Int,
				inull: sql.NullInt16{
					Int16: 100,
					Valid: true,
				},
			},
			want: want{
				valid:  true,
				result: 100,
			},
		},
		{
			name: "failed test int",
			args: args{
				typeData: reflect.Int,
				inull: sql.NullInt16{
					Int16: 100,
					Valid: false,
				},
			},
			want: want{
				valid:  false,
				result: nil,
			},
		},
		{
			name: "success test bool",
			args: args{
				typeData: reflect.Bool,
				bnull: sql.NullBool{
					Bool:  true,
					Valid: true,
				},
			},
			want: want{
				valid:  true,
				result: true,
			},
		},
		{
			name: "success test float64",
			args: args{
				typeData: reflect.Float64,
				fnull: sql.NullFloat64{
					Float64: 100.00,
					Valid:   true,
				},
			},
			want: want{
				valid:  true,
				result: 100.00,
			},
		},
		{
			name: "success test string",
			args: args{
				typeData: reflect.String,
				snull: sql.NullString{
					String: "100",
					Valid:  true,
				},
			},
			want: want{
				valid:  true,
				result: "100",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.args.typeData {
			case reflect.Int:
				res := SQLNullInt16ToNil(tt.args.inull)
				if tt.want.valid {
					if int(*res) != tt.want.result {
						t.Fatalf("want is '%v' but got '%v'", tt.want.result, *res)
					}
				} else {
					if res != nil {
						t.Fatalf("want is '%v' but got '%v'", tt.want.result, res)
					}
				}
			case reflect.Float64:
				res := SQLNullFloat64ToNil(tt.args.fnull)
				if tt.want.valid {
					if *res != tt.want.result {
						t.Fatalf("want is '%v' but got '%v'", tt.want.result, *res)
					}
				} else {
					if res != nil {
						t.Fatalf("want is '%v' but got '%v'", tt.want.result, res)
					}
				}
			case reflect.Bool:
				res := SQLNullBoolToNil(tt.args.bnull)
				if tt.want.valid {
					if *res != tt.want.result {
						t.Fatalf("want is '%v' but got '%v'", tt.want.result, *res)
					}
				} else {
					if res != nil {
						t.Fatalf("want is '%v' but got '%v'", tt.want.result, res)
					}
				}
			case reflect.String:
				res := SQLNullStringToNil(tt.args.snull)
				if tt.want.valid {
					if *res != tt.want.result {
						t.Fatalf("want is '%v' but got '%v'", tt.want.result, *res)
					}
				} else {
					if res != nil {
						t.Fatalf("want is '%v' but got '%v'", tt.want.result, res)
					}
				}
			}
		})
	}
}
