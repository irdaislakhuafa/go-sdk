package appcontext

import (
	"context"

	"github.com/irdaislakhuafa/go-sdk/language"
	"github.com/irdaislakhuafa/go-sdk/operator"
)

type contextKey string

const (
	acceptLanguage   contextKey = "AcceptLanguage"
	requestID        contextKey = "RequestID"
	serviceVersion   contextKey = "ServiceVersion"
	userAgent        contextKey = "UserAgent"
	requestStartTime contextKey = "RequestStartTime"
)

// Set accept language to context
func SetAcceptLanguage(ctx context.Context, lang language.Language) context.Context {
	return context.WithValue(ctx, acceptLanguage, lang)
}

// Get accept language from context and if not exists, this function will return `language.English` as default value
func GetAcceptLanguage(ctx context.Context) language.Language {
	acceptLanguage, isOk := ctx.Value(acceptLanguage).(string)
	return operator.Ternary[language.Language](!isOk, language.English, language.Language(acceptLanguage))
}
