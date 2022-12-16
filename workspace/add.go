package workspace

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/taubyte/office-space/common"
)

/*
workspace AddUse takes a directory path (dir) that can be one of the following:
  - relative to cwd
  - relative to the workspace
  - absolute

returning an error after adding the directory to the workspace
*/
func (ws *workspace) AddUse(dir string) error {
	_, err := os.Stat(dir)
	if err != nil {
		return err
	}

	// Get directory relative to ws.Dir from the abs path
	relativeDir, err := ws.RelativeTo(dir)
	if err != nil {
		return err
	}

	return ws.Edit(func(vs *common.VsWorkspace) error {
		vs.Folders = append(vs.Folders, common.VsFolder{Path: relativeDir})

		return nil
	})
}

func (ws *workspace) RelativeTo(dir string) (string, error) {
	absPath, err := filepath.Abs(dir)
	if err != nil {
		return "", fmt.Errorf("getting absolute path from dir `%s` failed with: %s", dir, err)
	}

	relativePath, err := filepath.Rel(ws.Dir(), absPath)
	if err != nil {
		return "", fmt.Errorf("getting relative path from (`%s` => `%s`) failed with: %s", ws.Dir(), absPath, err)
	}

	return relativePath, nil
}
