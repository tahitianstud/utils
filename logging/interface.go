package logging

// Logger is an interface for different Logging implementations
type Logger interface {
	Trace(message string, a ...interface{})
	Debug(message string, a ...interface{})
	Info(message string, a ...interface{})
	Warn(message string, a ...interface{})
	Error(message string, a ...interface{})
	Fatal(message string, a ...interface{})

	SetLevel(level LogLevel)
	ActivateVerboseOutput(verboseFlag bool)
}

// LogLevel is a type used for expressing the log level
type LogLevel uint8

// loggerFactory is a factory method type
type loggerFactory func() Logger

// New permits the instantiation of a new Logger
var New loggerFactory

// TraceLevel is the indicator for trace messages
// DebugLevel is the indicator for debug messages
// InfoLevel is the indicator for info messages
// WarnLevel is the indicator for warning messages
// ErrorLevel is the indicator for error messages
// FatalLevel is the indicator for fatal messages
const (
	TraceLevel LogLevel = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

// String allows for conversion between string and integer representation of a level
func (l LogLevel) String() string {
	var outputLevel = ""
	switch l {
	case TraceLevel:
		outputLevel = "trace"
	case DebugLevel:
		outputLevel = "debug"
	case InfoLevel:
		outputLevel = "info"
	case WarnLevel:
		outputLevel = "warn"
	case ErrorLevel:
		outputLevel = "error"
	case FatalLevel:
		outputLevel = "fatal"
	}
	return outputLevel
}
