package log

import "context"

type (
	StorageOpt struct {
		Driver   StorageDriver
		FileLocation string
		Rotation StorageRotation
	}
	StorageRotation struct {
		Enable     bool
		FileFormat func(ctx context.Context, cfg Config) string
	}
	StorageDriver string
)

const (
	STORAGE_DRIVER_CONSOLE = StorageDriver("CONSOLE")
	STORAGE_DRIVER_FILE    = StorageDriver("FILE")
)
