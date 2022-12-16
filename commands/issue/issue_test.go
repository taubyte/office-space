package issue_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/taubyte/office-space/common"
	"github.com/taubyte/office-space/internal/mocks"
	. "github.com/taubyte/office-space/singletons"
)

func TestIssueBasic(t *testing.T) {
	ctx, err := mocks.CLI()
	if err != nil {
		t.Error(err)
		return
	}
	defer ctx.Close()

	err = ctx.FakeWorkspace("repo1", "repo2", "repo3", "repo4")
	if err != nil {
		t.Error(err)
		return
	}

	// Empty the workspace, as we'll be creating workspaces from the repos
	err = Workspace().Write(common.VsWorkspace{})
	if err != nil {
		t.Error(err)
		return
	}

	branch1 := "TP-1_some_1branch_stuff"
	branch2 := "TP-2_some_2branch_stuff"
	branch3 := "TP-3_some_3branch_stuff"
	branch4 := "TP-4_some_4branch_stuff"

	testData := map[string][]string{
		"repo1": {branch1, branch2, branch3},
		"repo2": {branch1, branch2, branch3},
		"repo3": {branch1, branch2, branch3},
		"repo4": {branch1, branch4},
	}
	paths := map[string]string{}
	for name, branches := range testData {
		paths[name], err = ctx.FakeModuleWithBranches(name, branches...)
		if err != nil {
			return
		}
	}

	confirmPathsSelectedBranch := func(branch string, paths ...string) error {
		for _, path := range paths {
			out, _, err := ctx.ExecuteCaptureInDir(path, "git", "branch")
			if err != nil {
				return err
			}

			if strings.Contains(out, "* "+branch) == false {
				return fmt.Errorf("Expected branch %s to be checked out on path %s, but it was not got: %s", branch, path, out)
			}
		}

		return nil
	}

	err = ctx.Run("issue", "TP-4")
	if err != nil {
		t.Error(err)
		return
	}

	err = confirmPathsSelectedBranch(branch4, paths["repo4"])
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.ConfirmInWorkspace("repo4")
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.Run("issue", "TP-1")
	if err != nil {
		t.Error(err)
		return
	}

	err = confirmPathsSelectedBranch(branch1, paths["repo1"], paths["repo2"], paths["repo3"], paths["repo4"])
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.ConfirmInWorkspace("repo1", "repo2", "repo3", "repo4")
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.Run("issue", "TP-3")
	if err != nil {
		t.Error(err)
		return
	}

	err = confirmPathsSelectedBranch(branch3, paths["repo1"], paths["repo2"], paths["repo3"])
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.ConfirmInWorkspace("repo1", "repo2", "repo3")
	if err != nil {
		t.Error(err)
		return
	}
}
