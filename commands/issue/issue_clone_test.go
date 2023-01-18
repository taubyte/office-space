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
	"gotest.tools/v3/assert"
)

func TestIssueClone(t *testing.T) {
	ctx, err := mocks.CLI()
	assert.NilError(t, err)
	defer ctx.Close()

	err = ctx.FakeWorkspace("repo1", "repo2", "repo3", "repo4")
	assert.NilError(t, err)

	// Empty the workspace, as we'll be creating workspaces from the repos
	err = Workspace().Write(common.VsWorkspace{})
	assert.NilError(t, err)

	err = GoWork().Init()
	assert.NilError(t, err)

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
		assert.NilError(t, err)
	}

	confirmPathsSelectedBranch := func(branch string, paths ...string) error {
		for _, path := range paths {
			out, _, err := ctx.ExecuteCaptureInDir(path, "git", "branch")
			assert.NilError(t, err)

			if strings.Contains(out, "* "+branch) == false {
				return fmt.Errorf("Expected branch %s to be checked out on path %s, but it was not got: %s", branch, path, out)
			}
		}

		return nil
	}

	issue.CloneRepositoryOnBranch = func(runtime_ctx *runtime.Context, dir, branch string, gitPrefix string) error {
		name := filepath.Base(dir)

		path, err := ctx.FakeModuleWithBranches(name, branch)
		assert.NilError(t, err)

		err = ctx.ExecuteInDir(path, "git", "checkout", branch)
		assert.NilError(t, err)

		return nil
	}

	err = ctx.Run("issue-clone", "TP-4")
	assert.NilError(t, err)

	testDirPath := func(issue string, repo string) string {
		return path.Join(mocks.TestDirectory, issue, repo)
	}

	err = confirmPathsSelectedBranch(branch4, testDirPath("TP-4", "repo4"))
	assert.NilError(t, err)

	err = ctx.ConfirmInWorkspace("repo4")
	assert.NilError(t, err)

	err = ctx.ConfirmInGoWork("repo4")
	assert.NilError(t, err)

	_, err = mocks.ResetTestDir()
	assert.NilError(t, err)

	err = ctx.Run("issue-clone", "TP-1")
	assert.NilError(t, err)

	err = confirmPathsSelectedBranch(
		branch1,
		testDirPath("TP-1", "repo1"),
		testDirPath("TP-1", "repo2"),
		testDirPath("TP-1", "repo3"),
		testDirPath("TP-1", "repo4"),
	)
	assert.NilError(t, err)

	err = ctx.ConfirmInWorkspace("repo1", "repo2", "repo3", "repo4")
	assert.NilError(t, err)

	err = ctx.ConfirmInGoWork("repo1", "repo2", "repo3", "repo4")
	assert.NilError(t, err)

	_, err = mocks.ResetTestDir()
	assert.NilError(t, err)

	err = ctx.Run("issue-clone", "TP-3")
	assert.NilError(t, err)

	err = confirmPathsSelectedBranch(
		branch3,
		testDirPath("TP-3", "repo1"),
		testDirPath("TP-3", "repo2"),
		testDirPath("TP-3", "repo3"),
	)
	assert.NilError(t, err)

	err = ctx.ConfirmInWorkspace("repo1", "repo2", "repo3")
	assert.NilError(t, err)

	err = ctx.ConfirmInGoWork("repo1", "repo2", "repo3")
	assert.NilError(t, err)
}
