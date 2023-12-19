package header

const (
	// Headers keys
	KeyRequestID      = "x-request-id"
	KeyAuthorization  = "authorization"
	KeyUserAgent      = "user-agent"
	KeyContentType    = "content-type"
	KeyContentAccept  = "accept"
	KeyAcceptLanguage = "accept-language"
	KeyCacheControl   = "cache-control"

	// Content type.Specifying the payload in the request
	ContentTypeJSON = "application/json"
	ContentTypeXML  = "application/xml"
	ContentTypeForm = "application/x-www-form-urlencoded"

	// Accepting media. Specifying the types of requested media (in the response)
	// See here: https://en.wikipedia.org/wiki/Content_negotiation
	MediaTextPlain = "text/plain"
	MediaTextHTML  = "text/html"
	MediaTextCSV   = "text/csv"
	MediaTextXML   = "text/xml"

	MediaImageGIF  = "image/gif"
	MediaImageJPEG = "image/jpeg"
	MediaImagePNG  = "image/png"
	MediaImageWEBP = "image/webp"

	// Cache control
	CacheControlNoCache = "no-cache"
	CacheControlNoStore = "no-store"
)
