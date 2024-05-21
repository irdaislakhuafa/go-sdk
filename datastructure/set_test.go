package datastructure

import (
	"fmt"
	"strings"
	"testing"

	"github.com/irdaislakhuafa/go-sdk/collections"
)

func Test_Set(t *testing.T) {
	type (
		MODE int
		args struct {
			value  string
			values []string
		}

		want struct {
			equals string
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
		MODE_DELETE_IF
		MODE_FILTER
	)

	tests := []test{
		{
			name: "set slice",
			mode: MODE_SLICE,
			want: want{equals: strings.Join([]string{"A", "B", "C"}, ",")},
			beforeFunc: func(_ SetInterface[string]) SetInterface[string] {
				return NewSet("A", "B", "C")
			},
		},
		{
			name: "set add",
			mode: MODE_ADD,
			args: args{values: []string{"D", "E"}},
			want: want{equals: strings.Join([]string{"A", "B", "C", "D", "E"}, ",")},
			beforeFunc: func(_ SetInterface[string]) SetInterface[string] {
				return NewSet("A", "B", "C")
			},
		},
		{
			name: "set delete",
			mode: MODE_DELETE,
			args: args{values: []string{"A", "B"}},
			want: want{equals: strings.Join([]string{"C"}, ",")},
			beforeFunc: func(_ SetInterface[string]) SetInterface[string] {
				return NewSet("A", "B", "C")
			},
		},
		{
			name: "set exists",
			mode: MODE_EXISTS,
			args: args{value: "A"},
			want: want{equals: "true"},
			beforeFunc: func(_ SetInterface[string]) SetInterface[string] {
				return NewSet("A", "B", "C")
			},
		},
		{
			name: "delete if",
			mode: MODE_DELETE_IF,
			args: args{value: "A"},
			want: want{equals: "A,B"},
			beforeFunc: func(_ SetInterface[string]) SetInterface[string] {
				s := NewSet("A", "B", "C")
				s.DelIf(func(value string) bool { return value == "C" })
				return s
			},
		},
		{
			name: "filter",
			mode: MODE_FILTER,
			args: args{value: "A"},
			want: want{equals: "A"},
			beforeFunc: func(_ SetInterface[string]) SetInterface[string] {
				s := NewSet("A", "B", "C")
				s = s.Filter(func(value string) bool { return value == "A" })
				return s
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
				if want := strings.Split(tt.want.equals, ","); !collections.IsElementsEquals(want, set.Slice()) {
					t.Fatalf("mode add with values '%v' want '%v' but got '%v'", tt.args.values, want, set.Slice())
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
				if want := strings.Split(tt.want.equals, ","); !collections.IsElementsEquals(want, set.Slice()) {
					t.Fatalf("mode slice want '%v' but got '%v'", want, set.Slice())
				}
			case MODE_DELETE_IF:
				if want := strings.Split(tt.want.equals, ","); !collections.IsElementsEquals(want, set.Slice()) {
					t.Fatalf("mode delete if want '%v' but got '%v'", want, set.Slice())
				}
			case MODE_FILTER:
				if want := strings.Split(tt.want.equals, ","); !collections.IsElementsEquals(want, set.Slice()) {
					t.Fatalf("mode filter want '%v' but got '%v'", want, set.Slice())
				}
			}

			if tt.afterFunc != nil {
				set = tt.afterFunc(set)
			}
		})
	}
}
