package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Logger interface {
	Debug(module string, format string, args ...interface{})
	Info(module string, format string, args ...interface{})
	Warn(module string, format string, args ...interface{})
	Error(module string, format string, args ...interface{})
	Fatal(module string, format string, args ...interface{})

	SetLevel(level LogLevel)
	SetFormat(format string)
	SetOutput(w io.Writer)
}

type platformLogger struct {
	mu        sync.Mutex
	level     LogLevel
	formatter Formatter
	output    io.Writer
}

func (cl *platformLogger) logf(level LogLevel, module string, format string, args ...interface{}) {
	if level < cl.level {
		return
	}

	cl.mu.Lock()
	defer cl.mu.Unlock()

	msg := fmt.Sprintf(format, args...)
	entry := LogEntry{
		Timestamp: time.Now(),
		Module:    module,
		Level:     level.String(),
		Message:   msg,
	}

	line := cl.formatter.Format(entry)
	_, _ = cl.output.Write([]byte(line + "\n"))

	if level == FatalLevel {
		os.Exit(1)
	}
}

func (cl *platformLogger) Debug(module string, format string, args ...interface{}) {
	cl.logf(DebugLevel, module, format, args...)
}
func (cl *platformLogger) Info(module string, format string, args ...interface{}) {
	cl.logf(InfoLevel, module, format, args...)
}
func (cl *platformLogger) Warn(module string, format string, args ...interface{}) {
	cl.logf(WarnLevel, module, format, args...)
}
func (cl *platformLogger) Error(module string, format string, args ...interface{}) {
	cl.logf(ErrorLevel, module, format, args...)
}
func (cl *platformLogger) Fatal(module string, format string, args ...interface{}) {
	cl.logf(FatalLevel, module, format, args...)
}

func (cl *platformLogger) SetLevel(l LogLevel) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	cl.level = l
}
func (cl *platformLogger) SetFormat(name string) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	cl.formatter = getFormatter(name)
}
func (cl *platformLogger) SetOutput(w io.Writer) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	cl.output = w
}
