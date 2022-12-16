package workspace_test

import (
	"testing"

	"github.com/taubyte/office-space/internal/mocks"
	. "github.com/taubyte/office-space/singletons"
)

func TestInit(t *testing.T) {
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

	vs, err := Workspace().Read()
	if err != nil {
		t.Error(err)
		return
	}

	if len(vs.Folders) != 0 || len(vs.Settings) != 0 {
		t.Errorf("expected 0 folders and 0 settings, got %d and %d", len(vs.Folders), len(vs.Settings))
		return
	}
}
