package mocks

import (
	"fmt"
	"os"
	"path"

	"github.com/taubyte/office-space/common"
	. "github.com/taubyte/office-space/singletons"
)

func (m *mockCLI) FakeModule(name string) (path string, err error) {
	return m.FakeModuleWithPackage(name, "bitbucket.org/function/lib")
}

func (m *mockCLI) FakeModuleWithPackage(name, _package string) (filePath string, err error) {
	filePath = path.Join(Workspace().Dir(), name)
	os.Mkdir(filePath, 0777)

	err = m.ExecuteInDir(filePath, "go", "mod", "init", _package)
	if err != nil {
		return
	}

	return
}

func (m *mockCLI) FakeModules(names ...string) (err error) {
	for _, name := range names {
		_, err = m.FakeModule(name)
		if err != nil {
			return err
		}
	}

	return
}

func (m *mockCLI) FakeModuleWithBranches(name string, branches ...string) (path string, err error) {
	path, err = m.FakeModule(name)
	if err != nil {
		return
	}

	err = m.Runtime.ExecuteInDir(path, "git", "init")
	if err != nil {
		return
	}

	err = m.Runtime.ExecuteInDir(path, "git", "-c", "user.name='taubyte'", "-c", "user.email='test@taubyte.com'", "commit", "--allow-empty", "-m", `"empty"`)
	if err != nil {
		return
	}

	for _, branch := range branches {
		err = m.Runtime.ExecuteInDir(path, "git", "branch", branch)
		if err != nil {
			return
		}
	}

	return
}

func (m *mockCLI) FakeWorkspace(paths ...string) error {
	err := Workspace().New()
	if err != nil {
		return err
	}

	return Workspace().Edit(func(vs *common.VsWorkspace) error {
		vs.Folders = []common.VsFolder{}
		for _, path := range paths {
			vs.Folders = append(vs.Folders, common.VsFolder{Path: path})
		}

		return nil
	})
}

func (m *mockCLI) ConfirmInWorkspace(paths ...string) error {
	ws, err := Workspace().Read()
	if err != nil {
		return err
	}

	for _, path := range paths {
		found := false
		for _, folder := range ws.Folders {
			if folder.Path == path {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("path `%s` not found in workspace: %#v", path, ws.Folders)
		}
	}

	return nil
}

func (m *mockCLI) ConfirmInGoWork(paths ...string) error {
	work, err := GoWork().Read()
	if err != nil {
		return err
	}

	if len(work.Use) != len(paths) {
		return fmt.Errorf("expected %d paths, got %d", len(paths), len(work.Use))
	}

	for _, path := range paths {
		found := false
		for _, using := range work.Use {
			if using.Path == path {
				found = true
				break
			}
		}

		if !found {
			using := []string{}
			for _, path := range work.Use {
				using = append(using, path.Path)
			}

			return fmt.Errorf("path `%s` not found in GoWork: %#v", path, using)
		}
	}

	return nil
}
