package work

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/taubyte/office-space/runtime"
	. "github.com/taubyte/office-space/singletons"
)

func add(ctx *runtime.Context) error {
	dir := ctx.Args().First()
	if dir == "" {
		return errors.New("Must provide an argument, ex: `work use .`")
	}

	if filepath.IsAbs(dir) == false {
		// Try relative
		_, err := os.Stat(dir)
		if err != nil {
			// Try relative to workspace
			dir = path.Join(Workspace().Dir(), dir)
		}
	}

	// Add to go.work
	err := GoWork().AddUse(dir)
	if err != nil {
		return fmt.Errorf("Adding relative `%s` to go.work failed with: %s", dir, err)
	}

	// Add it to vs workspace
	err = Workspace().AddUse(dir)
	if err != nil {
		return fmt.Errorf("Adding relative `%s` to workspace failed with: %s", dir, err)
	}

	return Workspace().OpenWithCode()
}
