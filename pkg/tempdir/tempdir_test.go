package tempdir_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ssongin/core/pkg/tempdir"
)

func TestTempdir_CreateAndCleanup(t *testing.T) {
	base := t.TempDir()
	path := filepath.Join(base, "subdir1", "subdir2")
	absPath, cleanup := tempdir.CreateTempPath(path)
	// path should exist
	if fi, err := os.Stat(absPath); err != nil {
		t.Fatalf("expected path to exist after CreateTempPath: %v", err)
	} else if !fi.IsDir() {
		t.Fatalf("expected path to be a directory")
	}
	// run cleanup and ensure removed
	cleanup()
	if _, err := os.Stat(absPath); !os.IsNotExist(err) {
		t.Fatalf("expected path to be removed after cleanup, stat err=%v", err)
	}
}

func TestTempdir_CleanupAll(t *testing.T) {
	base := t.TempDir()
	p1, _ := tempdir.CreateTempPath(filepath.Join(base, "a"))
	p2, _ := tempdir.CreateTempPath(filepath.Join(base, "b"))
	// ensure they exist
	if _, err := os.Stat(p1); err != nil {
		t.Fatalf("expected p1 to exist: %v", err)
	}
	if _, err := os.Stat(p2); err != nil {
		t.Fatalf("expected p2 to exist: %v", err)
	}
	// call CleanupAll and assert both removed
	tempdir.CleanupAll()
	if _, err := os.Stat(p1); !os.IsNotExist(err) {
		t.Fatalf("expected p1 to be removed after CleanupAll, stat err=%v", err)
	}
	if _, err := os.Stat(p2); !os.IsNotExist(err) {
		t.Fatalf("expected p2 to be removed after CleanupAll, stat err=%v", err)
	}
}
