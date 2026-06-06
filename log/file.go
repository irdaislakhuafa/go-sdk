package log

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/irdaislakhuafa/go-sdk/appcontext"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

var onceFile sync.Once = sync.Once{}

type (
	fileImpl struct {
		log           zerolog.Logger
		funcCtxFields func(ctx context.Context) map[string]any
		cleanup       func(ctx context.Context)
		isMerge       bool
	}
)

// InitFile initializes a new file logger with the given configuration.
// It returns an Interface that logs to a file using lumberjack for rotation.
func InitFile(cfg Config) Interface {
	var zerologger zerolog.Logger
	var fileLog io.Writer
	var cleanup func(ctx context.Context) = func(ctx context.Context) {}

	onceFile.Do(func() {
		level, err := zerolog.ParseLevel(string(cfg.Level))
		if err != nil {
			zlog.Fatal().Msg(fmt.Sprintf("failed to parse log level from config with err: %v", err))
		}

		if cfg.Storage.Rotation.Enable {
			file := &lumberjack.Logger{
				Filename:   cfg.Storage.FileLocation,        // Log file path
				MaxSize:    cfg.Storage.Rotation.MaxSize,    // Megabytes before rolling
				MaxBackups: cfg.Storage.Rotation.MaxBackups, // Maximum number of old log files to retain
				MaxAge:     cfg.Storage.Rotation.MaxAge,     // Maximum number of days to retain old log files
				Compress:   cfg.Storage.Rotation.Compress,   // Compress old log files (gzip)
			}
			fileLog = file
			cleanup = func(ctx context.Context) {
				file.Close()
			}
		} else {
			dir := filepath.Dir(cfg.Storage.FileLocation)
			if err := os.MkdirAll(dir, 0755); err != nil {
				zlog.Fatal().Msg(fmt.Sprintf("failed to make dir '%v', %v", dir, err))
			}

			file, err := os.OpenFile(cfg.Storage.FileLocation, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				zlog.Fatal().Msg(err.Error())
				return
			}
			fileLog = file
			cleanup = func(ctx context.Context) {
				file.Close()
			}
		}

		zerologger = zerolog.New(fileLog).
			With().
			Timestamp().
			CallerWithSkipFrameCount(cfg.SkipFrameCount).
			Logger().
			Level(level)
	})

	result := &fileImpl{
		log:           zerologger,
		funcCtxFields: nil,
		cleanup:       cleanup,
	}

	return result
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

	base := map[string]any{
		"request_id":      appcontext.GetRequestID(ctx),
		"user_agent":      appcontext.GetUserAgent(ctx),
		"service_version": appcontext.GetServiceVersion(ctx),
		"time_elapsed":    timeElapsed,
	}

	if f.funcCtxFields != nil {
		custom := f.funcCtxFields(ctx)
		if f.isMerge {
			for k, v := range custom {
				base[k] = v
			}
			return base
		}

		return custom
	}

	return base
}

// Cleanup implements Interface.
func (f *fileImpl) Cleanup(ctx context.Context) {
	f.cleanup(ctx)
}

// MergedFields implements Interface.
func (f *fileImpl) MergedFields(ctx context.Context) Interface {
	f.isMerge = true
	return f
}
