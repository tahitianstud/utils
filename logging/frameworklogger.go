package logging

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// frameworkLogger is an implementation of a Logger using a framework
type frameworkLogger struct {
	verboseMode bool
}

// FrameworkLogger is a factory method for interface Logger
func FrameworkLogger() Logger {
	return &frameworkLogger{verboseMode: false}
}

func init() {
	New = FrameworkLogger
}

var needsInit = true

func initLogger() {
	if needsInit {
		fmt.Println("Needs init")

		log.SetOutput(os.Stdout)
		formatter := &prefixed.TextFormatter{ForceColors: true, TimestampFormat: "15:04:05.000"}

		log.SetFormatter(formatter)
		log.SetLevel(log.InfoLevel)

		needsInit = false
	}
}

// SetLevel sets the desired level
func (f frameworkLogger) SetLevel(level LogLevel) {
	initLogger()
	logLevel := LogLevel.String(level)
	parsedLevel, err := log.ParseLevel(logLevel)

	if err == nil {
		log.SetLevel(parsedLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
		fmt.Println("Got an error: '" + err.Error() + "', setting log level to ERROR")
	}
}

// ActivateVerboseOutput defines if the output contains a full set of information or not
func (f frameworkLogger) ActivateVerboseOutput(verboseFlag bool) {
	fmt.Println("Activate Verbose Output")
	initLogger()

	if verboseFlag == false {
		formatter := &log.TextFormatter{DisableTimestamp: true}
		log.SetFormatter(formatter)
	}
}

// Trace prints out more Debug level messages if Verbose activated
func (f frameworkLogger) Trace(message string) {
	initLogger()
	if f.verboseMode {
		log.Debug(message)
	}
}

// Debug prints out a Debug level message
func (f frameworkLogger) Debug(message string) {
	initLogger()
	log.Debug(message)
}

// Info prints out a Info level message
func (f frameworkLogger) Info(message string) {
	initLogger()
	log.Info(message)
}

// Warn prints out a Info level message but only when in Verbose mode
func (f frameworkLogger) Warn(message string) {
	initLogger()
	log.Warn(message)
}

// Die will exit after printing system error
func (f frameworkLogger) Error(message string) {
	initLogger()
	log.Error(message)
}

// Fatal exits after printing out critical error
func (f frameworkLogger) Fatal(message string) {
	initLogger()
	log.Fatal(message)
}
