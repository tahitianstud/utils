package logging

import "fmt"

type customLogger struct {
}

// CustomLogger is factory method for customLogger
func CustomLogger() Logger {
	return &customLogger{}
}

func init() {
	New = CustomLogger
}

// SetLevel will set the desired level
func (c customLogger) SetLevel(level LogLevel) {
	fmt.Println("Setting level to " + LogLevel.String(level))
}

// ActivateVerboseOutput defines if the output contains a full set of information or not
func (c customLogger) ActivateVerboseOutput(verboseFlag bool) {
	fmt.Println("Activate the right verbose mode")
}

// Trace prints out more Debug level messages if Verbose activated
func (c customLogger) Trace(message string) {
	fmt.Println(message)
}

// Debug prints out a Debug level message
func (c customLogger) Debug(message string) {
	fmt.Println(message)
} // Debug prints out a Debug level message

// Info prints out a Info level message
func (c customLogger) Info(message string) {
	fmt.Println(message)
}

// Warn prints out a Info level message
func (c customLogger) Warn(message string) {
	fmt.Println(message)
}

// Error prints out a Info level message
func (c customLogger) Error(message string) {
	fmt.Println(message)
}

// Fatal exits after printing out critical error
func (c customLogger) Fatal(message string) {
	fmt.Println(message)
}
