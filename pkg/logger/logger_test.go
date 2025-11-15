package logger_test

import (
	"log/slog"
	"testing"

	"github.com/ssongin/core/pkg/logger"
)

func TestInitFromEnv_SetsExpectedLevel(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want slog.Level
	}{
		{"debug", "debug", slog.LevelDebug},
		{"info", "info", slog.LevelInfo},
		{"warn", "warn", slog.LevelWarn},
		{"error", "error", slog.LevelError},
		{"default_on_unknown", "unexpected", slog.LevelInfo},
		{"default_on_empty", "", slog.LevelInfo},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			logger.Reset()

			logger.SetLevelFromString(tc.in)

			if logger.GetLogLevel() != tc.want {
				t.Fatalf("SetLevel(%q) set logLevel = %v; want %v", tc.in, logger.GetLogLevel(), tc.want)
			}
		})
	}
}
