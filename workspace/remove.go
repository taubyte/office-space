package workspace

import (
	"os"

	"github.com/taubyte/office-space/common"
)

func (ws *workspace) RemoveUse(dir string) error {
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
		for i, folder := range vs.Folders {
			if folder.Path == relativeDir {
				vs.Folders = append(vs.Folders[:i], vs.Folders[i+1:]...)
				break
			}
		}

		return nil
	})
}
