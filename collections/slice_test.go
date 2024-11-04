package collections

import "testing"

func Test_IsElementEquals(t *testing.T) {
	type (
		params struct {
			param1, param2 []string
		}
		test struct {
			name       string
			params     params
			wantResult bool
		}
	)

	tests := []test{
		{
			name: "element should equals",
			params: params{
				param1: []string{"A", "B", "C"},
				param2: []string{"B", "C", "A"},
			},
			wantResult: true,
		},
		{
			name: "element should not equals",
			params: params{
				param1: []string{"A", "B", "C"},
				param2: []string{"A", "B"},
			},
			wantResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsElementsEquals(tt.params.param1, tt.params.param2)
			if tt.wantResult != result {
				t.Fatalf("param1: '%+v', param2: '%+v', want result '%+v' but got '%+v'\n", tt.params.param1, tt.params.param2, tt.wantResult, result)
			}
		})
	}
}
