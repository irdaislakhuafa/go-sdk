package appcontext

import (
	"context"
	"time"

	"github.com/irdaislakhuafa/go-sdk/language"
	"github.com/irdaislakhuafa/go-sdk/operator"
)

type contextKey string

const (
	AcceptLanguage   contextKey = "AcceptLanguage"
	RequestID        contextKey = "RequestID"
	ServiceVersion   contextKey = "ServiceVersion"
	UserAgent        contextKey = "UserAgent"
	RequestStartTime contextKey = "RequestStartTime"
)

// Set accept language to context
func SetAcceptLanguage(ctx context.Context, lang language.Language) context.Context {
	return context.WithValue(ctx, AcceptLanguage, lang)
}

// Get accept language from context and if not exists, this function will return `language.English` as default value
func GetAcceptLanguage(ctx context.Context) language.Language {
	acceptLanguage, isOk := ctx.Value(AcceptLanguage).(string)
	return operator.Ternary(!isOk, language.English, language.Language(acceptLanguage))
}

// Set request id to context
func SetRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, RequestID, id)
}

// Get request id from context and will return empty string if not exist
func GetRequestID(ctx context.Context) string {
	requestID, isOk := ctx.Value(RequestID).(string)
	return operator.Ternary(!isOk, "", requestID)
}

// Set service version to context
func SetServiceVersion(ctx context.Context, version string) context.Context {
	return context.WithValue(ctx, ServiceVersion, version)
}

// Get service version from context and will return empty string if not exist
func GetServiceVersion(ctx context.Context) string {
	serviceVersion, isOk := ctx.Value(ServiceVersion).(string)
	return operator.Ternary(!isOk, "", serviceVersion)
}

// Set user agent to context
func SetUserAgent(ctx context.Context, ua string) context.Context {
	return context.WithValue(ctx, UserAgent, ua)
}

// Get user agent from context and will return empty string if not exist
func GetUserAgent(ctx context.Context) string {
	userAgent, isOk := ctx.Value(UserAgent).(string)
	return operator.Ternary(!isOk, "", userAgent)
}

// Set request start time to context
func SetRequestStartTime(ctx context.Context, rst time.Time) context.Context {
	return context.WithValue(ctx, RequestStartTime, rst)
}

// Get request start time from context will return zero value of `time.Time` if not exist
func GetRequestStartTime(ctx context.Context) time.Time {
	requestStartTime, isOk := ctx.Value(RequestStartTime).(time.Time)
	return operator.Ternary(!isOk, time.Time{}, requestStartTime)
}
