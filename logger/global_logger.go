package logger

import (
	"io"
)

var lg Logger

func InitLogger(cfg Config) {
	lg = newLogger(cfg)
}

func Debug(module string, format string, args ...interface{}) {
	if lg != nil {
		lg.Debug(module, format, args...)
	}
}

func Info(module string, format string, args ...interface{}) {
	if lg != nil {
		lg.Info(module, format, args...)
	}
}

func Warn(module string, format string, args ...interface{}) {
	if lg != nil {
		lg.Warn(module, format, args...)
	}
}
func Error(module string, format string, args ...interface{}) {
	if lg != nil {
		lg.Error(module, format, args...)
	}
}
func Fatal(module string, format string, args ...interface{}) {
	if lg != nil {
		lg.Fatal(module, format, args...)
	}
}

func SetLevel(level string) {
	if lg != nil {
		lg.SetLevel(parseLevel(level))
	}
}
func SetFormat(fmt string) {
	if lg != nil {
		lg.SetFormat(fmt)
	}
}
func SetOutput(w io.Writer) {
	if lg != nil {
		lg.SetOutput(w)
	}
}
