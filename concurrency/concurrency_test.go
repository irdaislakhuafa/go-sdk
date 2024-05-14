package concurrency

import (
	"context"
	"testing"
	"time"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/collections"
	"github.com/irdaislakhuafa/go-sdk/errors"
)

func Test_Concurrency(t *testing.T) {
	type (
		args struct {
			worker  int
			process int
		}

		want struct {
			intValues []int
			errCode   codes.Code
		}

		test struct {
			name string
			args args
			want want
		}
	)

	tests := []test{
		{
			name: "concurrency get data",
			args: args{worker: 2, process: 10},
			want: want{intValues: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, errCode: codes.NoCode},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intValues := []int{}
			ctx := context.Background()
			c := NewConcurrency().WithMaxWorker(int64(tt.args.worker))
			for index := range tt.want.intValues[:len(tt.want.intValues)/2] {
				i := tt.want.intValues[index]
				c.AddFunc(func(ctx context.Context, c Interface) {
					time.Sleep(time.Second * 3)
					c.Lock()
					intValues = append(intValues, i)
					t.Logf("%#v", intValues)
					c.Unlock()
				})
			}
			if err := c.Do(ctx); err != nil {
				if code := errors.GetCode(err); code.IsNotOneOf(tt.want.errCode) {
					t.Fatalf("want err code '%v' but got '%v'", tt.want.errCode, errors.GetCode(err))
				}
			}

			for index := range tt.want.intValues[len(tt.want.intValues)/2:] {
				i := tt.want.intValues[len(tt.want.intValues)/2+index]
				c.AddFunc(func(ctx context.Context, c Interface) {
					time.Sleep(time.Second * 3)
					c.Lock()
					intValues = append(intValues, i)
					t.Logf("%#v", intValues)
					c.Unlock()
				})
			}

			if err := c.Do(ctx); err != nil {
				if code := errors.GetCode(err); code.IsNotOneOf(tt.want.errCode) {
					t.Fatalf("want err code '%v' but got '%v'", tt.want.errCode, errors.GetCode(err))
				}
			}

			if !collections.IsElementsEquals(tt.want.intValues, intValues) {
				t.Fatalf("want result values is '%#v' but got '%#v'", tt.want.intValues, intValues)
			}
		})
	}
}
