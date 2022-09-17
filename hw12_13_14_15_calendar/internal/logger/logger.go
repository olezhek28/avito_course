package logger

import (
	"fmt"
	"strings"
	"time"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/config"
)

const dateLayout = "2006-01-02"

// Logger ...
type Logger struct {
	showTime bool
}

// New ...
func New(config *config.LoggerConf) *Logger {
	return &Logger{
		showTime: config.ShowTime,
	}
}

// Info ...
func (l *Logger) Info(msg string, a ...any) {
	l.print("[INFO]", msg, a...)
}

// Warn ...
func (l *Logger) Warn(msg string, a ...any) {
	l.print("[WARNING]", msg, a...)
}

// Debug ...
func (l *Logger) Debug(msg string, a ...any) {
	l.print("[DEBUG]", msg, a...)
}

// Error ...
func (l *Logger) Error(msg string, a ...any) {
	l.print("[ERROR]", msg, a...)
}

func (l *Logger) print(level string, msg string, a ...any) {
	now := time.Now().Format(dateLayout)
	var str strings.Builder
	str.WriteString(level)
	str.WriteString(" ")

	if l.showTime {
		str.WriteString(now)
		str.WriteString(": ")
	}

	str.WriteString(fmt.Sprintf(msg, a...))

	fmt.Println(str.String())
}
