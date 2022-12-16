package go_work

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/taubyte/office-space/env"
	. "github.com/taubyte/office-space/singletons"
)

func (w *goWork) Remove() {
	Runtime().Remove(path.Join(w.Dir(), "go.work"))
	Runtime().Remove(path.Join(w.Dir(), "go.work.sum"))
}

func (w *goWork) Init() error {
	if len(w.Dir()) == 0 {
		return fmt.Errorf("env variable `%s` not defined", env.WorkspaceDirectoryEnvKey)
	}

	err := Runtime().ExecuteInDir(w.Dir(), "go", "work", "init")
	if err != nil {
		return err
	}

	return nil
}

func (w *goWork) RelativeTo(dir string) (string, error) {
	absPath, err := filepath.Abs(dir)
	if err != nil {
		return "", fmt.Errorf("getting absolute path from dir `%s` failed with: %s", dir, err)
	}

	relativePath, err := filepath.Rel(w.Dir(), absPath)
	if err != nil {
		return "", fmt.Errorf("getting relative path from (`%s` => `%s`) failed with: %s", w.Dir(), absPath, err)
	}

	return relativePath, nil
}
