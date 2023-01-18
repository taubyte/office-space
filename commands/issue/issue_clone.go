package issue

import (
	"fmt"
	"os"
	"path"
	"strings"
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
var CloneRepositoryOnBranch = func(ctx *runtime.Context, dir string, branch string) error {
	parser, err := GitConfig().Open(dir + "/" + branch + "/.git/config")
	url, err := parser.Remote()
	if err != nil {
		return err
	}

	fmt.Print(url.String())

	err = ctx.ExecuteInDir(Workspace().Dir(), "git", "clone", url.String(), "-b", branch)
	if err != nil {
		return fmt.Errorf("git clone %s on branch %s failed with: %s", url.String(), branch, err)
	}

	return nil
}

func issueClone(ctx *runtime.Context) error {
	// branchPrefix, err := getBranchPrefix(ctx)
	// if err != nil {
	// 	return err
	// }
	//How do I grab the path here so I can create the parser?
	//Is there a way to execute pwd in using the ctx?
	parser, err := GitConfig().Open("/.git/config")
	if err != nil {
		return err
	}

	branchPrefixUrl, err := parser.Remote()
	if err != nil {
		return err
	}

	branchPrefix := branchPrefixUrl.String()

	gitPrefix := env.Get().GitPrefix()
	if len(gitPrefix) == 0 {
		return fmt.Errorf("environment variable %s not set", env.GitPrefixEnvKey)
	}

	newWsDir := path.Join(Workspace().Dir(), branchPrefix)
	if !ctx.Dry() {
		err = os.Mkdir(newWsDir, 0744)
		if err == nil {
			pterm.Success.Printf("Created dir %s\n", newWsDir)
		} else {
			return err
		}
	}

	// Get valid repo/branch combinations
	var itemsLock sync.Mutex
	items := make(map[string]string)
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
				relativeDir, err := Workspace().RelativeTo(dir)
				if err != nil {
					relativeDir = dir
				}
				itemsLock.Lock()
				items[relativeDir] = branch
				itemsLock.Unlock()

				break
			}
		}

		return nil
	}).Root()
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
			errChan <- CloneRepositoryOnBranch(ctx, _dir, _branch)
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
