package core

import (
	"log/slog"
	"os"
	"sync"
)

var (
	once   sync.Once
	Logger *slog.Logger
)

func GetLogger() *slog.Logger {
	once.Do(func() {
		Logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	})
	return Logger
}
