package workspace

import (
	"encoding/json"
	"os"

	"github.com/taubyte/office-space/common"
)

func (w *workspace) read() (vs common.VsWorkspace, err error) {
	data, err := os.ReadFile(w.Loc())
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &vs)
	return
}

// Reads the vscode workspace file into a struct
func (w *workspace) Read() (vs common.VsWorkspace, err error) {
	w.fileLock.RLock()
	defer w.fileLock.RUnlock()

	return w.read()
}
