package logging

import "fmt"

// CustomLogger is an implementation of a logger
type CustomLogger struct {
}

// Trace prints out more Debug level messages if Verbose activated
func (c CustomLogger) Trace(message string) {
	fmt.Println(message)
}

// Debug prints out a Debug level message
func (c CustomLogger) Debug(message string) {
	fmt.Println(message)
} // Debug prints out a Debug level message

// Info prints out a Info level message
func (c CustomLogger) Info(message string) {
	fmt.Println(message)
}

// Warn prints out a Info level message
func (c CustomLogger) Warn(message string) {
	fmt.Println(message)
}

// Error prints out a Info level message
func (c CustomLogger) Error(message string) {
	fmt.Println(message)
}

// Fatal exits after printing out critical error
func (c CustomLogger) Fatal(message string) {
	fmt.Println(message)
}
