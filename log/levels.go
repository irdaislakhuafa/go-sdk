package log

type (
	LEVEL string
)

// Reference: github.com/rs/zerolog
var (
	// LevelTraceValue is the value used for the trace level field.
	LEVEL_TRACE_VALUE = LEVEL("trace")
	// LevelDebugValue is the value used for the debug level field.
	LEVEL_DEBUG_VALUE = LEVEL("debug")
	// LevelInfoValue is the value used for the info level field.
	LEVEL_INFO_VALUE = LEVEL("info")
	// LevelWarnValue is the value used for the warn level field.
	LEVEL_WARN_VALUE = LEVEL("warn")
	// LevelErrorValue is the value used for the error level field.
	LEVEL_ERROR_VALUE = LEVEL("error")
	// LevelFatalValue is the value used for the fatal level field.
	LEVEL_FATAL_VALUE = LEVEL("fatal")
	// LevelPanicValue is the value used for the panic level field.
	LEVEL_PANIC_VALUE = LEVEL("panic")
)
