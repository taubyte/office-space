package workspace_test

import (
	"testing"

	"github.com/taubyte/office-space/common"
	"github.com/taubyte/office-space/internal/mocks"
	. "github.com/taubyte/office-space/singletons"
)

func TestWrite(t *testing.T) {
	ctx, err := mocks.CLI()
	if err != nil {
		t.Error(err)
		return
	}
	defer ctx.Close()

	err = Workspace().Write(common.VsWorkspace{
		Folders: []common.VsFolder{
			{Path: "repo1"},
			{Path: "repo2"},
			{Path: "repo3"},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
}
