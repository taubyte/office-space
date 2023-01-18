package issue

import (
	"fmt"
	"strings"
	"sync"

	"github.com/pterm/pterm"
	"github.com/taubyte/office-space/runtime"
	. "github.com/taubyte/office-space/singletons"
)

func getBranchPrefix(ctx *runtime.Context) (branchPrefix string, err error) {
	branchPrefix = ctx.Args().First()
	if len(branchPrefix) == 0 {
		err = fmt.Errorf("Must provide a branch prefix, ex: `asd issue <prefix>`")
	}

	return
}

func getValidBranchesWithPrefix(ctx *runtime.Context, branchPrefix string) (items map[string]string, err error) {
	var itemsLock sync.Mutex
	items = make(map[string]string)

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

		// TODO checkout most recent branch with prefix as an issue could have multiple branches on one repository
		// TODO confirm branch is still active
		branches := strings.Split(out, "\n")
		for _, branch := range branches {
			// Remove origin/ to prevent checking out detached commit
			branch = strings.TrimPrefix(branch, "origin/")

			if strings.HasPrefix(branch, branchPrefix) {
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

	return
}
