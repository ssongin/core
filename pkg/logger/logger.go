package logger

import (
	"log/slog"
	"os"
	"sync"
)

var (
	once     sync.Once
	logLevel slog.Level
	logger   *slog.Logger
)

func Reset() {
	once = sync.Once{}
	logger = nil
	logLevel = -10
}

func GetLogLevel() slog.Level {
	return logLevel
}

func SetLevelFromString(level string) {
	switch level {
	case "debug":
		SetLevel(slog.LevelDebug)
	case "info":
		SetLevel(slog.LevelInfo)
	case "warn":
		SetLevel(slog.LevelWarn)
	case "error":
		SetLevel(slog.LevelError)
	default:
		SetLevel(slog.LevelInfo)
	}
}

func SetLevel(level slog.Level) {
	logLevel = level
}

func GetLogger() *slog.Logger {
	once.Do(func() {
		handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: logLevel,
		})
		logger = slog.New(handler)
	})
	return logger
}
