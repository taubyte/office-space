package work_test

import (
	"os"
	"path"
	"testing"

	_ "embed"

	"github.com/taubyte/office-space/internal/mocks"
	. "github.com/taubyte/office-space/singletons"
)

//go:embed update_test_main.go_
var mainGo []byte

func TestUpdate(t *testing.T) {
	ctx, err := mocks.CLI()
	if err != nil {
		t.Error(err)
		return
	}
	defer ctx.Close()

	err = ctx.FakeWorkspace("repo1", "go-sdk")
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.ConfirmInWorkspace("repo1", "go-sdk")
	if err != nil {
		t.Error(err)
		return
	}

	repoPath, err := ctx.FakeModule("repo1")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = ctx.FakeModuleWithPackage("go-sdk", "bitbucket.org/taubyte/go-sdk")
	if err != nil {
		t.Error(err)
		return
	}

	err = os.WriteFile(path.Join(repoPath, "main.go"), mainGo, 0777)
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.ExecuteInDir(repoPath, "go", "get", "bitbucket.org/taubyte/go-sdk@v0.1.0")
	if err != nil {
		t.Error(err)
		return
	}

	mod, err := GoMod().Open(repoPath)
	if err != nil {
		t.Error(err)
		return
	}

	err = mod.File().AddReplace("bitbucket.org/taubyte/go-sdk", "", "bitbucket.org/taubyte/go-sdk", "./go-sdk")
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.Run("work")
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.ConfirmInGoWork("repo1", "go-sdk")
	if err != nil {
		t.Error(err)
		return
	}

	err = ctx.Run("work", "update", "go-sdk", "--no-git")
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
