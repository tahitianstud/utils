package logging

import (
	"os"

	log "github.com/Sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// FrameworkLogger is an implementation of a Logger using a framework
type FrameworkLogger struct {
	verboseMode bool
}

// FrameworkLoggerInstance is a constructor
func FrameworkLoggerInstance(verbose bool) *FrameworkLogger {
	logger := new(FrameworkLogger)
	logger.verboseMode = verbose
	return logger
}

var needsInit = true

func initLogger() {
	if needsInit {
		log.SetOutput(os.Stdout)
		prefixedFormatter := &prefixed.TextFormatter{ForceColors: true, TimestampFormat: "15:04:05.000"}

		log.SetFormatter(prefixedFormatter)
		log.SetLevel(log.InfoLevel)

		needsInit = false
	}
}

// DebugMode activate the debug mode, lowering the log level to print out debug traces
func (f FrameworkLogger) DebugMode() {
	initLogger()
	log.SetLevel(log.DebugLevel)
}

// Print prints out arguments
func (f FrameworkLogger) Print(stype string, args ...interface{}) {
	initLogger()
	log.Printf(stype+" ==> %s", args)
}

// Info prints out a Info level message
func (f FrameworkLogger) Info(message string) {
	initLogger()
	log.Info(message)
}

// Step prints out a Info level message but only when in Verbose mode
func (f FrameworkLogger) Step(message string) {
	initLogger()
	if f.verboseMode == true {
		log.Info(message)
	}
}

// Debug prints out a Debug level message
func (f FrameworkLogger) Debug(message string) {
	initLogger()
	log.Debug(message)
} // Debug prints out a Debug level message

// Trace prints out more Debug level messages if Verbose activated
func (f FrameworkLogger) Trace(message string) {
	initLogger()
	if f.verboseMode {
		log.Debug(message)
	}
}

// Fatal exits after printing out critical error
func (f FrameworkLogger) Fatal(message string) {
	initLogger()
	log.Fatal(message)
}

// Die will exit after printing system error
func (f FrameworkLogger) Die(message string) {
	initLogger()
	log.Panic(message)
}
