package logging

// Logger is an interface for different Logging implementations
type Logger interface {
	Trace(message string)
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Fatal(message string)

	SetLevel(level LogLevel)
	ActivateVerboseOutput(verboseFlag bool)
}

// LogLevel is a type used for expressing the log level
type LogLevel uint8

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
