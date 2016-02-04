package logging

import (
	"os"

	log "github.com/Sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var needsInit = true
var verboseMode = false

// InitLogger initializes the logging system
func InitLogger() {
	log.SetOutput(os.Stdout)
	prefixedFormatter := &prefixed.TextFormatter{ForceColors: true, TimestampFormat: "15:04:05.000"}

	log.SetFormatter(prefixedFormatter)
	log.SetLevel(log.InfoLevel)

	needsInit = false
}

// VerboseMode activates the verbose mode, allowing more traces to be outputted
func VerboseMode() {
	if needsInit {
		InitLogger()
	}
	verboseMode = true
}

// DebugMode activate the debug mode, lowering the log level to print out debug traces
func DebugMode() {
	if needsInit {
		InitLogger()
	}
	log.SetLevel(log.DebugLevel)

}

// Info prints out a Info level message
func Info(message string) {
	if needsInit {
		InitLogger()
	}
	log.Info(message)
}

// Step prints out a Info level message but only when in Verbose mode
func Step(message string) {
	if needsInit {
		InitLogger()
	}
	if verboseMode == true {
		log.Info(message)
	}
}

// Debug prints out a Debug level message
func Debug(message string) {
	if needsInit {
		InitLogger()
	}
	log.Debug(message)
} // Debug prints out a Debug level message

// Trace prints out more Debug level messages if Verbose activated
func Trace(message string) {
	if needsInit {
		InitLogger()
	}
	if verboseMode {
		log.Debug(message)
	}
}

// Fatal exits after printing out critical error
func Fatal(message string) {
	if needsInit {
		InitLogger()
	}
	log.Fatal(message)
}

// Die will exit after printing system error
func Die(message string) {
	if needsInit {
		InitLogger()
	}
	log.Panic(message)
}
