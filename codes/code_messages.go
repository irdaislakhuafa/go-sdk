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
			language.English:    Message{},
			language.Indonesian: Message{},
		},

		// HTTP Status 5xx
	}
)

func getMessages(msgCode int) map[language.Language]Message {
	if messages == nil {
		return map[language.Language]Message{}
	}
	return messages[msgCode]
}
