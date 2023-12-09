package codes

import (
	"net/http"

	"github.com/irdaislakhuafa/go-sdk/language"
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
		http.StatusOK: {
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
		http.StatusBadRequest: {
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

		// HTTP Status 5xx
	}
)

func getMessages(httpStatusCode int) map[language.Language]Message {
	if messages == nil {
		return map[language.Language]Message{}
	}
	return messages[httpStatusCode]
}
