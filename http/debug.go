package http

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
)

// escapeShellArg safely escapes a string for POSIX shell single-quoted arguments.
func escapeShellArg(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "'\\''") + "'"
}

func ConvertToCurl(ctx context.Context, req http.Request) (string, error) {
	curl := strings.Builder{}
	curl.WriteString("curl \\\n")

	// convert method
	curl.WriteString(fmt.Sprintf("--request %s \\\n", escapeShellArg(req.Method)))

	// convert url
	if req.URL != nil {
		curl.WriteString(fmt.Sprintf("--url %s \\\n", escapeShellArg(req.URL.String())))
	}

	// convert headers
	for k, v := range req.Header.Clone() {
		headerVal := strings.Join(v, ",")
		curl.WriteString(fmt.Sprintf("--header %s \\\n", escapeShellArg(fmt.Sprintf("%s: %s", k, headerVal))))
	}

	// convert body
	if req.Body != nil {
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			return "", errors.NewWithCode(codes.CodeClientErrorOnReadBody, "%s", err.Error())
		}
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))

		curl.WriteString(fmt.Sprintf("--data %s ", escapeShellArg(string(bodyBytes))))
	} else {
		curl.WriteString("--data '' ")
	}

	return curl.String(), nil
}
