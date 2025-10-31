package core

import (
	"os"
)

func CheckError(message string, err error) {
	if err != nil {
		GetLogger().Error(message, "error", err)
		panic(err)
	}
}

func CheckFatalError(message string, err error) {
	if err != nil {
		GetLogger().Error(message, "error", err)
		os.Exit(1)
	}
}
