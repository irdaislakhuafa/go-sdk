package auth

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/irdaislakhuafa/go-sdk/files"
)

func Test_JWT(t *testing.T) {
	type Mode int

	const (
		MODE_GENERATE = Mode(iota + 1)
		MODE_VALIDATE
		MODE_EXTRACT
	)

	type claims struct {
		UserID string
		jwt.RegisteredClaims
	}

	type params struct {
		claims      claims
		tokenString string
		secretKey   string
	}

	type want struct {
		fn func(token string) error
	}

	type wantErr struct {
		code codes.Code
	}

	type test struct {
		ctx        context.Context
		beforeFunc func(ctx context.Context, j JWTInterface[*claims], p *params)
		want       want
		params     params
		name       string
		mode       Mode
		wantErr    wantErr
		isWantErr  bool
	}

	tests := []test{
		{
			ctx:  context.Background(),
			name: "generate jwt token string",
			mode: MODE_GENERATE,
			params: params{
				claims: claims{
					uuid.NewString(),
					jwt.RegisteredClaims{
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
					},
				},
				secretKey: "secret",
			},
			isWantErr: false,
			want: want{
				fn: func(token string) error {
					if len(strings.Split(token, ".")) == 3 {
						return nil
					}
					return errors.NewWithCode(codes.CodeJWT, "generated jwt token not valid")
				},
			},
			wantErr: wantErr{},
		},
		{
			ctx:  context.Background(),
			name: "validate jwt token string",
			beforeFunc: func(ctx context.Context, j JWTInterface[*claims], p *params) {
				s, _ := j.Generate(ctx)
				p.tokenString = s
			},
			mode: MODE_VALIDATE,
			params: params{
				secretKey: "secret",
			},
			isWantErr: false,
			want: want{
				fn: func(token string) error {
					return nil
				},
			},
			wantErr: wantErr{},
		},
		{
			ctx:  context.Background(),
			name: "extract claims jwt token string",
			beforeFunc: func(ctx context.Context, j JWTInterface[*claims], p *params) {
				s, _ := j.Generate(ctx)
				p.tokenString = s
			},
			mode: MODE_EXTRACT,
			params: params{
				secretKey: "secret",
			},
			isWantErr: false,
			want: want{
				fn: func(token string) error {
					return nil
				},
			},
			wantErr: wantErr{},
		},
	}

	f := files.GetCurrentMethodName()
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v:%v", f, tt.name), func(t *testing.T) {
			jwtFunc := InitJWT([]byte(tt.params.secretKey), &tt.params.claims)

			if tt.beforeFunc != nil {
				tt.beforeFunc(tt.ctx, jwtFunc, &tt.params)
			}

			switch tt.mode {
			case MODE_GENERATE:
				s, err := jwtFunc.Generate(tt.ctx)
				if tt.isWantErr {
					if err != nil {
						if code := errors.GetCode(err); code != tt.wantErr.code {
							t.Fatalf("want err code is %#v but got err code %#v", tt.wantErr.code, code)
						}
					} else {
						t.Fatalf("want err is %#v but got err %#v", tt.isWantErr, err)
					}
				}

				if err := tt.want.fn(s); err != nil {
					t.Fatalf(err.Error())
				}

				t.Logf("generated token: %#v", s)

			case MODE_VALIDATE:
				_, err := jwtFunc.Validate(tt.ctx, tt.params.tokenString)
				if tt.isWantErr {
					if err != nil {
						if code := errors.GetCode(err); code != tt.wantErr.code {
							t.Fatalf("want err code is %#v but got err code %#v", tt.wantErr.code, code)
						}
					} else {
						t.Fatalf("want err is %#v but got err %#v", tt.isWantErr, err)
					}
				}
			case MODE_EXTRACT:
				jt, err := jwtFunc.Validate(tt.ctx, tt.params.tokenString)
				if err != nil {
					if code := errors.GetCode(err); code != tt.wantErr.code {
						t.Fatalf("want err code is %#v but got err code %#v", tt.wantErr.code, code)
					}
				}

				c, err := jwtFunc.ExtractClaims(tt.ctx, jt)
				if tt.isWantErr {
					if err != nil {
						if code := errors.GetCode(err); code != tt.wantErr.code {
							t.Fatalf("want err code is %#v but got err code %#v", tt.wantErr.code, code)
						}
					} else {
						t.Fatalf("want err is %#v but got err %#v", tt.isWantErr, err)
					}
				}

				t.Logf("claims: %#v", *c)
			}
		})
		fmt.Println("")
	}
}
