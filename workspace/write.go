package workspace

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/taubyte/office-space/common"
	. "github.com/taubyte/office-space/singletons"
)

const fourSpaces = "    "

func (w *workspace) write(vs common.VsWorkspace) (err error) {
	if vs.Folders == nil {
		vs.Folders = []common.VsFolder{}
	}

	// Remove duplicate folders
	seen := map[string]bool{}
	newFolders := []common.VsFolder{}
	for _, folder := range vs.Folders {
		if seen[folder.Path] {
			continue
		}

		seen[folder.Path] = true
		newFolders = append(newFolders, folder)
	}

	vs.Folders = newFolders

	// Avoid writing null into a JSON file
	if vs.Settings == nil {
		vs.Settings = map[string]interface{}{}
	}

	data, err := json.MarshalIndent(vs, "", fourSpaces)
	if err != nil {
		return fmt.Errorf("Marshall workspace failed with: %s", err)
	}

	if Runtime().Dry() == false {
		err = os.WriteFile(w.Loc(), data, 0777)
	}
	if err == nil {
		Display().WroteDirectoriesInFile(vs.List(), w.FileName(), w.Dir())
	}

	return
}

// Writes the vscode workspace file from a struct
func (w *workspace) Write(vs common.VsWorkspace) error {
	w.fileLock.Lock()
	defer w.fileLock.Unlock()

	return w.write(vs)
}
