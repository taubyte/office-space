package work_test

import (
	"os"
	"testing"

	"github.com/taubyte/office-space/internal/mocks"
)

func TestClean(t *testing.T) {
	ctx, err := mocks.CLI()
	if err != nil {
		t.Error(err)
		return
	}
	defer ctx.Close()

	err = ctx.FakeWorkspace("./repo1", "./repo2")
	if err != nil {
		t.Error(err)
		return
	}

	ctx.Run("work")
	ctx.Run("work", "clean")

	expectedGoWorkFile := "go 1.19\n"

	data, err := os.ReadFile(ctx.Dir + "/go.work")
	if err != nil {
		t.Error(err)
		return
	}

	if string(data) != expectedGoWorkFile {
		t.Errorf("\nExpected \n%s, \ngot \n%s", expectedGoWorkFile, string(data))
		return
	}
}
