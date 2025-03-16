package collections

import (
	"fmt"
	"testing"
)

func Test_Reduce(t *testing.T) {
	type (
		want struct {
			result string
		}
		arg struct {
			items []int
		}
		test struct {
			name string
			arg  arg
			want want
		}
	)

	tests := []test{
		{
			name: "test convert array of int to string",
			arg: arg{
				items: []int{1, 2, 3, 4, 5},
			},
			want: want{
				result: "|1|2|3|4|5|",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Reduce(tt.arg.items, "", func(carry string, idx, value int) string {
				if idx == 0 {
					carry += "|"
				}
				carry += fmt.Sprintf("%v|", value)
				return carry
			})
			if res != tt.want.result {
				t.Fatalf("want is '%v' but got '%v'", tt.want.result, res)
			}
		})
	}
}
