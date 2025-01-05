package caches

import (
	"fmt"
	"testing"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
)

func TestCaches(t *testing.T) {
	type FN int
	const (
		FN_REMEMBER = iota
		FN_CLEAR
		FN_FORGET
	)

	type (
		args struct {
			key      string
			ttl      uint64
			callback func() string
		}

		params struct {
			cache Interface[string]
			args  args
		}

		want struct {
			equals string
		}

		test struct {
			name     string
			fn       FN
			params   params
			want     want
			fnBefore func(test *test)
			fnAfter  func(test *test)
		}
	)

	tests := []test{
		{
			name: "remember 3 seconds success",
			params: params{
				cache: NewCache[string](),
				args: args{
					key: "key_3s",
					ttl: 3,
					callback: func() string {
						fmt.Println("callback called for key_3s") // should be called only once because cache will remember value for 3 seconds
						return "value_3s"
					},
				},
			},
			fn: FN_REMEMBER,
			want: want{
				equals: "value_3s",
			},
			fnBefore: func(test *test) {},
			fnAfter:  func(test *test) {},
		},
		{
			name: "forgot key_3s success",
			params: params{
				cache: NewCache[string](),
				args: args{
					key: "key_3s",
					ttl: 3,
					callback: func() string {
						fmt.Println("callback called for key_3s")
						return "value_3s"
					},
				},
			},
			fn: FN_FORGET,
			want: want{
				equals: "value_3s",
			},
			fnBefore: func(test *test) {
				tt := test
				tt.params.cache.Remember(tt.params.args.key, tt.params.args.ttl, tt.params.args.callback)
			},
			fnAfter: func(test *test) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fnBefore != nil {
				tt.fnBefore(&tt)
			}

			switch tt.fn {
			case FN_REMEMBER:
				value, err := tt.params.cache.Remember(tt.params.args.key, tt.params.args.ttl, tt.params.args.callback)
				if err != nil {
					t.Fatal(err)
				}

				if value != tt.want.equals {
					t.Fatalf("want '%v' but got '%v'", tt.want.equals, value)
				}

				value, err = tt.params.cache.Remember(tt.params.args.key, tt.params.args.ttl, tt.params.args.callback)
				if err != nil {
					t.Fatal(err)
				}

				if value != tt.want.equals {
					t.Fatalf("want '%v' but got '%v'", tt.want.equals, value)
				}
			case FN_CLEAR:
				tt.params.cache.Clear()
			case FN_FORGET:
				value, err := tt.params.cache.Forget(tt.params.args.key)
				if err != nil {
					t.Fatal(err)
				}

				if value != tt.want.equals {
					t.Fatalf("want '%v' but got '%v'", tt.want.equals, value)
				}

				_, err = tt.params.cache.Forget(tt.params.args.key)
				if code := errors.GetCode(err); code != codes.CodeNotFound {
					t.Fatalf("want '%v' but got '%v'", codes.CodeNotFound, code)
				}
			}

			if tt.fnAfter != nil {
				tt.fnAfter(&tt)
			}
		})
	}

}
