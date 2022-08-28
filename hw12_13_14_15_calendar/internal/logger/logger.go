package logger

import (
	"fmt"
	"time"
)

const dateLayout = "2006-01-02"

// Logger ...
type Logger struct{}

// New ...
func New() *Logger {
	return &Logger{}
}

// Info ...
func (l *Logger) Info(msg string) {
	fmt.Printf("[INFO] %s: %s\n", time.Now().Format(dateLayout), msg)
}

// Warn ...
func (l *Logger) Warn(msg string) {
	fmt.Printf("[WARNING] %s: %s\n", time.Now().Format(dateLayout), msg)
}

// Debug ...
func (l *Logger) Debug(msg string) {
	fmt.Printf("[DEBUG] %s: %s\n", time.Now().Format(dateLayout), msg)
}

// Error ...
func (l *Logger) Error(msg string) {
	fmt.Printf("[ERROR] %s: %s\n", time.Now().Format(dateLayout), msg)
}
