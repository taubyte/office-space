package workspace

import (
	"os"
	"path"
	"path/filepath"

	"github.com/taubyte/office-space/env"
)

// Used in help commands
func PreLoc() string {
	env := env.Get()

	dir := path.Join(env.WorkspaceDirectory(), env.WorkspaceName()+env.WorkspaceExt())

	cwd, err := os.Getwd()
	if err != nil {
		return dir
	}

	relativeDir, err := filepath.Rel(cwd, dir)
	if err != nil {
		return dir
	}

	return relativeDir
}
