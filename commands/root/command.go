package root

import (
	"fmt"
	"path/filepath"

	"github.com/taubyte/office-space/env"
	"github.com/taubyte/office-space/runtime"
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:   "root",
		Usage:  "Sets the environment variable for the work directory to args[0]",
		Action: runtime.Wrap(command),
	}
}

func command(ctx *runtime.Context) error {
	root := ctx.Args().First()
	if len(root) == 0 {
		return fmt.Errorf("root variable required, ex: `asd root .`")
	}

	absPath, err := filepath.Abs(root)
	if err != nil {
		return fmt.Errorf("Getting absolute path failed with: %s", err)
	}

	fmt.Printf("export %s=%s\n", env.WorkspaceDirectoryEnvKey, absPath)

	return nil
}
