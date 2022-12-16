package workspace_test

import (
	"testing"

	"github.com/taubyte/office-space/common"
	"github.com/taubyte/office-space/internal/mocks"
	. "github.com/taubyte/office-space/singletons"
)

func TestMethods(t *testing.T) {
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

	err = Workspace().Edit(func(vs *common.VsWorkspace) error {
		vs.Folders = []common.VsFolder{
			{Path: "repo1"},
			{Path: "repo2"},
			{Path: "repo3"},
		}

		return nil
	})
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
