package random

import (
	"fmt"
	"strconv"
	"time"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/irdaislakhuafa/go-sdk/operator"
)

// Generate random number based on time.Now().UnixNano() with return string of random int.
func SInt(maxLength int) (string, error) {
	s := fmt.Sprint(time.Now().UnixNano())
	for maxLength > len(s) {
		s += s
	}

	s = s[len(s)-maxLength:]
	return s, nil
}

// Generate random number based on time.Now().UnixNano() with return random int with current maximum character is 19.
//
// Will return error if generate is failure.
//
// errors:
//
// - `codes.CodeInvalidValue`.
func Int(maxLength int) (int, error) {
	s, err := SInt(maxLength)
	if err != nil {
		return 0, errors.NewWithCode(errors.GetCode(err), "%s", err.Error())
	}
	s = operator.Ternary(s[:1] == "0", "1"+s[1:], s)

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, errors.NewWithCode(codes.CodeInvalidValue, "%s", err.Error())
	}

	return int(i), nil
}
