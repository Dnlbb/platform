package logger

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
	JsonFORMAT = "json"
	UmlFORMAT  = "uml"
	TextFORMAT = "text"
)

type Formatter interface {
	Format(entry LogEntry) string
}

type LogEntry struct {
	Timestamp time.Time
	Module    string
	Level     string
	Message   string
}

type textFormatter struct{}

func (tf *textFormatter) Format(e LogEntry) string {
	t := e.Timestamp.Format(TimeFormat)
	return fmt.Sprintf("%s [%s] %s", t, e.Module, e.Message)
}

type jsonFormatter struct{}

func (jf *jsonFormatter) Format(e LogEntry) string {
	data := map[string]interface{}{
		"timestamp": e.Timestamp.Format(time.RFC3339Nano),
		"module":    e.Module,
		"level":     e.Level,
		"message":   e.Message,
	}
	b, err := json.Marshal(data)
	if err != nil {
		return `{"ERROR":"json marshal failed"}`
	}
	return string(b)
}

type umlFormatter struct{}

func (uf *umlFormatter) Format(e LogEntry) string {
	sb := &strings.Builder{}
	sb.WriteString("@startuml\n")
	sb.WriteString(fmt.Sprintf("title %s [%s] %s\n",
		e.Timestamp.Format("2006-01-02T15:04:05.000Z"),
		e.Module,
		e.Level,
	))
	sb.WriteString(fmt.Sprintf("%s -> Logger: %s\n", e.Module, e.Message))
	sb.WriteString("@enduml")
	return sb.String()
}

func getFormatter(name string) Formatter {
	switch strings.ToLower(name) {
	case JsonFORMAT:
		return &jsonFormatter{}
	case UmlFORMAT:
		return &umlFormatter{}
	case TextFORMAT:
		return &textFormatter{}
	default:
		return &textFormatter{}
	}
}
