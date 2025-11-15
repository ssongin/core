package errors

import (
	"os"

	"github.com/ssongin/core/pkg/logger"
)

func CheckWarn(message string, err error, args ...any) {
	if err != nil {
		allArgs := append([]any{"error", err}, args...)
		logger.GetLogger().Warn(message, allArgs...)
	}
}

func CheckError(message string, err error, args ...any) {
	if err != nil {
		allArgs := append([]any{"error", err}, args...)
		logger.GetLogger().Error(message, allArgs...)
		panic(err)
	}
}

func CheckFatalError(message string, err error, args ...any) {
	if err != nil {
		allArgs := append([]any{"error", err}, args...)
		logger.GetLogger().Error(message, allArgs...)
		os.Exit(1)
	}
}
