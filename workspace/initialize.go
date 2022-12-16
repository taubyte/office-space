package workspace

import (
	"fmt"
	"os"

	"github.com/taubyte/office-space/common"
)

func Initialize() (ws common.Workspace, err error) {
	ws = &workspace{}

	err = ws.Stat()
	if err != nil {
		// Ignoring init warning as we are running init
		if len(os.Args) > 1 && os.Args[1] == "init" {
			return ws, nil
		}

		return ws, fmt.Errorf("Workspace not initialized, run `init` :: %s", err)
	}

	return
}
