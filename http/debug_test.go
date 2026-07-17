package http

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/irdaislakhuafa/go-sdk/header"
)

type errReader struct{}

func (errReader) Read([]byte) (int, error) {
	return 0, io.ErrClosedPipe
}

func TestConvertToCurl(t *testing.T) {
	type (
		args struct {
			req http.Request
		}

		want struct {
			wantCode codes.Code
			contains []string
		}

		test struct {
			name     string
			args     args
			want     want
			fnBefore func(test *test)
			fnAfter  func(test *test)
		}
	)

	url := &url.URL{Scheme: "http", Host: "localhost"}
	tests := []test{
		{
			name: "GET with no body and no headers success",
			args: args{
				req: http.Request{
					Method: http.MethodGet,
					URL:    url,
					Header: http.Header{},
					Body:   io.NopCloser(strings.NewReader("")),
				},
			},
			want: want{
				wantCode: codes.NoCode,
				contains: []string{
					"--method 'GET'",
					"--data ''",
				},
			},
			fnBefore: func(test *test) {},
			fnAfter:  func(test *test) {},
		},
		{
			name: "POST with header and body success",
			args: args{
				req: http.Request{
					Method: http.MethodPost,
					URL:    url,
					Header: http.Header{
						header.KeyContentType: {header.ContentTypeJSON},
					},
					Body: io.NopCloser(strings.NewReader(`{"a":1}`)),
				},
			},
			want: want{
				wantCode: codes.NoCode,
				contains: []string{
					"--method 'POST'",
					"--header '" + header.KeyContentType + ": " + header.ContentTypeJSON + "'",
					"--data '{\"a\":1}'",
				},
			},
			fnBefore: func(test *test) {},
			fnAfter:  func(test *test) {},
		},
		{
			name: "error on read body",
			args: args{
				req: http.Request{
					Method: http.MethodPost,
					URL:    url,
					Header: http.Header{},
					Body:   io.NopCloser(errReader{}),
				},
			},
			want: want{
				wantCode: codes.CodeClientErrorOnReadBody,
			},
			fnBefore: func(test *test) {},
			fnAfter:  func(test *test) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fnBefore != nil {
				tt.fnBefore(&tt)
			}

			curl, err := ConvertToCurl(context.Background(), tt.args.req)
			code := errors.GetCode(err)

			if tt.want.wantCode != code {
				t.Fatalf("want error code '%v' but got '%v'", tt.want.wantCode, code)
			}

			for _, c := range tt.want.contains {
				if !strings.Contains(curl, c) {
					t.Fatalf("want curl contains '%v' but got '%v'", c, curl)
				}
			}

			if tt.fnAfter != nil {
				tt.fnAfter(&tt)
			}
		})
	}
}
