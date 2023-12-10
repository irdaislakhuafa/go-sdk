package errors

import (
	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/operator"
)

type stacktrace struct {
	message  string
	cause    error
	code     codes.Code
	file     string
	function string
	line     int
}

func (s *stacktrace) Error() string {
	return s.message
}

func (s *stacktrace) ExitCode() int {
	return operator.Ternary[int](s.code == codes.NoCode, 1, int(s.code))
}
