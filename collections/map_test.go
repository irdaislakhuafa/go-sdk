package collections

import (
	"fmt"
	"strings"
	"testing"
)

func Test_Map(t *testing.T) {
	type (
		args struct {
			values []int
		}

		want struct {
			values []string
			err    error
		}

		mode int
		test struct {
			name string
			mode mode
			args args
			want want
		}
	)

	const (
		MODE_MAP = mode(iota)
		MODE_MAP_WITH_ERR
	)

	tests := []test{
		{
			name: "map int to string success",
			mode: MODE_MAP,
			args: args{values: []int{1, 2, 3}},
			want: want{values: []string{"1", "2", "3"}, err: nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.mode {
			case MODE_MAP:
				res := Map(tt.args.values, func(_, value int) string {
					return fmt.Sprint(value)
				})

				if strings.Join(res, ",") != strings.Join(tt.want.values, ",") {
					t.Fatalf("want is %+v but got %+v", tt.want.values, res)
				}
			case MODE_MAP_WITH_ERR:
				res, err := MapWithErr(tt.args.values, func(_, value int) (string, error) {
					return fmt.Sprint(value), nil
				})
				if err != nil {
					t.Fatal(err)
				}

				if strings.Join(res, ",") != strings.Join(tt.want.values, ",") {
					t.Fatalf("want is %+v but got %+v", tt.want.values, res)
				}
			}
		})
	}
}
