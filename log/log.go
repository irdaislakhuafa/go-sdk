package log

import (
	"context"
	"fmt"

	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/rs/zerolog"
)

// Interface defines the methods that a logger must implement.
type Interface interface {
	Trace(ctx context.Context, obj interface{})
	Debug(ctx context.Context, obj interface{})
	Info(ctx context.Context, obj interface{})
	Warn(ctx context.Context, obj interface{})
	Error(ctx context.Context, obj interface{})
	Fatal(ctx context.Context, obj interface{})
	WithCtxFields(funcCtxField func(ctx context.Context) map[string]any) Interface
}

// Config holds the configuration for the logger.
type Config struct {
	Level          LEVEL
	SkipFrameCount int
	Storage        StorageOpt
}

// logger is an internal struct that holds the zerolog logger and context fields function.
type logger struct {
	log           zerolog.Logger
	funcCtxFields func(ctx context.Context) map[string]any
}

const (
	DEFAULT_SKIP_FRAME_COUNT = 3 // NOTE: temporary 3 for now
)

// Init initializes a new logger based on the provided configuration.
// It returns an Interface that logs to either console or file, depending on the config.
func Init(cfg Config) Interface {
	cfg.ParseDefault()

	switch cfg.Storage.Driver {
	case STORAGE_DRIVER_CONSOLE:
		return InitConsole(cfg)
	case STORAGE_DRIVER_FILE:
		return InitFile(cfg)
	default:
		panic(fmt.Sprintf("log storage driver '%v' not implemented!", cfg.Storage.Driver))
	}
}

// GetCaller extracts caller information from an error or returns the input as is.
// If the value is an error, it returns a formatted string with file, line, and message.
// If the value is a string, it returns the string as is.
// Otherwise, it returns the value formatted as a Go syntax representation.
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

// ParseDefault sets default values for the logger configuration if not already set.
// It sets the driver to CONSOLE if empty, skip frame count to default, and file location for file driver.
func (cfg *Config) ParseDefault() {
	if cfg.Storage.Driver == "" {
		cfg.Storage.Driver = STORAGE_DRIVER_CONSOLE
	}

	if cfg.SkipFrameCount <= 0 {
		cfg.SkipFrameCount = DEFAULT_SKIP_FRAME_COUNT
	}

	if cfg.Storage.Driver == STORAGE_DRIVER_FILE {
		cfg.Storage.FileLocation = "logs/log.json"
	}
}
