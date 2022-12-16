package work_test

import (
	"testing"

	"github.com/taubyte/office-space/internal/mocks"
)

func TestWork(t *testing.T) {
	ctx, err := mocks.CLI()
	if err != nil {
		t.Error(err)
		return
	}
	defer ctx.Close()

	err = ctx.Run("init")
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.FakeModules("repo1", "repo2")
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.FakeWorkspace("repo1", "repo2")
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.Run("work")
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.ConfirmInGoWork("repo1", "repo2")
	if err != nil {
		t.Error(err)
		return
	}
}
