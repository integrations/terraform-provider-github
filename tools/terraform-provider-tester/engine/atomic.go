package engine

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func writeAtomic(path string, perm fs.FileMode, write func(io.Writer) error) (err error) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}

	f, err := os.CreateTemp(dir, "."+filepath.Base(path)+".*.tmp")
	if err != nil {
		return err
	}
	tmp := f.Name()
	defer func() {
		_ = f.Close()
		if err != nil {
			_ = os.Remove(tmp)
		}
	}()

	if err = f.Chmod(perm); err != nil {
		return err
	}
	if err = write(f); err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}

	err = os.Rename(tmp, path)
	return err
}
