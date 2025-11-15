package tempdir

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/ssongin/core/pkg/errors"
)

var (
	mu      sync.Mutex
	created []string
)

func CreateTempPath(path string) (string, func()) {
	absPath, err := filepath.Abs(path)
	errors.CheckError("failed to resolve absolute path", err)

	err = os.MkdirAll(absPath, 0o755)
	errors.CheckError("failed to create directories", err)

	mu.Lock()
	created = append(created, absPath)
	mu.Unlock()

	cleanup := func() {
		_ = os.RemoveAll(absPath)
	}

	return absPath, cleanup
}

func CleanupAll() {
	mu.Lock()
	defer mu.Unlock()

	for _, dir := range created {
		_ = os.RemoveAll(dir)
	}
	created = nil
}
