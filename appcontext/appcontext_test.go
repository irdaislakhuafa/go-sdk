package appcontext

import (
	"context"
	"reflect"
	"testing"

	"github.com/irdaislakhuafa/go-sdk/language"
)

func Test_Appcontext_acceptLanguage(t *testing.T) {
	const (
		SET = int(iota)
		GET
	)

	type args struct {
		ctx  context.Context
		lang language.Language
	}

	type want struct {
		value any
	}

	type test struct {
		name string
		mode int
		args args
		want want
	}

	tests := []test{
		{
			name: "test SetAcceptLanguage success",
			mode: SET,
			args: args{ctx: context.Background(), lang: language.Indonesian},
			want: want{value: context.WithValue(context.Background(), acceptLanguage, language.Indonesian)},
		},
		{
			name: "test GetAcceptLanguage success with default language",
			mode: GET,
			args: args{ctx: context.Background()},
			want: want{value: language.English},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.mode {
			case SET:
				if ctx := SetAcceptLanguage(tt.args.ctx, tt.args.lang); !reflect.DeepEqual(ctx, tt.want.value) {
					t.Fatalf("want ctx is '%+v' but got '%+v'", tt.want.value, ctx)
				}
			case GET:
				if ctx := GetAcceptLanguage(tt.args.ctx); !reflect.DeepEqual(ctx, tt.want.value) {
					t.Fatalf("want ctx is '%+v' but got '%+v'", tt.want.value, ctx)
				}
			}
		})
	}
}
