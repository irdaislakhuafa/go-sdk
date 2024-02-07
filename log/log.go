package log

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/irdaislakhuafa/go-sdk/appcontext"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var once sync.Once = sync.Once{}

type Interface interface {
	// TODO: added method Debugf
	Trace(ctx context.Context, obj interface{})
	Debug(ctx context.Context, obj interface{})
	Info(ctx context.Context, obj interface{})
	Warn(ctx context.Context, obj interface{})
	Error(ctx context.Context, obj interface{})
	Fatal(ctx context.Context, obj interface{})
}

type Config struct {
	Level string
}

type logger struct {
	log zerolog.Logger
}

const (
	skipFrameCount = 3 // NOTE: temporary 3 for now
)

func Init(cfg Config) Interface {
	var zerologger zerolog.Logger
	once.Do(func() {
		level, err := zerolog.ParseLevel(cfg.Level)
		if err != nil {
			log.Fatal().Msg(fmt.Sprintf("failed to parse log level from config with err: %v", err))
		}

		zerologger = zerolog.New(os.Stdout).
			With().
			Timestamp().
			CallerWithSkipFrameCount(skipFrameCount).
			Logger().
			Level(level)
	})

	return &logger{log: zerologger}
}

func (l *logger) Trace(ctx context.Context, obj interface{}) {
	l.log.Trace().
		Fields(getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

func (l *logger) Debug(ctx context.Context, obj interface{}) {
	l.log.Debug().
		Fields(getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

func (l *logger) Info(ctx context.Context, obj interface{}) {
	l.log.Info().
		Fields(getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

func (l *logger) Warn(ctx context.Context, obj interface{}) {
	l.log.Warn().
		Fields(getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

func (l *logger) Error(ctx context.Context, obj interface{}) {
	l.log.Error().
		Fields(getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

func (l *logger) Fatal(ctx context.Context, obj interface{}) {
	l.log.Fatal().
		Fields(getContextFields(ctx)).
		Msg(fmt.Sprint(GetCaller(obj)))
}

func GetCaller(value any) any {
	switch tErr := value.(type) {
	case error:
		file, line, message, err := errors.GetCaller(tErr)
		if err != nil {
			return fmt.Sprintf("error cannot get caller, %v", err)
		}
		return fmt.Sprintf("%s:%#v --- %s", file, line, message)
	case string:
		return tErr
	default:
		return fmt.Sprintf("%#v", tErr)
	}
}

func getContextFields(ctx context.Context) map[string]any {
	reqStartTime := appcontext.GetRequestStartTime(ctx)
	timeElapsed := "0ms"

	if !reqStartTime.IsZero() {
		timeElapsed = fmt.Sprintf("%dms", uint64(time.Since(reqStartTime)/time.Millisecond))
	}

	return map[string]any{
		"request_id":      appcontext.GetRequestID(ctx),
		"user_agent":      appcontext.GetUserAgent(ctx),
		"service_version": appcontext.GetServiceVersion(ctx),
		"time_elapsed":    timeElapsed,
	}
}
