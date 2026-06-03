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
		log zerolog.Logger
		funcCtxFields func(ctx context.Context) map[string]any
	}
)

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

	return &consoleImpl{
		log:           zerologger,
		funcCtxFields: nil,
	}
}

// Debug implements log.Interface.
func (f *fileImpl) Debug(ctx context.Context, obj interface{}) {
	f.log.Debug().
		Fields(f.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// Error implements log.Interface.
func (f *fileImpl) Error(ctx context.Context, obj interface{}) {
	f.log.Error().
		Fields(f.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// Fatal implements log.Interface.
func (f *fileImpl) Fatal(ctx context.Context, obj interface{}) {
	f.log.Fatal().
		Fields(f.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// Info implements log.Interface.
func (f *fileImpl) Info(ctx context.Context, obj interface{}) {
	f.log.Info().
		Fields(f.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// Trace implements log.Interface.
func (f *fileImpl) Trace(ctx context.Context, obj interface{}) {
	f.log.Trace().
		Fields(f.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// Warn implements log.Interface.
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

func (c *fileImpl) getContextFields(ctx context.Context) map[string]any {
	reqStartTime := appcontext.GetRequestStartTime(ctx)
	timeElapsed := "0ms"

	if !reqStartTime.IsZero() {
		timeElapsed = fmt.Sprintf("%dms", uint64(time.Since(reqStartTime)/time.Millisecond))
	}

	if c.funcCtxFields != nil {
		return c.funcCtxFields(ctx)
	}

	return map[string]any{
		"request_id":      appcontext.GetRequestID(ctx),
		"user_agent":      appcontext.GetUserAgent(ctx),
		"service_version": appcontext.GetServiceVersion(ctx),
		"time_elapsed":    timeElapsed,
	}
}
