package work_test

import (
	"testing"

	"github.com/taubyte/office-space/internal/mocks"
)

func TestRemove(t *testing.T) {
	ctx, err := mocks.CLI()
	if err != nil {
		t.Error(err)
		return
	}
	defer ctx.Close()

	err = ctx.FakeWorkspace()
	if err != nil {
		t.Error(err)
		return
	}

	repoPath, err := ctx.FakeModule("repo1")
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.Run("work", "add", repoPath)
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.ConfirmInGoWork("repo1")
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.ConfirmInWorkspace("repo1")
	if err != nil {
		t.Error(err)
		return
	}

	// The above is copied from TestAdd

	err = ctx.Run("work", "remove", repoPath)
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.ConfirmInGoWork()
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.ConfirmInWorkspace()
	if err != nil {
		t.Error(err)
		return
	}
}
