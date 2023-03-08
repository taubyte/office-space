package work

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/pterm/pterm"
	"github.com/taubyte/office-space/runtime"
	. "github.com/taubyte/office-space/singletons"
)

func remove(ctx *runtime.Context) error {
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

	// Remove from go.work
	err := GoWork().RemoveUse(dir)
	if err != nil {
		pterm.Warning.Printfln("Removing Relative dir `%s` from go.work failed with: %s", dir, err)
	}

	// Remove from vs workspace
	err = Workspace().RemoveUse(dir)
	if err != nil {
		return fmt.Errorf("Removing Relative dir `%s` from workspace failed with: %s", dir, err)
	}

	return Workspace().OpenWithCode()
}
