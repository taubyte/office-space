package issue

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/pterm/pterm"
	"github.com/taubyte/office-space/commands/work"
	"github.com/taubyte/office-space/common"
	"github.com/taubyte/office-space/env"
	"github.com/taubyte/office-space/runtime"
	. "github.com/taubyte/office-space/singletons"
	"github.com/urfave/cli/v2"

	goRuntime "runtime"
)

func issueCloneCommand() *cli.Command {
	return &cli.Command{
		Name:      "issue-clone",
		ArgsUsage: "<branch-prefix>",
		Usage:     "Used for checking out and cloning a git issue into a prefixed directory based on branch prefix",
		Action:    runtime.Wrap(issueClone),
	}
}

func issueSubCloneCommand() *cli.Command {
	c := issueCloneCommand()
	c.Name = "clone"

	return c
}

// Setting as a variable so that it can be overridden with a mkdir for tests
var CloneRepositoryOnBranch = func(ctx *runtime.Context, dir, branch, gitPrefix string) error {
	cloneUrl := path.Join(gitPrefix, filepath.Base(dir)+".git")

	err := ctx.ExecuteInDir(Workspace().Dir(), "git", "clone", cloneUrl, "-b", branch)
	if err != nil {
		return fmt.Errorf("git clone %s on branch %s failed with: %s", cloneUrl, branch, err)
	}

	return nil
}

func initializeIssueClone(ctx *runtime.Context) (branchPrefix, gitPrefix, newWsDir string, err error) {
	branchPrefix, err = getBranchPrefix(ctx)
	if err != nil {
		return
	}

	gitPrefix = env.Get().GitPrefix()
	if len(gitPrefix) == 0 {
		err = fmt.Errorf("environment variable %s not set", env.GitPrefixEnvKey)
		return
	}

	newWsDir = path.Join(Workspace().Dir(), branchPrefix)
	if ctx.Dry() == false {
		err = os.Mkdir(newWsDir, 0744)
		if err == nil {
			pterm.Success.Printf("Created dir %s\n", newWsDir)
		}
	}

	return
}

func issueClone(ctx *runtime.Context) error {
	branchPrefix, gitPrefix, newWsDir, err := initializeIssueClone(ctx)
	if err != nil {
		return err
	}

	// Get valid repo/branch combinations
	items, err := getValidBranchesWithPrefix(ctx, branchPrefix)

	if err != nil {
		return err
	}

	env.Set().WorkspaceDirectory(newWsDir)

	err = Workspace().New()
	if err != nil {
		return err
	}

	// Add items to workspace
	err = Workspace().Edit(func(vs *common.VsWorkspace) error {
		vs.Folders = []common.VsFolder{}
		for dir := range items {
			vs.Folders = append(vs.Folders, common.VsFolder{Path: dir})
		}

		return nil
	})
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(items))
	wg.Add(len(items))

	// TODO checkout most recent branch with prefix
	for dir, branch := range items {
		go func(_dir, _branch string) {
			errChan <- CloneRepositoryOnBranch(ctx, _dir, _branch, gitPrefix)
			wg.Done()
		}(dir, branch)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	// call $ asd work
	{
		err = work.Work(ctx, true)
		if err != nil {
			pterm.Warning.Printfln("Using provided repositories failed with: %s", err)
		}
	}

	err = Workspace().Edit(func(vs *common.VsWorkspace) error {
		vs.Settings = map[string]interface{}{
			"terminal.integrated.env." + goRuntime.GOOS: map[string]string{
				env.WorkspaceDirectoryEnvKey: Workspace().Dir(),
			},
		}
		return nil
	})
	if err != nil {
		return err
	}

	pterm.Info.Println(
		fmt.Sprintf("To use `asd` in this workspace be sure to run `export %s=%s`\n", env.WorkspaceDirectoryEnvKey, Workspace().Dir()),
		fmt.Sprintf("Or set your default %s in your ~/.bash_profile, and the workspace settings will override\n", env.WorkspaceDirectoryEnvKey),
	)

	return Workspace().OpenWithCode()
}
