package go_work_test

import (
	"testing"

	"github.com/taubyte/office-space/internal/mocks"
	. "github.com/taubyte/office-space/singletons"
)

func TestRemove(t *testing.T) {
	ctx, err := mocks.CLI()
	if err != nil {
		t.Error(err)
		return
	}
	defer ctx.Close()

	repoPath, err := ctx.FakeModule("repo1")
	if err != nil {
		t.Error(err)
		return
	}

	err = GoWork().AddUse(repoPath)
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.ConfirmInGoWork("repo1")
	if err != nil {
		t.Error(err)
		return
	}

	// Copied above from add_test.go

	err = GoWork().RemoveUse(repoPath)
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.ConfirmInGoWork()
	if err != nil {
		t.Error(err)
		return
	}
}
