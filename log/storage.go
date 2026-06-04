package log

// StorageOpt holds options for log storage.
type StorageOpt struct {
	Driver       StorageDriver
	FileLocation string
	Rotation     StorageRotation
}

// StorageRotation holds options for log file rotation.
type StorageRotation struct {
	// Enable storage rotation, if disabled then other options on storage rotation will be ignored.
	Enable bool

	// Megabytes before rolling
	MaxSize int

	// Maximum number of old log files to retain
	MaxBackups int

	// Maximum number of days to retain old log files
	MaxAge int

	// Compress old log files (gzip)
	Compress bool
}

// StorageDriver represents the type of storage driver for logs.
type StorageDriver string

const (
	// STORAGE_DRIVER_CONSOLE specifies the console storage driver.
	STORAGE_DRIVER_CONSOLE = StorageDriver("CONSOLE")
	// STORAGE_DRIVER_FILE specifies the file storage driver.
	STORAGE_DRIVER_FILE = StorageDriver("FILE")
)
