package work_test

import (
	"testing"

	"github.com/taubyte/office-space/internal/mocks"
	. "github.com/taubyte/office-space/singletons"
)

func TestAdd(t *testing.T) {
	ctx, err := mocks.CLI()
	if err != nil {
		t.Error(err)
		return
	}
	defer ctx.Close()

	err = Workspace().New()
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
}
