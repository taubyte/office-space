package issue_test

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/taubyte/office-space/commands/issue"
	"github.com/taubyte/office-space/common"
	"github.com/taubyte/office-space/internal/mocks"
	"github.com/taubyte/office-space/runtime"
	. "github.com/taubyte/office-space/singletons"
)

func TestIssueClone(t *testing.T) {
	ctx, err := mocks.CLI()
	if err != nil {
		t.Error(err)
		return
	}

	defer ctx.Close() //comment this line out to see generated file structure

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

	err = GoWork().Init()
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
				return fmt.Errorf("%s git branch failed with: %s", path, err)
			}

			if strings.Contains(out, "* "+branch) == false {
				return fmt.Errorf("Expected branch %s to be checked out on path %s, but it was not got: %s", branch, path, out)
			}
		}

		return nil
	}

	issue.CloneRepositoryOnBranch = func(runtime_ctx *runtime.Context, dir, branch string, gitPrefix string) error {
		name := filepath.Base(dir)

		path, err := ctx.FakeModuleWithBranches(name, branch)
		if err != nil {
			return err
		}

		err = ctx.ExecuteInDir(path, "git", "checkout", branch)
		if err != nil {
			return err
		}

		return nil
	}

	err = ctx.Run("issue-clone", "TP-4")
	if err != nil {
		t.Error(err)
		return
	}

	testDirPath := func(issue string, repo string) string {
		return path.Join(mocks.TestDirectory, issue, repo)
	}

	err = confirmPathsSelectedBranch(branch4, testDirPath("TP-4", "repo4"))
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.ConfirmInWorkspace("repo4")
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.ConfirmInGoWork("repo4")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = mocks.ResetTestDir()
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.Run("issue-clone", "TP-1")
	if err != nil {
		t.Error(err)
		return
	}

	err = confirmPathsSelectedBranch(
		branch1,
		testDirPath("TP-1", "repo1"),
		testDirPath("TP-1", "repo2"),
		testDirPath("TP-1", "repo3"),
		testDirPath("TP-1", "repo4"),
	)
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.ConfirmInWorkspace("repo1", "repo2", "repo3", "repo4")
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.ConfirmInGoWork("repo1", "repo2", "repo3", "repo4")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = mocks.ResetTestDir()
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.Run("issue-clone", "TP-3")
	if err != nil {
		t.Error(err)
		return
	}

	err = confirmPathsSelectedBranch(
		branch3,
		testDirPath("TP-3", "repo1"),
		testDirPath("TP-3", "repo2"),
		testDirPath("TP-3", "repo3"),
	)
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.ConfirmInWorkspace("repo1", "repo2", "repo3")
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.ConfirmInGoWork("repo1", "repo2", "repo3")
	if err != nil {
		t.Error(err)
		return
	}
}
