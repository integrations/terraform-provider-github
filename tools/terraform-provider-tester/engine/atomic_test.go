package engine

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteAtomicPreservesTargetOnWriteError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "target")
	if err := os.WriteFile(path, []byte("old"), 0o600); err != nil {
		t.Fatalf("seed target: %v", err)
	}

	wantErr := errors.New("write failed")
	err := writeAtomic(path, 0o600, func(io.Writer) error { return wantErr })
	if !errors.Is(err, wantErr) {
		t.Fatalf("error = %v, want %v", err, wantErr)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read target: %v", err)
	}
	if string(got) != "old" {
		t.Fatalf("target = %q, want old", got)
	}

	matches, err := filepath.Glob(filepath.Join(filepath.Dir(path), ".*.tmp"))
	if err != nil {
		t.Fatalf("glob temp files: %v", err)
	}
	if len(matches) != 0 {
		t.Fatalf("leftover temp files: %v", matches)
	}
}

func TestWriteAtomicWritesPrivateFileAndCleansUpTemps(t *testing.T) {
	path := filepath.Join(t.TempDir(), "target")

	err := writeAtomic(path, 0o600, func(w io.Writer) error {
		_, err := io.WriteString(w, "new")
		return err
	})
	if err != nil {
		t.Fatalf("writeAtomic: %v", err)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat target: %v", err)
	}
	if got := info.Mode().Perm(); got != 0o600 {
		t.Fatalf("mode = %v, want 0600", got)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read target: %v", err)
	}
	if string(got) != "new" {
		t.Fatalf("target = %q, want new", got)
	}

	matches, err := filepath.Glob(filepath.Join(filepath.Dir(path), ".*.tmp"))
	if err != nil {
		t.Fatalf("glob temp files: %v", err)
	}
	if len(matches) != 0 {
		t.Fatalf("leftover temp files: %v", matches)
	}
}

func TestWriteAtomicCreatesNestedDirectories(t *testing.T) {
	path := filepath.Join(t.TempDir(), "nested", "target")

	err := writeAtomic(path, 0o600, func(w io.Writer) error {
		_, err := io.WriteString(w, strings.Repeat("x", 3))
		return err
	})
	if err != nil {
		t.Fatalf("writeAtomic: %v", err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read target: %v", err)
	}
	if string(got) != "xxx" {
		t.Fatalf("target = %q, want xxx", got)
	}
}
