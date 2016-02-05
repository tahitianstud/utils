package logging

// Logger is an interface for different Logging implementations
type Logger interface {
	Trace(message string)
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Fatal(message string)
}

type loggerFactory func() Logger

// New permits the instantiation of a new Logger
var New loggerFactory
