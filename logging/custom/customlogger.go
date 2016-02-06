package custom

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/logutils"
	"github.com/tahitianstud/utils/logging"
	"github.com/tahitianstud/utils/logging/terminal"
)

// package variables
var (
	needsInit = true
	logLevels = []logutils.LogLevel{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

// my logger type
type customLogger struct {
}

// Logger is a factory method for customLogger
func Logger() logging.Logger {
	return &customLogger{}
}

func init() {
	logging.New = Logger
}

func initLogger() {
	if needsInit {
		fmt.Println("Initializing the logger")
		filter := &logutils.LevelFilter{
			Levels:   []logutils.LogLevel(logLevels),
			MinLevel: logutils.LogLevel("WARN"),
			Writer:   os.Stderr,
		}
		log.SetOutput(filter)
		log.SetFlags(0)

		fmt.Printf("filter is: %v\n", filter)

		needsInit = false
	}
}

// SetLevel will set the desired level
func (c customLogger) SetLevel(level logging.LogLevel) {
	initLogger()

	logLevel := strings.Trim(strings.ToUpper(logging.LogLevel.String(level)), "")
	fmt.Printf("setting level to : #%s#\n", logLevel)

	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel(logLevels),
		MinLevel: logutils.LogLevel(logLevel),
		Writer:   os.Stderr,
	}

	log.SetOutput(filter)
}

// ActivateVerboseOutput defines if the output contains a full set of information or not
func (c customLogger) ActivateVerboseOutput(verboseFlag bool) {
	initLogger()

	if verboseFlag {
		log.SetFlags(log.LstdFlags)
	}
}

// Trace prints out more Debug level messages if Verbose activated
func (c customLogger) Trace(message string) {

	initLogger()
	log.Printf("%s %s", terminal.Colorize("[TRACE]", terminal.CYAN), message)
}

// Debug prints out a Debug level message
func (c customLogger) Debug(message string) {
	initLogger()
	log.Printf("%s %s", terminal.Colorize("[DEBUG]", terminal.GREEN), message)
}

// Info prints out a Info level message
func (c customLogger) Info(message string) {
	initLogger()

	log.Printf("%s %s", terminal.Colorize("[INFO]", terminal.BLUE), message)
}

// Warn prints out a Info level message
func (c customLogger) Warn(message string) {
	initLogger()
	log.Printf("%s %s", terminal.Colorize("[WARN]", terminal.YELLOW), message)
}

// Error prints out a Info level message
func (c customLogger) Error(message string) {
	initLogger()
	log.Printf("%s %s", terminal.Colorize("[ERROR]", terminal.MAGENTA), message)
}

// Fatal exits after printing out critical error
func (c customLogger) Fatal(message string) {
	initLogger()
	log.Printf("%s %s", terminal.Colorize("[FATAL]", terminal.RED), message)
}
