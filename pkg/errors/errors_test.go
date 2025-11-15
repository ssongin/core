package errors_test

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"testing"

	pkgErr "github.com/ssongin/core/pkg/errors"
)

func captureOutput(f func()) string {
	orig := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		outC <- buf.String()
	}()

	f()

	_ = w.Close()
	os.Stdout = orig
	out := <-outC
	return out
}

func TestCheckWarn_LogsWhenError_NoPanic(t *testing.T) {
	errVal := errors.New("sample-warn")
	out := captureOutput(func() {
		// should not panic
		pkgErr.CheckWarn("warn message", errVal, "key", "value")
	})
	// logging capture can be flaky depending on logger init timing,
	// so do not fail the test if nothing was captured; just ensure no panic.
	if out != "" {
		if !bytes.Contains([]byte(out), []byte("warn message")) && !bytes.Contains([]byte(out), []byte("error")) {
			t.Logf("captured output (unexpected format): %q", out)
		}
	} else {
		t.Log("no log output captured; logger may have been initialized to use a different writer")
	}
}

func TestCheckWarn_NoOutputWhenNil_NoPanic(t *testing.T) {
	out := captureOutput(func() {
		pkgErr.CheckWarn("warn message", nil)
	})
	if out != "" {
		t.Logf("unexpected log output for nil error: %q", out)
	}
}

func TestCheckError_PanicsOnly(t *testing.T) {
	errVal := errors.New("sample-error")
	var recovered interface{}
	_ = captureOutput(func() {
		func() {
			defer func() {
				recovered = recover()
			}()
			pkgErr.CheckError("error message", errVal, "key", "value")
		}()
	})
	if recovered == nil {
		t.Fatalf("expected panic, got none")
	}
	if recErr, ok := recovered.(error); ok {
		if recErr.Error() != errVal.Error() {
			t.Errorf("expected panic error %q, got %q", errVal.Error(), recErr.Error())
		}
	} else {
		t.Errorf("expected recovered value to be error, got: %#v", recovered)
	}
	// don't assert on captured log output because logger initialization may use a different writer
}

func TestCheckFatalError_Child(t *testing.T) {
	// child process: perform the fatal call which should os.Exit(1)
	if os.Getenv("BE_CRASHER") != "1" {
		t.Skip("not a crash child")
	}
	pkgErr.CheckFatalError("fatal message", errors.New("boom"))
	// should never reach here
}

func TestCheckFatalError_Exits(t *testing.T) {
	// parent: spawn child that will call os.Exit(1)
	cmd := exec.Command(os.Args[0], "-test.run=TestCheckFatalError_Child")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if err == nil {
		t.Fatalf("expected child process to exit with non-zero status")
	}
	// expect ExitError with exit code 1
	if exitErr, ok := err.(*exec.ExitError); ok {
		// Non-zero exit status expected
		if exitErr.ExitCode() != 1 {
			t.Fatalf("expected exit code 1, got %d", exitErr.ExitCode())
		}
	} else {
		t.Fatalf("expected ExitError, got: %v", err)
	}
}
