package codes

import "math"

type Code uint64

// Default code/no code
const (
	NoCode = math.MaxUint64
)

// General result error codes (1-1000)
const (
	CodeSuccess = Code(iota + 1)
	// Other codes
)

// Common error codes (1001-2000)
const (
	CodeInvalidValue = Code(iota + 1001)
	CodeContextDeadlineExceeded
	CodeContextCanceled
	CodeInternalServerError
	CodeServerUnavailable
	CodeNotImplemented
	CodeBadRequest
	CodeNotFound
	CodeConflict
	CodeUnauthorized
	CodeTooManyRequest
	CodeMarshal
	CodeUnmarshal
)
