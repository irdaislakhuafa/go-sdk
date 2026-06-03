package log

import (
	"context"
	"fmt"

	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/rs/zerolog"
)

type Interface interface {
	// TODO: added method Debugf
	Trace(ctx context.Context, obj interface{})
	Debug(ctx context.Context, obj interface{})
	Info(ctx context.Context, obj interface{})
	Warn(ctx context.Context, obj interface{})
	Error(ctx context.Context, obj interface{})
	Fatal(ctx context.Context, obj interface{})
	WithCtxFields(funcCtxField func(ctx context.Context) map[string]any) Interface
}

type (
	Config struct {
		Level          LEVEL
		SkipFrameCount int
		Storage        StorageOpt
	}
	logger struct {
		log           zerolog.Logger
		funcCtxFields func(ctx context.Context) map[string]any
	}
)

const (
	DEFAULT_SKIP_FRAME_COUNT = 3 // NOTE: temporary 3 for now
)

func Init(cfg Config) Interface {
	cfg.ParseDefault()

	switch cfg.Storage.Driver {
	case STORAGE_DRIVER_CONSOLE:
		return InitConsole(cfg)
	default:
		panic(fmt.Sprintf("log storage driver '%v' not implemented!", cfg.Storage.Driver))
	}
}

func GetCaller(value any) any {
	switch tErr := value.(type) {
	case error:
		file, line, message, err := errors.GetCaller(tErr)
		if err != nil {
			return fmt.Sprintf("%v", err)
		}
		return fmt.Sprintf("%s:%#v --- %s", file, line, message)
	case string:
		return tErr
	default:
		return fmt.Sprintf("%#v", tErr)
	}
}

func (cfg *Config) ParseDefault() {
	if cfg.Storage.Driver == "" {
		cfg.Storage.Driver = STORAGE_DRIVER_CONSOLE
	}

	if cfg.SkipFrameCount <= 0 {
		cfg.SkipFrameCount = DEFAULT_SKIP_FRAME_COUNT
	}
}
