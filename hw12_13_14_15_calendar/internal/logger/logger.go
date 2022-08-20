package logger

import (
	"fmt"
	"time"
)

const dateLayout = "2006-01-02"

type Logger struct {
}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Info(msg string) {
	fmt.Printf("[INFO] %s: %s\n", time.Now().Format(dateLayout), msg)
}

func (l *Logger) Warn(msg string) {
	fmt.Printf("[WARNING] %s: %s\n", time.Now().Format(dateLayout), msg)
}

func (l *Logger) Debug(msg string) {
	fmt.Printf("[DEBUG] %s: %s\n", time.Now().Format(dateLayout), msg)
}

func (l *Logger) Error(msg string) {
	fmt.Printf("[ERROR] %s: %s\n", time.Now().Format(dateLayout), msg)
}
