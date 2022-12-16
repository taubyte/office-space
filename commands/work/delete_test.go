package work_test

import (
	"os"
	"testing"

	"github.com/taubyte/office-space/internal/mocks"
)

func TestDelete(t *testing.T) {
	ctx, err := mocks.CLI()
	if err != nil {
		t.Error(err)
		return
	}
	defer ctx.Close()

	err = ctx.FakeWorkspace("repo1")
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.Run("work")
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.Run("work", "d")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = os.ReadFile(ctx.Dir + "/go.work")
	if err == nil {
		t.Error("Expected error")
		return
	} else {
		err = nil
	}
}
