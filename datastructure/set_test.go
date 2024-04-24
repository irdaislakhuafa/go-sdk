package datastructure

import (
	"fmt"
	"strings"
	"testing"
)

func Test_Set(t *testing.T) {
	type (
		MODE int
		args struct {
			value  string
			values []string
		}

		want struct {
			contains string
			equals   string
		}

		test struct {
			name       string
			mode       MODE
			args       args
			want       want
			beforeFunc func(s SetInterface[string]) SetInterface[string]
			afterFunc  func(s SetInterface[string]) SetInterface[string]
		}
	)

	const (
		MODE_ADD = iota
		MODE_EXISTS
		MODE_DELETE
		MODE_SLICE
	)

	tests := []test{
		{
			name: "set slice",
			mode: MODE_ADD,
			want: want{contains: "D", equals: strings.Join([]string{"A", "B", "C", "D", "E"}, ",")},
			beforeFunc: func(_ SetInterface[string]) SetInterface[string] {
				return NewSet("A", "B", "C")
			},
		},
		{
			name: "set add",
			mode: MODE_ADD,
			args: args{values: []string{"D", "E"}},
			want: want{contains: "D", equals: strings.Join([]string{"A", "B", "C", "D", "E"}, ",")},
			beforeFunc: func(_ SetInterface[string]) SetInterface[string] {
				return NewSet("A", "B", "C")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := NewSet[string]()
			if tt.beforeFunc != nil {
				set = tt.beforeFunc(set)
			}

			switch tt.mode {
			case MODE_ADD:
				set.Add(tt.args.values...)
				if results := strings.Join(set.Slice(), ","); results != tt.want.equals {
					t.Fatalf("mode add with values '%v' want '%v' but got '%v'", tt.args.values, tt.want.equals, results)
				}
			case MODE_EXISTS:
				isExists := set.IsExists(tt.args.value)
				if fmt.Sprint(isExists) != tt.want.equals {
					t.Fatalf("mode exists with value '%v' want '%v' but got '%v'", tt.args.value, tt.want.equals, isExists)
				}
			case MODE_DELETE:
				set.Del(tt.args.values...)
				if results := strings.Join(set.Slice(), ","); results != tt.want.equals {
					t.Fatalf("mode delete with values '%v' want '%v' but got '%v'", tt.args.values, tt.want.equals, results)
				}
			case MODE_SLICE:
				if results := strings.Join(set.Slice(), ","); results != tt.want.equals {
					t.Fatalf("mode slice want '%v' but got '%v'", tt.want.equals, results)
				}
			}

			if tt.afterFunc != nil {
				set = tt.afterFunc(set)
			}
		})
	}
}
