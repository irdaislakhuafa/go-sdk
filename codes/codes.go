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
	// Other codes
)

// SQL error codes (2001-3000)
const (
	CodeSQL = Code(iota + 2001)
	CodeSQLInit
	CodeSQLBuilder
	CodeSQLTxBegin
	CodeSQLTxCommit
	CodeSQLTxRollback
	CodeSQLTxExec
	CodeSQLPrepareStmt
	CodeSQLRead
	CodeSQLRowScan
	CodeSQLRecordDoesNotExist
	CodeSQLUniqueConstraint
	CodeSQLConflict
	CodeSQLNoRowsAffected
	// Other codes
)

// Client error codes (3001-4000)
const (
	CodeClient = Code(iota + 3001)
	CodeClientMarshal
	CodeClientUnmarshal
	CodeClientErrorOnRequest
	CodeClientErrorOnReadBody
	// Other codes
)

// Auth error codes (4001-5000)
const (
	CodeAuth = Code(iota + 4001)
	CodeAuthRefreshTokenExpired
	CodeAuthAccessTokenExpired
	CodeAuthFailure
	CodeAuthInvalidToken
	CodeForbidden
	// Other codes
)

// JSON encoding/decoding error codes (5001-6000)
const (
	CodeJSONSchema = Code(iota + 5001)
	CodeJSONSchemaInvalid
	CodeJSONSchemaNotFound
	CodeJSONStructInvalid
	CodeJSONRawInvalid
	CodeJSONValidationError
	CodeJSONMarshalError
	CodeJSONUnmarshalError
	// Other codes
)

// Storage error codes (6001-7000)
const (
	CodeStorage = Code(iota + 6001)
	CodeStorageNoFile
	CodeStorageGenerateURLFailure
	CodeStorageNoClient
	// Other codes
)
