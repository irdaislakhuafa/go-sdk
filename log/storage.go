package log

import "context"

// StorageOpt holds options for log storage.
type StorageOpt struct {
	Driver       StorageDriver
	FileLocation string
	Rotation     StorageRotation
}

// StorageRotation holds options for log file rotation.
type StorageRotation struct {
	Enable     bool
	FileFormat func(ctx context.Context, cfg Config) string
}

// StorageDriver represents the type of storage driver for logs.
type StorageDriver string

const (
	// STORAGE_DRIVER_CONSOLE specifies the console storage driver.
	STORAGE_DRIVER_CONSOLE = StorageDriver("CONSOLE")
	// STORAGE_DRIVER_FILE specifies the file storage driver.
	STORAGE_DRIVER_FILE = StorageDriver("FILE")
)
