package cli

import (
	"fmt"

	"github.com/taubyte/office-space/commands/air"
	"github.com/taubyte/office-space/commands/autocomplete"
	"github.com/taubyte/office-space/commands/initialize"
	"github.com/taubyte/office-space/commands/issue"
	"github.com/taubyte/office-space/commands/root"
	"github.com/taubyte/office-space/commands/run"
	"github.com/taubyte/office-space/commands/work"
	"github.com/taubyte/office-space/workspace"
	"github.com/urfave/cli/v2"
)

func defineCLI() (*cli.App, error) {
	globalFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "dry",
			Usage:   "Show which commands will be executed",
			Aliases: []string{"d"},
		},

		&cli.BoolFlag{
			Name:    "nocode",
			Aliases: []string{"n"},
			Usage:   fmt.Sprintf(`"This prevents running "code %s"`, workspace.PreLoc()),
		},
	}

	app := &cli.App{
		UseShortOptionHandling: true,
		Flags:                  globalFlags,
		EnableBashCompletion:   true,
	}

	app.Commands = []*cli.Command{
		run.Command(),
		work.Command(),
		initialize.Command(),
		air.Command(),
		root.Command(),
		autocomplete.Command(),
	}
	app.Commands = append(app.Commands, issue.Commands()...)

	// Generate valid commands to give ability for default command
	validCommands = make([]string, len(app.Commands))
	for idx, cmd := range app.Commands {
		validCommands[idx] = cmd.Name
	}

	return app, nil
}
