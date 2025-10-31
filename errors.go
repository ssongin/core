package core

import (
	"log/slog"
	"os"
)

func CheckError(message string, logger *slog.Logger, err error) {
	if err != nil {
		logger.Error(message, "error", err)
		panic(err)
	}
}

func CheckFatalError(message string, logger *slog.Logger, err error) {
	if err != nil {
		logger.Error(message, "error", err)
		os.Exit(1)
	}
}
