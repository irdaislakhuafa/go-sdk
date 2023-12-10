package log

import (
	"context"
	"fmt"
	"os"
	"sync"

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
	panic("not implemented") // TODO: Implement
}

func (l *logger) Debug(ctx context.Context, obj interface{}) {
	panic("not implemented") // TODO: Implement
}

func (l *logger) Info(ctx context.Context, obj interface{}) {
	panic("not implemented") // TODO: Implement
}

func (l *logger) Warn(ctx context.Context, obj interface{}) {
	panic("not implemented") // TODO: Implement
}

func (l *logger) Error(ctx context.Context, obj interface{}) {
	panic("not implemented") // TODO: Implement
}

func (l *logger) Fatal(ctx context.Context, obj interface{}) {
	panic("not implemented") // TODO: Implement
}

func GetCaller(value any) any {
	switch tErr := value.(type) {
	case error:
		file, line, message, err := errors.GetCaller(tErr)
		if err != nil {
			return fmt.Sprintf("%s:%#v --- %s", file, line, message)
		}
	case string:
		return tErr
	default:
		return fmt.Sprintf("#%v", tErr)
	}

	return value
}

// func getContextFields(ctx context.Context) map[string]any {

// }
