package initialize

import (
	"fmt"

	"github.com/taubyte/office-space/common"
	"github.com/taubyte/office-space/runtime"
	. "github.com/taubyte/office-space/singletons"
	"github.com/taubyte/office-space/workspace"
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:   "init",
		Usage:  fmt.Sprintf("Initializes a vscode workspace at: %s", workspace.PreLoc()),
		Action: runtime.Wrap(command),
	}
}

func command(ctx *runtime.Context) (err error) {
	err = Workspace().Stat()
	if err != nil {
		err = Workspace().Write(common.VsWorkspace{})
	}

	if err == nil {
		err = Workspace().OpenWithCode()
	}

	return
}
