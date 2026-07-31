package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
)

func ConvertToCurl(ctx context.Context, req http.Request) (string, error) {
	curl := strings.Builder{}
	curl.WriteString("curl \\\n")

	// convert method
	curl.WriteString(fmt.Sprintf("--request '%s' \\\n", req.Method))

	// convert url
	curl.WriteString(fmt.Sprintf("--url '%v' \\\n", req.URL.String()))

	// convert headers
	for k, v := range req.Header.Clone() {
		curl.WriteString(fmt.Sprintf("--header '%s: %s' \\\n", k, strings.Join(v, ",")))
	}

	// convert body
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return "", errors.NewWithCode(codes.CodeClientErrorOnReadBody, "%s", err.Error())
	}

	curl.WriteString(fmt.Sprintf("--data '%s' ", string(bodyBytes)))

	return curl.String(), nil
}
