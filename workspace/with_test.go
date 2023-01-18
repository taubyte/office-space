package workspace_test

import (
	"fmt"
	"os"
	"path"
	"sort"
	"sync"
	"testing"

	"github.com/taubyte/office-space/common"
	"github.com/taubyte/office-space/internal/mocks"
	. "github.com/taubyte/office-space/singletons"
)

func TestWithAll(t *testing.T) {
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

	folders := []common.VsFolder{
		{Path: "./repo1"},
		{Path: "./repo2"},
		{Path: "./repo3"},
	}

	ws := Workspace()

	err = ws.Edit(func(vs *common.VsWorkspace) error {
		vs.Folders = folders

		return nil
	})
	if err != nil {
		t.Error(err)
		return
	}

	var allDirsLock sync.Mutex
	allDirs := sort.StringSlice{}
	err = ws.ForEach(func(dir string) error {
		allDirsLock.Lock()
		defer allDirsLock.Unlock()

		allDirs = append(allDirs, dir)
		return nil
	}).Absolute()
	if err != nil {
		t.Error(err)
		return
	}

	cwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}

	absTestDirectory := path.Join(cwd, mocks.TestDirectory)

	expectedFolders := []string{}
	for _, folder := range folders {
		expectedFolders = append(expectedFolders, path.Join(absTestDirectory, folder.Path))
	}

	if len(allDirs) != len(expectedFolders) {
		t.Errorf("expected %d folders, got %d", len(expectedFolders), len(allDirs))
		return
	}

	allDirs.Sort()
	for idx, dir := range allDirs {
		if dir != expectedFolders[idx] {
			t.Errorf("expected %s, got %s", expectedFolders[idx], dir)
			return
		}
	}

	err = ws.ForEach(func(dir string) error {
		return fmt.Errorf("failed")
	}).Absolute()
	if err == nil {
		t.Error("expected an error, got nil")
		return
	}
}

func TestWithRelative(t *testing.T) {
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

	folders := []common.VsFolder{
		{Path: "./repo1"},
		{Path: "./repo2"},
		{Path: "./repo3"},
	}

	ws := Workspace()

	err = ws.Edit(func(vs *common.VsWorkspace) error {
		vs.Folders = folders

		return nil
	})
	if err != nil {
		t.Error(err)
		return
	}

	var relativeDirsLock sync.Mutex
	relativeDirs := sort.StringSlice{}
	err = ws.ForEach(func(dir string) error {
		relativeDirsLock.Lock()
		defer relativeDirsLock.Unlock()

		relativeDirs = append(relativeDirs, dir)
		return nil
	}).Relative()
	if err != nil {
		t.Error(err)
		return
	}

	expectedFolders := []string{}
	for _, folder := range folders {
		expectedFolders = append(expectedFolders, folder.Path)
	}

	if len(relativeDirs) != len(expectedFolders) {
		t.Errorf("expected %d folders, got %d", len(expectedFolders), len(relativeDirs))
		return
	}

	relativeDirs.Sort()
	for idx, dir := range relativeDirs {
		if dir != expectedFolders[idx] {
			t.Errorf("expected %s, got %s", expectedFolders[idx], dir)
			return
		}
	}

	err = ws.ForEach(func(dir string) error {
		return fmt.Errorf("failed")
	}).Relative()
	if err == nil {
		t.Error("expected an error, got nil")
		return
	}
}
