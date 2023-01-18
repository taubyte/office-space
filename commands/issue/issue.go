package issue

import (
	"fmt"
	"path"

	"github.com/pterm/pterm"
	"github.com/taubyte/office-space/commands/work"
	"github.com/taubyte/office-space/common"
	"github.com/taubyte/office-space/runtime"
	. "github.com/taubyte/office-space/singletons"
	"github.com/urfave/cli/v2"
)

func issueBasicCommand() *cli.Command {
	return &cli.Command{
		Name:        "issue",
		Usage:       "Used for checking out a git issue based on branch prefix",
		ArgsUsage:   "<branch-prefix>",
		Action:      runtime.Wrap(issue),
		Subcommands: []*cli.Command{issueSubCloneCommand()},
	}
}

func issue(ctx *runtime.Context) error {
	branchPrefix, err := getBranchPrefix(ctx)
	if err != nil {
		return err
	}

	err = Workspace().Write(common.VsWorkspace{})
	if err != nil {
		return fmt.Errorf("cleaning workspace failed with: %s", err)
	}

	items, err := getValidBranchesWithPrefix(ctx, branchPrefix)
	if err != nil {
		return err
	}

	Display().SetVerbose(false)

	// Checkout branches
	added := make([]string, len(items))
	var idx int
	for relativeDir, branch := range items {
		absoluteDir := path.Join(Workspace().Dir(), relativeDir)

		err = Workspace().AddUse(absoluteDir)
		if err != nil {
			return fmt.Errorf("adding `%s` to workspace failed with: %s", absoluteDir, err)
		}

		err = ctx.ExecuteInDir(absoluteDir, "git", "checkout", branch)
		if err != nil {
			return fmt.Errorf("checking out branch: %s in dir: %s failed with: %s", branch, absoluteDir, err)
		}

		added[idx] = relativeDir
		idx++
	}

	Display().SetVerbose(true)

	// call $ asd work
	{
		err = work.Work(ctx, true)
		if err != nil {
			pterm.Warning.Printfln("Using provided repositories failed with: %s", err)
		}
	}

	// Display the changes to main workspace as we hushed the verbose printouts for each iteration
	Display().WroteDirectoriesInFile(added, Workspace().FileName(), Workspace().Dir())

	return Workspace().OpenWithCode()
}
