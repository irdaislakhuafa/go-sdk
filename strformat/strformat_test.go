package strformat

import "testing"

func Test_IsOnlyNumber(t *testing.T) {
	type test struct {
		name       string
		arg        string
		wantResult bool
	}

	tests := []test{
		{
			name:       "test for value xxx",
			arg:        "xxx",
			wantResult: false,
		},
		{
			name:       "test for value 001",
			arg:        "001",
			wantResult: true,
		},
		{
			name:       "test for value 0x1",
			arg:        "0x1",
			wantResult: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsOnlyNumber(tt.arg)
			if result != tt.wantResult {
				t.Fatalf("error: want result is '%v' but got '%v'", tt.wantResult, result)
			}
		})
	}
}
