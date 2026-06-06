package storage

import "testing"

func TestHelperFormatBytes(t *testing.T) {
	type testCase struct {
		name string
		args uint64
		want string
	}

	testCases := []testCase{
		{
			name: "zero",
			args: 0,
			want: "0.00 B",
		},
		{
			name: "bytes",
			args: 1023,
			want: "1023.00 B",
		},
		{
			name: "bytes",
			args: 1024,
			want: "1.00 KB",
		},
		{
			name: "bytes",
			args: 1024 * 1024,
			want: "1.00 MB",
		},
		{
			name: "bytes",
			args: 1024 * 1024 * 1024,
			want: "1.00 GB",
		},
		{
			name: "bytes",
			args: 1024 * 1024 * 1024 * 1024,
			want: "1.00 TB",
		},
		{
			name: "bytes",
			args: 1024 * 1024 * 1024 * 1024 * 1024,
			want: "1.00 PB",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// t.Parallel()
			result := FormatBytes(tc.args)
			if result != tc.want {
				t.Fatalf("want '%v' but got '%v'", tc.want, result)
			}
		})
	}
}

func TestHelperPercentBytes(t *testing.T) {
	type (
		args struct {
			used  uint64
			total uint64
		}

		testCase struct {
			name string
			args args
			want float64
		}
	)

	testCases := []testCase{
		{
			name: "used 0B from 0MB",
			args: args{
				used:  0,
				total: 0,
			},
			want: 100.00,
		},
		{
			name: "used 0B from 10MB",
			args: args{
				used:  0,
				total: 1024 * 1024 * 10,
			},
			want: 0.00,
		},
		{
			name: "used 1MB from 10MB",
			args: args{
				used:  1024 * 1024,
				total: 1024 * 1024 * 10,
			},
			want: 10.00,
		},
		{
			name: "used 5MB from 10MB",
			args: args{
				used:  1024 * 1024 * 5,
				total: 1024 * 1024 * 10,
			},
			want: 50.00,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := PercentBytes(tc.args.used, tc.args.total)
			if result != tc.want {
				t.Fatalf("want '%.2f' but got '%.2f'", tc.want, result)
			}
		})
	}
}
