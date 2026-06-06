package log

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/irdaislakhuafa/go-sdk/appcontext"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

var onceConsole sync.Once = sync.Once{}

type (
	consoleImpl struct {
		log           zerolog.Logger
		funcCtxFields func(ctx context.Context) map[string]any
		isMerge       bool
	}
)

// InitConsole initializes a new console logger with the given configuration.
// It returns an Interface that logs to the console (stdout).
func InitConsole(cfg Config) Interface {
	var zerologger zerolog.Logger
	onceConsole.Do(func() {
		level, err := zerolog.ParseLevel(string(cfg.Level))
		if err != nil {
			zlog.Fatal().Msg(fmt.Sprintf("failed to parse log level from config with err: %v", err))
		}

		zerologger = zerolog.New(os.Stdout).
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

// Debug logs a debug message to the console.
// It takes a context and an object to log, and includes caller information.
func (c *consoleImpl) Debug(ctx context.Context, obj interface{}) {
	c.log.Debug().
		Fields(c.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// Error logs an error message to the console.
// It takes a context and an object to log, and includes caller information.
func (c *consoleImpl) Error(ctx context.Context, obj interface{}) {
	c.log.Error().
		Fields(c.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// Fatal logs a fatal message to the console, then exits the application.
// It takes a context and an object to log, and includes caller information.
func (c *consoleImpl) Fatal(ctx context.Context, obj interface{}) {
	c.log.Fatal().
		Fields(c.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// Info logs an info message to the console.
// It takes a context and an object to log, and includes caller information.
func (c *consoleImpl) Info(ctx context.Context, obj interface{}) {
	c.log.Info().
		Fields(c.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// Trace logs a trace message to the console.
// It takes a context and an object to log, and includes caller information.
func (c *consoleImpl) Trace(ctx context.Context, obj interface{}) {
	c.log.Trace().
		Fields(c.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// Warn logs a warning message to the console.
// It takes a context and an object to log, and includes caller information.
func (c *consoleImpl) Warn(ctx context.Context, obj interface{}) {
	c.log.Warn().
		Fields(c.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// WithCtxFields implements log.Interface.
func (c *consoleImpl) WithCtxFields(funcCtxField func(ctx context.Context) map[string]any) Interface {
	c.funcCtxFields = funcCtxField
	return c
}

func (c *consoleImpl) getContextFields(ctx context.Context) map[string]any {
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

	if c.funcCtxFields != nil {
		custom := c.funcCtxFields(ctx)
		if c.isMerge {
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
func (c *consoleImpl) Cleanup(ctx context.Context) {
	// do nothing because it's show on console
}

// MergedFields implements Interface.
func (c *consoleImpl) MergedFields(ctx context.Context) Interface {
	c.isMerge = true
	return c
}
