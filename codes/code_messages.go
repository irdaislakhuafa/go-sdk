package codes

import (
	"net/http"

	"github.com/irdaislakhuafa/go-sdk/language"
)

// code of messages
const (
	// 2xx (default)
	MsgCodeSuccessDefault = (iota + 1)

	// 4xx
	MsgCodeErrBadRequest
	MsgCodeErrUnauthorized
	MsgCodeErrInvalidToken
	MsgCodeErrRefreshTokenExpired
	MsgCodeErrAccessTokenExpired
	MsgCodeErrForbidden
	MsgCodeErrNotFound
	MsgCodeErrContextTimeout
	MsgCodeErrConflict
	MsgCodeErrTooManyRequest

	// 5xx
	MsgCodeErrInternalServerError
	MsgCodeErrNotImplemented
	MsgCodeErrServiceUnavailable
)

// Struct to store error message
type Message struct {
	StatusCode int
	Title      string
	Body       string
}

var (
	// Collections of messages in multiple language
	messages = map[int](map[language.Language]Message){
		// HTTP Status 1xx
		MsgCodeSuccessDefault: {
			language.English: Message{
				StatusCode: http.StatusOK,
				Title:      language.HTTPStatusText(language.English, http.StatusOK),
				Body:       "Request successful",
			},
			language.Indonesian: Message{
				StatusCode: http.StatusOK,
				Title:      language.HTTPStatusText(language.Indonesian, http.StatusOK),
				Body:       "Request berhasil",
			},
		},

		// HTTP Status 2xx
		// HTTP Status 3xx

		// HTTP Status 4xx
		MsgCodeErrBadRequest: {
			language.English: Message{
				StatusCode: http.StatusBadRequest,
				Title:      language.HTTPStatusText(language.English, http.StatusBadRequest),
				Body:       "Invalid input. Please validate your input.",
			},
			language.Indonesian: Message{
				StatusCode: http.StatusBadRequest,
				Title:      language.HTTPStatusText(language.Indonesian, http.StatusBadRequest),
				Body:       "Masukan data tidak valid. Mohon cek kembali masukan anda.",
			},
		},
		MsgCodeErrUnauthorized: {
			language.English: Message{
				StatusCode: http.StatusUnauthorized,
				Title:      language.HTTPStatusText(language.English, http.StatusUnauthorized),
				Body:       "Unauthorized access. You are not authorized to access this resource.",
			},
			language.Indonesian: Message{
				StatusCode: http.StatusUnauthorized,
				Title:      language.HTTPStatusText(language.Indonesian, http.StatusUnauthorized),
				Body:       "Akses ditolak. Anda tidak memilik izin untuk mengakses sumber daya ini.",
			},
		},
		MsgCodeErrInvalidToken: {
			language.English: Message{
				StatusCode: http.StatusUnauthorized,
				Title:      language.HTTPStatusText(language.English, http.StatusUnauthorized),
				Body:       "Invalid token. Please renew your session by reloading.",
			},
			language.Indonesian: Message{
				StatusCode: http.StatusUnauthorized,
				Title:      language.HTTPStatusText(language.Indonesian, http.StatusUnauthorized),
				Body:       "Token tidak valid. Mohon perbarui sesi anda dengan menggakses ulang laman.",
			},
		},
		MsgCodeErrRefreshTokenExpired: {
			language.English: Message{
				StatusCode: http.StatusUnauthorized,
				Title:      language.HTTPStatusText(language.English, http.StatusUnauthorized),
				Body:       "Session refresh token has expired. Please renew your session by reloading.",
			},
			language.Indonesian: Message{
				StatusCode: http.StatusUnauthorized,
				Title:      language.HTTPStatusText(language.Indonesian, http.StatusUnauthorized),
				Body:       "Token pembaruan sudah tidak berlaku. Mohon perbarui sesi anda dengan mengakses ulang laman.",
			},
		},
		MsgCodeErrAccessTokenExpired: {
			language.English:    Message{},
			language.Indonesian: Message{},
		},
		MsgCodeErrForbidden: {
			language.English:    Message{},
			language.Indonesian: Message{},
		},
		MsgCodeErrNotFound: {
			language.English:    Message{},
			language.Indonesian: Message{},
		},
		MsgCodeErrContextTimeout: {
			language.English:    Message{},
			language.Indonesian: Message{},
		},
		MsgCodeErrConflict: {
			language.English:    Message{},
			language.Indonesian: Message{},
		},
		MsgCodeErrTooManyRequest: {
			language.English:    Message{},
			language.Indonesian: Message{},
		},

		// HTTP Status 5xx
		MsgCodeErrInternalServerError: {
			language.English:    Message{},
			language.Indonesian: Message{},
		},
		MsgCodeErrNotImplemented: {
			language.English:    Message{},
			language.Indonesian: Message{},
		},
		MsgCodeErrServiceUnavailable: {
			language.English:    Message{},
			language.Indonesian: Message{},
		},
	}
)

func getMessages(msgCode int) map[language.Language]Message {
	if messages == nil {
		return map[language.Language]Message{}
	}
	return messages[msgCode]
}
