package operator

import "testing"

func Test_Ternary(t *testing.T) {
	type test struct {
		name       string
		condition  bool
		ifOk       string
		ifNot      string
		wantResult string
	}
	tests := []test{
		{
			name:       "test ternary if true",
			condition:  true,
			ifOk:       "a",
			ifNot:      "b",
			wantResult: "a",
		},
		{
			name:       "test ternary if false",
			condition:  false,
			ifOk:       "a",
			ifNot:      "b",
			wantResult: "b",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Ternary[string](tt.condition, tt.ifOk, tt.ifNot)
			if result != tt.wantResult {
				t.Fatalf("error: want result '%v' but got '%v'", tt.wantResult, result)
			}
		})
	}
}
