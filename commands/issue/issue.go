package issue

import (
	"fmt"
	"strings"

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

	Display().SetVerbose(false)

	err = Workspace().Write(common.VsWorkspace{})
	if err != nil {
		return fmt.Errorf("cleaning workspace failed with: %s", err)
	}

	added := []string{}
	err = Workspace().ForEach(func(dir string) error {
		err := ctx.ExecuteInDir(dir, "git", "fetch")
		if err != nil {
			pterm.Warning.Printfln("likely not a git repository: `%s` git fetch failed with: %s", dir, err)

			return nil
		}

		// Get branches
		out, errOut, err := ctx.ExecuteCaptureInDir(dir, "git", "for-each-ref", "--format=%(refname:short)")
		if err != nil {
			return fmt.Errorf("getting branch of `%s`, failed with: %s\nError Output:\n%s", dir, err, errOut)
		}

		// TODO checkout most recent branch with prefix
		branches := strings.Split(out, "\n")
		for _, branch := range branches {
			_branchPrefix := branchPrefix
			if strings.HasPrefix(branch, "origin/") {
				_branchPrefix = "origin/" + _branchPrefix
			}

			if strings.HasPrefix(branch, _branchPrefix) {
				err = Workspace().AddUse(dir)
				if err != nil {
					return fmt.Errorf("adding `%s` to workspace failed with: %s", dir, err)
				}

				relativeDir, err := Workspace().RelativeTo(dir)
				if err != nil {
					relativeDir = dir
				}

				added = append(added, relativeDir)

				err = ctx.ExecuteInDir(dir, "git", "checkout", branch)
				if err != nil {
					return fmt.Errorf("checking out branch: %s in dir: %s failed with: %s", branch, dir, err)
				}

				break
			}
		}

		return nil
	}).Root()
	if err != nil {
		return err
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
