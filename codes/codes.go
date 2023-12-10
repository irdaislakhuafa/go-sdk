package codes

import (
	"math"

	"github.com/irdaislakhuafa/go-sdk/language"
)

// Alias of unsigned int64 to identify errors by {Code}
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

// JWT error codes (7001-8000)
const (
	CodeJWT = Code(iota + 7001)
	CodeJWTInvalidMethod
	CodeJWTParseWithClaimsError
	CodeJWTInvalidClaimsType
	CodeJWTSignedStringError
	// Other codes
)

// GraphQL error codes (8001-9000)
const (
	CodeGQL = Code(iota + 8001)
	CodeGQLInvalidValue
	CodeGQLBuilder
	// Other codes
)

// Argon2 error codes (9001-10_000)
const (
	CodeArgon2 = Code(iota + 9001)
	CodeArgon2InvalidEncodedHash
	CodeArgon2EncodeHashError
	CodeArgon2DecodeHashError
	CodeArgon2IncompatibleVersion
	// Other codes
)

// aes 256 gcm error codes (10_001 - 11_000)
const (
	CodeAES256GCM = Code(iota + 10_001)
	CodeAES256GCMOpenError
	// Other codes
)

// smtp error codes (11_001 - 12_000)
const (
	CodeSMTP = Code(iota + 11_001)
	CodeSMTPBadRequest
	CodeSMTPRequestTimeout
	// Other codes
)

var codeMessages = map[Code](map[language.Language]Message){
	// Error messages
	CodeInvalidValue:            getMessages(MsgCodeErrBadRequest),
	CodeContextDeadlineExceeded: getMessages(MsgCodeErrContextTimeout),
	CodeContextCanceled:         getMessages(MsgCodeErrContextTimeout),
	CodeInternalServerError:     getMessages(MsgCodeErrInternalServerError),
	CodeServerUnavailable:       getMessages(MsgCodeErrServiceUnavailable),
	CodeNotImplemented:          getMessages(MsgCodeErrNotImplemented),
	CodeBadRequest:              getMessages(MsgCodeErrBadRequest),
	CodeNotFound:                getMessages(MsgCodeErrNotFound),
	CodeConflict:                getMessages(MsgCodeErrConflict),
	CodeUnauthorized:            getMessages(MsgCodeErrUnauthorized),
	CodeTooManyRequest:          getMessages(MsgCodeErrTooManyRequest),
	CodeMarshal:                 getMessages(MsgCodeErrBadRequest),
	CodeUnmarshal:               getMessages(MsgCodeErrBadRequest),
	CodeJSONMarshalError:        getMessages(MsgCodeErrBadRequest),
	CodeJSONUnmarshalError:      getMessages(MsgCodeErrBadRequest),

	CodeSQL:                   getMessages(MsgCodeErrInternalServerError),
	CodeSQLInit:               getMessages(MsgCodeErrInternalServerError),
	CodeSQLBuilder:            getMessages(MsgCodeErrInternalServerError),
	CodeSQLTxBegin:            getMessages(MsgCodeErrInternalServerError),
	CodeSQLTxCommit:           getMessages(MsgCodeErrInternalServerError),
	CodeSQLTxRollback:         getMessages(MsgCodeErrInternalServerError),
	CodeSQLTxExec:             getMessages(MsgCodeErrInternalServerError),
	CodeSQLPrepareStmt:        getMessages(MsgCodeErrInternalServerError),
	CodeSQLRead:               getMessages(MsgCodeErrInternalServerError),
	CodeSQLRowScan:            getMessages(MsgCodeErrInternalServerError),
	CodeSQLRecordDoesNotExist: getMessages(MsgCodeErrNotFound),
	CodeSQLUniqueConstraint:   getMessages(MsgCodeErrConflict),
	CodeSQLConflict:           getMessages(MsgCodeErrConflict),
	CodeSQLNoRowsAffected:     getMessages(MsgCodeErrNotFound),

	CodeClientMarshal:         getMessages(MsgCodeErrInternalServerError),
	CodeClientUnmarshal:       getMessages(MsgCodeErrInternalServerError),
	CodeClientErrorOnRequest:  getMessages(MsgCodeErrInternalServerError),
	CodeClientErrorOnReadBody: getMessages(MsgCodeErrInternalServerError),

	CodeAuth:                    getMessages(MsgCodeErrUnauthorized),
	CodeAuthRefreshTokenExpired: getMessages(MsgCodeErrRefreshTokenExpired),
	CodeAuthAccessTokenExpired:  getMessages(MsgCodeErrAccessTokenExpired),
	CodeAuthFailure:             getMessages(MsgCodeErrUnauthorized),
	CodeAuthInvalidToken:        getMessages(MsgCodeErrInvalidToken),
	CodeForbidden:               getMessages(MsgCodeErrForbidden),

	// Successfull messages
	CodeSuccess: getMessages(MsgCodeSuccessDefault),
}

func GetCodeMessages(code Code) map[language.Language]Message {
	if msg, isOk := codeMessages[code]; isOk {
		return msg
	}
	return map[language.Language]Message{}
}

func Compile(code Code, lang language.Language) Message {
	msgs := GetCodeMessages(code)
	if msg, isOk := msgs[lang]; isOk {
		return msg
	}

	msg := GetCodeMessages(CodeSuccess)[language.English]
	return msg
}
