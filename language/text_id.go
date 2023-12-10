package language

import "net/http"

// Surprisingly, we have a kinda official translation for HTTP status in Indonesian.
// Check it here: https://id.wikipedia.org/wiki/Daftar_kode_status_HTTP
var statusTextID = map[int]string{
	// HTTP Status 1xx
	http.StatusContinue:           "Lanjutkan",
	http.StatusSwitchingProtocols: "Beralih Protokol",
	http.StatusProcessing:         "Processing",
	http.StatusEarlyHints:         "Petunjuk Awal",

	// HTTP Status 2xx
	http.StatusOK:                   "OK",
	http.StatusCreated:              "Dibuat",
	http.StatusAccepted:             "Diterima",
	http.StatusNonAuthoritativeInfo: "Informasi Non-Resmi",
	http.StatusNoContent:            "Tanpa Konten",
	http.StatusResetContent:         "Setel Ulang Konten",
	http.StatusPartialContent:       "Konten Sebagian",
	http.StatusMultiStatus:          "Multi-Status",
	http.StatusAlreadyReported:      "Sudah Dilaporkan",
	http.StatusIMUsed:               "IM Used",

	// HTTP Status 3xx
	http.StatusMultipleChoices:   "Pilihan Ganda",
	http.StatusMovedPermanently:  "Dipindahkan Permanen",
	http.StatusFound:             "Ditemukan",
	http.StatusSeeOther:          "Lihat Lainnya",
	http.StatusNotModified:       "Tidak Dimodifikasi",
	http.StatusUseProxy:          "Gunakan Proxy",
	http.StatusTemporaryRedirect: "Pengalihan Sementara",
	http.StatusPermanentRedirect: "Pengalihan Permanen",

	// HTTP Status 4xx
	http.StatusBadRequest:                   "Bad Request",
	http.StatusUnauthorized:                 "Tidak Diperbolehkan",
	http.StatusPaymentRequired:              "Payment Required",
	http.StatusForbidden:                    "Terlarang",
	http.StatusNotFound:                     "Tidak Ditemukan",
	http.StatusMethodNotAllowed:             "Metode Tidak Diizinkan",
	http.StatusNotAcceptable:                "Tidak Dapat Diterima",
	http.StatusProxyAuthRequired:            "Diperlukan Otentikasi Proksi",
	http.StatusRequestTimeout:               "Request Timeout",
	http.StatusConflict:                     "Konflik",
	http.StatusGone:                         "Hilang",
	http.StatusLengthRequired:               "Panjang Diperlukan",
	http.StatusPreconditionFailed:           "Precondition Failed",
	http.StatusRequestEntityTooLarge:        "Payload Terlalu Besar",
	http.StatusRequestURITooLong:            "URI Terlalu Panjang",
	http.StatusUnsupportedMediaType:         "Jenis Media Yang Tidak Didukung",
	http.StatusRequestedRangeNotSatisfiable: "Requested Range Not Statisfiable",
	http.StatusExpectationFailed:            "Expectation Failed",
	http.StatusTeapot:                       "I'm a Teapot",
	http.StatusMisdirectedRequest:           "Misdirected Request",
	http.StatusUnprocessableEntity:          "Unprocessable Entity",
	http.StatusLocked:                       "Locked",
	http.StatusFailedDependency:             "Failed Dependency",
	http.StatusTooEarly:                     "Too Early",
	http.StatusUpgradeRequired:              "Diperlukan Peningkatan",
	http.StatusPreconditionRequired:         "Precondition Required",
	http.StatusTooManyRequests:              "Too Many Requests",
	http.StatusRequestHeaderFieldsTooLarge:  "Request Header Fields Too Large",
	http.StatusUnavailableForLegalReasons:   "Unavailable For Legal Reasons",

	// HTTP Status 5xx
	http.StatusInternalServerError:           "Masalah Peladen Dalam",
	http.StatusNotImplemented:                "'Not Implemented",
	http.StatusBadGateway:                    "Bad Gateway",
	http.StatusServiceUnavailable:            "Service Unavailable",
	http.StatusGatewayTimeout:                "Gateway Timeout",
	http.StatusHTTPVersionNotSupported:       "HTTP Version Not Supported",
	http.StatusVariantAlsoNegotiates:         "Variant Also Negotiates",
	http.StatusInsufficientStorage:           "Insufficient Storage",
	http.StatusLoopDetected:                  "Loop Detected",
	http.StatusNotExtended:                   "Not Extended",
	http.StatusNetworkAuthenticationRequired: "Network Authentication Required",
}
