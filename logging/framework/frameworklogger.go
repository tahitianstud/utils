package framework

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/tahitianstud/utils/logging"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// frameworkLogger is an implementation of a Logger using a framework
type frameworkLogger struct {
	verboseMode bool
}

// Logger is a factory method for interface Logger
func Logger() logging.Logger {
	return &frameworkLogger{verboseMode: false}
}

func init() {
	logging.New = Logger
}

var needsInit = true

func initLogger() {
	if needsInit {
		logrus.SetOutput(os.Stdout)
		formatter := &prefixed.TextFormatter{ForceColors: true, TimestampFormat: "15:04:05.000"}

		logrus.SetFormatter(formatter)
		logrus.SetLevel(logrus.InfoLevel)

		needsInit = false
	}
}

// SetLevel sets the desired level
func (f frameworkLogger) SetLevel(level logging.LogLevel) {
	initLogger()
	logLevel := logging.LogLevel.String(level)
	parsedLevel, err := logrus.ParseLevel(logLevel)

	if err == nil {
		logrus.SetLevel(parsedLevel)
	} else {
		logrus.SetLevel(logrus.ErrorLevel)
	}
}

// ActivateVerboseOutput defines if the output contains a full set of information or not
func (f frameworkLogger) ActivateVerboseOutput(verboseFlag bool) {
	initLogger()

	if verboseFlag == false {
		formatter := &logrus.TextFormatter{DisableTimestamp: true}
		logrus.SetFormatter(formatter)
	}
}

// Trace prints out more Debug level messages if Verbose activated
func (f frameworkLogger) Trace(message string, a ... interface{}) {
	initLogger()
	if f.verboseMode {
		if len(a) > 0 {
			logrus.Debugf(message, a...)
		} else {
			logrus.Debug(message)
		}
	}
}

// Debug prints out a Debug level message
func (f frameworkLogger) Debug(message string, a ... interface{}) {
	initLogger()
	if len(a) > 0 {
		logrus.Debugf(message, a...)
	} else {
		logrus.Debug(message)
	}
}

// Info prints out a Info level message
func (f frameworkLogger) Info(message string, a ... interface{}) {
	initLogger()
	if len(a) > 0 {
		logrus.Infof(message, a...)
	} else {
		logrus.Info(message)
	}
}

// Warn prints out a Info level message but only when in Verbose mode
func (f frameworkLogger) Warn(message string, a ... interface{}) {
	initLogger()
	if len(a) > 0 {
		logrus.Warnf(message, a...)
	} else {
		logrus.Warn(message)
	}
}

// Die will exit after printing system error
func (f frameworkLogger) Error(message string, a ... interface{}) {
	initLogger()
	if len(a) > 0 {
		logrus.Errorf(message, a...)
	} else {
		logrus.Error(message)
	}
}

// Fatal exits after printing out critical error
func (f frameworkLogger) Fatal(message string, a ... interface{}) {
	initLogger()
	if len(a) > 0 {
		logrus.Fatalf(message, a...)
	} else {
		logrus.Fatal(message)
	}
}
