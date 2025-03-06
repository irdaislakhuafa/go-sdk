package codes

import (
	"github.com/irdaislakhuafa/go-sdk/language"
)

// Alias of unsigned int64 to identify errors by {Code}
type Code uint64

// Default code/no code
const (
	NoCode = Code(0)
)

const (
	// start code from sdk
	CodeStart = Code(iota + 1)

	// General result error codes
	CodeSuccess
	// Other codes

	// Common error codes
	CodeInvalidValue
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
	CodeCommonEnd
	// Other codes

	// SQL error codes
	CodeSQLStart
	CodeSQL
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
	CodeSQLEnd
	// Other codes

	// Client error codes
	CodeClientStart
	CodeClient
	CodeClientMarshal
	CodeClientUnmarshal
	CodeClientErrorOnRequest
	CodeClientErrorOnReadBody
	CodeClientEnd
	// Other codes

	// Auth error codes
	CodeAuthStart
	CodeAuth
	CodeAuthRefreshTokenExpired
	CodeAuthAccessTokenExpired
	CodeAuthFailure
	CodeAuthInvalidToken
	CodeForbidden
	CodeAuthEnd
	// Other codes

	// JSON encoding/decoding error codes
	CodeJSONSchemaStart
	CodeJSONSchema
	CodeJSONSchemaInvalid
	CodeJSONSchemaNotFound
	CodeJSONStructInvalid
	CodeJSONRawInvalid
	CodeJSONValidationError
	CodeJSONMarshalError
	CodeJSONUnmarshalError
	CodeJSONSchemaEnd
	// Other codes

	// Storage error codes
	CodeStorageStart
	CodeStorage
	CodeStorageNoFile
	CodeStorageGenerateURLFailure
	CodeStorageNoClient
	CodeStorageDelFailure
	CodeStorageEnd
	// Other codes

	// JWT error codes
	CodeJWTStart
	CodeJWT
	CodeJWTInvalidMethod
	CodeJWTParseWithClaimsError
	CodeJWTInvalidClaimsType
	CodeJWTSignedStringError
	CodeJWTEnd
	// Other codes

	// GraphQL error codes
	CodeGQLStart
	CodeGQL
	CodeGQLInvalidValue
	CodeGQLBuilder
	CodeGQLEnd
	// Other codes

	// Argon2 error codes
	CodeArgon2Start
	CodeArgon2
	CodeArgon2InvalidEncodedHash
	CodeArgon2EncodeHashError
	CodeArgon2DecodeHashError
	CodeArgon2IncompatibleVersion
	CodeArgon2NotMatch
	CodeArgon2End
	// Other codes

	// Bcrypt error codes
	CodeBcryptStart
	CodeBcrypt
	CodeBcryptEnd
	// Other codes

	// AES 256 GCM error codes
	CodeAES256GCMStart
	CodeAES256GCM
	CodeAES256GCMOpenError
	CodeAES256GCMEnd
	// Other codes

	// SMTP error codes
	CodeSMTPStart
	CodeSMTP
	CodeSMTPBadRequest
	CodeSMTPRequestTimeout
	CodeSMTPEnd
	// Other codes

	// Go Identiface codes (go lib based on https://github.com/Kagami/go-face.git to identify face)
	CodeIdentifaceStart
	CodeIdentiface
	CodeIdentifaceNoFaceDetected
	CodeIdentifaceFaceNotRecognized
	CodeIdentifaceMultipleFaceDetected
	CodeIdentifaceFaceNotRegistered
	CodeIdentifaceEnd
	// Other codes

	// Go string template codes
	CodeStrTemplateStart
	CodeStrTemplateInvalidFormat
	CodeStrTemplateExecuteErr
	CodeStrTemplateEnd
	// Other codes

	// Go Queue codes
	CodeQueueEmpty
	CodeQueueFull
	// Other codes

	// Cache codes
	CodeCacheStart
	CodeCacheKeyNotFound
	CodeCacheEnd
	// Other codes

	// end of sdk code
	CodeEnd
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

// Get code messages in multiple language by `codes.Code` and `language.Language`. Will return empty value if language not available.
func GetCodeMessages(code Code) map[language.Language]Message {
	if msg, isOk := codeMessages[code]; isOk {
		return msg
	}
	return map[language.Language]Message{}
}

// Compile error codes with specific language to available messages from `codes.GetCodeMessages`
func Compile(code Code, lang language.Language) Message {
	msgs := GetCodeMessages(code)
	if msg, isOk := msgs[lang]; isOk {
		return msg
	}

	msg := GetCodeMessages(CodeSuccess)[language.English]
	return msg
}
