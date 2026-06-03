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

var once sync.Once = sync.Once{}

type (
	consoleImpl struct {
		log           zerolog.Logger
		funcCtxFields func(ctx context.Context) map[string]any
	}
)

func InitConsole(cfg Config) Interface {
	var zerologger zerolog.Logger
	once.Do(func() {
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

// Debug implements log.Interface.
func (c *consoleImpl) Debug(ctx context.Context, obj interface{}) {
	c.log.Debug().
		Fields(c.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// Error implements log.Interface.
func (c *consoleImpl) Error(ctx context.Context, obj interface{}) {
	c.log.Error().
		Fields(c.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// Fatal implements log.Interface.
func (c *consoleImpl) Fatal(ctx context.Context, obj interface{}) {
	c.log.Fatal().
		Fields(c.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// Info implements log.Interface.
func (c *consoleImpl) Info(ctx context.Context, obj interface{}) {
	c.log.Info().
		Fields(c.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// Trace implements log.Interface.
func (c *consoleImpl) Trace(ctx context.Context, obj interface{}) {
	c.log.Trace().
		Fields(c.getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

// Warn implements log.Interface.
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
