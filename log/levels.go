package log

// LEVEL represents a log level as a string.
type LEVEL string

// Reference: github.com/rs/zerolog
var (
	// LEVEL_TRACE_VALUE is the string representation for the trace log level.
	LEVEL_TRACE_VALUE = LEVEL("trace")
	// LEVEL_DEBUG_VALUE is the string representation for the debug log level.
	LEVEL_DEBUG_VALUE = LEVEL("debug")
	// LEVEL_INFO_VALUE is the string representation for the info log level.
	LEVEL_INFO_VALUE = LEVEL("info")
	// LEVEL_WARN_VALUE is the string representation for the warn log level.
	LEVEL_WARN_VALUE = LEVEL("warn")
	// LEVEL_ERROR_VALUE is the string representation for the error log level.
	LEVEL_ERROR_VALUE = LEVEL("error")
	// LEVEL_FATAL_VALUE is the string representation for the fatal log level.
	LEVEL_FATAL_VALUE = LEVEL("fatal")
	// LEVEL_PANIC_VALUE is the string representation for the panic log level.
	LEVEL_PANIC_VALUE = LEVEL("panic")
)
