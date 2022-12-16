package work_test

import (
	"testing"

	"github.com/taubyte/office-space/internal/mocks"
	. "github.com/taubyte/office-space/singletons"
)

func TestBuild(t *testing.T) {
	ctx, err := mocks.CLI()
	if err != nil {
		t.Error(err)
		return
	}
	defer ctx.Close()

	err = ctx.FakeModules("repo1", "repo2")
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.FakeWorkspace("./repo1", "./repo2")
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.Run("work")
	if err != nil {
		t.Error(err)
		return
	}

	err = Workspace().Delete()
	if err != nil {
		t.Error(err)
		return
	}

	err = Workspace().New()
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.Run("work", "build")
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.ConfirmInWorkspace("repo1", "repo2")
	if err != nil {
		t.Error(err)
		return
	}
}
