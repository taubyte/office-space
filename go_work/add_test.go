package go_work_test

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
}
