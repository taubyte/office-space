package work

import (
	"fmt"

	"github.com/taubyte/office-space/runtime"
	"github.com/taubyte/office-space/workspace"
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "work",
		Usage: fmt.Sprintf("Generates a new `go.work` from %s", workspace.PreLoc()),
		Action: runtime.Wrap(func(ctx *runtime.Context) error {
			return Work(ctx, false)
		}),
		Subcommands: []*cli.Command{
			{
				Name:    "clean",
				Aliases: []string{"c"},
				Usage:   "Removes all replaces from `go.work`",
				Action:  runtime.Wrap(clean),
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "Deletes `go.work`",
				Action:  runtime.Wrap(delete),
			},
			{
				Name:    "build",
				Aliases: []string{"b"},
				Usage:   fmt.Sprintf("Builds workspace `%s` with replaces from `go.work`", workspace.PreLoc()),
				Action:  runtime.Wrap(build),
			},
			{
				Name:      "add",
				Aliases:   []string{"a"},
				Usage:     fmt.Sprintf("Adds arg[0] to workspace and replaces `%s` in `go.work`", workspace.PreLoc()),
				Action:    runtime.Wrap(add),
				ArgsUsage: "<relative path | absolute path | relative to workspace path>",
			},
			{
				Name:      "remove",
				Aliases:   []string{"rm"},
				Usage:     fmt.Sprintf("Removes arg[0] from workspace and removes replace for `%s` from `go.work`", workspace.PreLoc()),
				Action:    runtime.Wrap(remove),
				ArgsUsage: "<relative path | absolute path | relative to workspace path>",
			},
			{
				Name:        "update",
				Aliases:     []string{"u"},
				Usage:       fmt.Sprintf("Removes replaces of a given package, removes from go.work, removes from %s, and updates versions throughout to latest", workspace.PreLoc()),
				Subcommands: updateCommands(),
			},
		},
	}
}
