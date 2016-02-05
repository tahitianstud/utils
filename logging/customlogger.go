package logging

import "fmt"

// customLogger is an implementation of a logger
type customLogger struct {
}

// NewCustomLogger is factory method for customLogger
func NewCustomLogger() Logger {
	return &customLogger{}
}

func init() {
	New = NewCustomLogger
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
