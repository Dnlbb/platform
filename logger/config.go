package logger

import (
	"io"
	"os"

	"github.com/natefinch/lumberjack"
)

type Config struct {
	Level      string `yaml:"level"`
	Format     string `yaml:"format"`
	Filename   string `yaml:"filename"`
	MaxSizeMB  int    `yaml:"max_size_mb"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAgeDays int    `yaml:"max_age_days"`
	Compress   bool   `yaml:"compress"`
}

func DefaultConfig() Config {
	return Config{
		Level:      "DEBUG",
		Format:     "text",
		Filename:   "logs/app.log",
		MaxSizeMB:  10,
		MaxBackups: 3,
		MaxAgeDays: 7,
		Compress:   false,
	}
}

func newLogger(cfg Config) Logger {
	lvl := parseLevel(cfg.Level)
	f := getFormatter(cfg.Format)

	var outputWriter io.Writer = os.Stdout
	if cfg.Filename != "" {
		outputWriter = &lumberjack.Logger{
			Filename:   cfg.Filename,
			MaxSize:    cfg.MaxSizeMB,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAgeDays,
			Compress:   cfg.Compress,
		}
	}

	return &platformLogger{
		level:     lvl,
		formatter: f,
		output:    outputWriter,
	}
}
