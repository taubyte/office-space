package workspace

import (
	"os"

	"github.com/taubyte/office-space/common"
)

func (w *workspace) New() (err error) {
	file, err := os.Create(w.Loc())
	if err != nil {
		return
	}
	file.Close()

	return w.Write(common.VsWorkspace{
		Folders:  []common.VsFolder{},
		Settings: map[string]interface{}{},
	})
}
