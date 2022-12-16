package run

import (
	"fmt"

	"github.com/taubyte/office-space/runtime"
	. "github.com/taubyte/office-space/singletons"
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:   "run",
		Usage:  "Default command, runs arguments as a command in each workspace directory",
		Action: runtime.Wrap(command),
	}
}

func command(ctx *runtime.Context) error {
	args := ctx.Args().Slice()

	if len(args) == 0 {
		return fmt.Errorf("Invalid arguments for run, expected a command.")
	}

	name := args[0]
	if len(args) > 1 {
		args = args[1:]
	} else {
		args = nil
	}

	return Workspace().ForEach(func(dir string) error {
		return ctx.ExecuteInDir(dir, name, args...)
	}).Absolute()
}
