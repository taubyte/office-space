package cli

import (
	"context"
	"os"

	"github.com/pterm/pterm"
	"github.com/taubyte/office-space/display"
	"github.com/taubyte/office-space/go_mod"
	"github.com/taubyte/office-space/go_work"
	"github.com/taubyte/office-space/runtime"
	. "github.com/taubyte/office-space/singletons"
	"github.com/taubyte/office-space/workspace"
)

func Initialize(ctx context.Context) (*App, error) {
	var err error

	// Initialize the vscode workspace file
	SingletonWorkspace, err = workspace.Initialize()
	if err != nil {
		if IsHelpCommand() == false {
			pterm.Warning.Println(err)
		}
	}

	// Initialize the go work file
	SingletonGoWork, err = go_work.Initialize()
	if err != nil {
		return nil, err
	}

	SingletonGoMod, err = go_mod.Initialize()
	if err != nil {
		return nil, err
	}

	SingletonDisplay, err = display.Initialize()
	if err != nil {
		return nil, err
	}

	err = runtime.Initialize()
	if err != nil {
		return nil, err
	}
	cli, err := defineCLI()
	if err != nil {
		return nil, err
	}

	// Show help when only running `asd`
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "-h")
	}

	validFlags := getValidFlags(cli.Flags)

	structureDefaultCommand()
	os.Args = movePostfixOptions(os.Args, validFlags)

	return &App{cli}, nil
}
