package log

import (
	"context"
	"fmt"
	"time"

	"github.com/irdaislakhuafa/go-sdk/appcontext"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

type (
	fileImpl struct {
		log           zerolog.Logger
		funcCtxFields func(ctx context.Context) map[string]any
	}
)

// InitFile initializes a new file logger with the given configuration.
// It returns an Interface that logs to a file using lumberjack for rotation.
func InitFile(cfg Config) Interface {
	var zerologger zerolog.Logger
	once.Do(func() {
		level, err := zerolog.ParseLevel(string(cfg.Level))
		if err != nil {
			zlog.Fatal().Msg(fmt.Sprintf("failed to parse log level from config with err: %v", err))
		}

		fileLog := lumberjack.Logger{
			Filename:   cfg.Storage.FileLocation, // Log file path
			MaxSize:    30,                       // Megabytes before rolling
			MaxBackups: 3,                        // Maximum number of old log files to retain
			MaxAge:     30,                       // Maximum number of days to retain old log files
			Compress:   true,                     // Compress old log files (gzip)
		}

		zerologger = zerolog.New(&fileLog).
			With().
			Timestamp().
			CallerWithSkipFrameCount(cfg.SkipFrameCount).
			Logger().
			Level(level)
	})

	return &fileImpl{
		log:           zerologger,
		funcCtxFields: nil,
	}
}

// Debug logs a debug message to a file.
// It takes a context and an object to log, and includes caller information.
func (f *fileImpl) Debug(ctx context.Context, obj interface{}) {
	f.log.Debug().
		Fields(f.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// Error logs an error message to a file.
// It takes a context and an object to log, and includes caller information.
func (f *fileImpl) Error(ctx context.Context, obj interface{}) {
	f.log.Error().
		Fields(f.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// Fatal logs a fatal message to a file, then exits the application.
// It takes a context and an object to log, and includes caller information.
func (f *fileImpl) Fatal(ctx context.Context, obj interface{}) {
	f.log.Fatal().
		Fields(f.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// Info logs an info message to a file.
// It takes a context and an object to log, and includes caller information.
func (f *fileImpl) Info(ctx context.Context, obj interface{}) {
	f.log.Info().
		Fields(f.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// Trace logs a trace message to a file.
// It takes a context and an object to log, and includes caller information.
func (f *fileImpl) Trace(ctx context.Context, obj interface{}) {
	f.log.Trace().
		Fields(f.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// Warn logs a warning message to a file.
// It takes a context and an object to log, and includes caller information.
func (f *fileImpl) Warn(ctx context.Context, obj interface{}) {
	f.log.Warn().
		Fields(f.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// WithCtxFields implements log.Interface.
func (f *fileImpl) WithCtxFields(funcCtxField func(ctx context.Context) map[string]any) Interface {
	f.funcCtxFields = funcCtxField
	return f
}

// getContextFields returns a map of fields extracted from the context for logging.
// It includes request_id, user_agent, service_version, and time_elapsed.
func (f *fileImpl) getContextFields(ctx context.Context) map[string]any {
	reqStartTime := appcontext.GetRequestStartTime(ctx)
	timeElapsed := "0ms"

	if !reqStartTime.IsZero() {
		timeElapsed = fmt.Sprintf("%dms", uint64(time.Since(reqStartTime)/time.Millisecond))
	}

	if f.funcCtxFields != nil {
		return f.funcCtxFields(ctx)
	}

	return map[string]any{
		"request_id":      appcontext.GetRequestID(ctx),
		"user_agent":      appcontext.GetUserAgent(ctx),
		"service_version": appcontext.GetServiceVersion(ctx),
		"time_elapsed":    timeElapsed,
	}
}
