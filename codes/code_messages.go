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
			language.English: Message{
				StatusCode: http.StatusUnauthorized,
				Title:      language.HTTPStatusText(language.English, http.StatusUnauthorized),
				Body:       "Session access token has expired. Please renew your session by reloading.",
			},
			language.Indonesian: Message{
				StatusCode: http.StatusUnauthorized,
				Title:      language.HTTPStatusText(language.Indonesian, http.StatusUnauthorized),
				Body:       "Token akses sudah tidak berlaku. Mohon perbarui sesi anda dengan mengakses ulang laman.",
			},
		},
		MsgCodeErrForbidden: {
			language.English: Message{
				StatusCode: http.StatusForbidden,
				Title:      language.HTTPStatusText(language.English, http.StatusForbidden),
				Body:       "Forbidden. You don't have permission to access this resource.",
			},
			language.Indonesian: Message{
				StatusCode: http.StatusForbidden,
				Title:      language.HTTPStatusText(language.Indonesian, http.StatusForbidden),
				Body:       "Terlarang. Anda tidak memiliki izin untuk mengakses laman ini",
			},
		},
		MsgCodeErrNotFound: {
			language.English: Message{
				StatusCode: http.StatusNotFound,
				Title:      language.HTTPStatusText(language.English, http.StatusNotFound),
				Body:       "Record doesn't exist. Please validate your input or contact the administrator.",
			},
			language.Indonesian: Message{
				StatusCode: http.StatusNotFound,
				Title:      language.HTTPStatusText(language.Indonesian, http.StatusNotFound),
				Body:       "Data tidak ditemukan. Mohon cek kembali masukan anda atau hubungi administrator.",
			},
		},
		MsgCodeErrContextTimeout: {
			language.English: Message{
				StatusCode: http.StatusRequestTimeout,
				Title:      language.HTTPStatusText(language.English, http.StatusRequestTimeout),
				Body:       "Request time has been exceeded.",
			},
			language.Indonesian: Message{
				StatusCode: http.StatusRequestTimeout,
				Title:      language.HTTPStatusText(language.Indonesian, http.StatusRequestTimeout),
				Body:       "Waktu permintaan habis.",
			},
		},
		MsgCodeErrConflict: {
			language.English: Message{
				StatusCode: http.StatusConflict,
				Title:      language.HTTPStatusText(language.English, http.StatusConflict),
				Body:       "Record has existed. Please validate your input or contact the administrator.",
			},
			language.Indonesian: Message{
				StatusCode: http.StatusConflict,
				Title:      language.HTTPStatusText(language.Indonesian, http.StatusConflict),
				Body:       "Data sudah ada. Mohon cek kembali masukan anda atau hubungi administrator.",
			},
		},
		MsgCodeErrTooManyRequest: {
			language.English: Message{
				StatusCode: http.StatusTooManyRequests,
				Title:      language.HTTPStatusText(language.English, http.StatusTooManyRequests),
				Body:       "Too many requests. Please wait and try again after a few seconds.",
			},
			language.Indonesian: Message{
				StatusCode: http.StatusTooManyRequests,
				Title:      language.HTTPStatusText(language.Indonesian, http.StatusTooManyRequests),
				Body:       "Terlalu banyak permintaan. Mohon tunggu dan coba lagi setelah beberapa saat.",
			},
		},

		// HTTP Status 5xx
		MsgCodeErrInternalServerError: {
			language.English: Message{
				StatusCode: http.StatusInternalServerError,
				Title:      language.HTTPStatusText(language.English, http.StatusInternalServerError),
				Body:       "Internal server error. Please contact the administrator.",
			},
			language.Indonesian: Message{
				StatusCode: http.StatusInternalServerError,
				Title:      language.HTTPStatusText(language.Indonesian, http.StatusInternalServerError),
				Body:       "Terjadi kendala diserver. Mohon hubungi administrator.",
			},
		},
		MsgCodeErrNotImplemented: {
			language.English: Message{
				StatusCode: http.StatusNotImplemented,
				Title:      language.HTTPStatusText(language.English, http.StatusNotImplemented),
				Body:       "Not Implemented. Please contact the administrator.",
			},
			language.Indonesian: Message{
				StatusCode: http.StatusNotImplemented,
				Title:      language.HTTPStatusText(language.Indonesian, http.StatusNotImplemented),
				Body:       "Layanan tidak tersedia. Mohon hubungi administrator.",
			},
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
