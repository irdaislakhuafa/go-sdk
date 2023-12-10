package appcontext

import (
	"context"
	"time"

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

// Set request id to context
func SetRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestID, id)
}

// Get request id from context and will return empty string if not exist
func GetRequestID(ctx context.Context) string {
	requestID, isOk := ctx.Value(requestID).(string)
	return operator.Ternary[string](!isOk, "", requestID)
}

// Set service version to context
func SetServiceVersion(ctx context.Context, version string) context.Context {
	return context.WithValue(ctx, serviceVersion, version)
}

// Get service version from context and will return empty string if not exist
func GetServiceVersion(ctx context.Context) string {
	serviceVersion, isOk := ctx.Value(serviceVersion).(string)
	return operator.Ternary[string](!isOk, "", serviceVersion)
}

// Set user agent to context
func SetUserAgent(ctx context.Context, ua string) context.Context {
	return context.WithValue(ctx, userAgent, ua)
}

// Get user agent from context and will return empty string if not exist
func GetUserAgent(ctx context.Context) string {
	userAgent, isOk := ctx.Value(userAgent).(string)
	return operator.Ternary[string](!isOk, "", userAgent)
}

// Set request start time to context
func SetRequestStartTime(ctx context.Context, rst time.Time) context.Context {
	return context.WithValue(ctx, requestStartTime, rst)
}

// Get request start time from context will return zero value of `time.Time` if not exist
func GetRequestStartTime(ctx context.Context) time.Time {
	requestStartTime, isOk := ctx.Value(requestStartTime).(time.Time)
	return operator.Ternary[time.Time](!isOk, time.Time{}, requestStartTime)
}
