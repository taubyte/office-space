package workspace

import (
	"errors"
	"fmt"
	"sync"

	"github.com/taubyte/office-space/common"
)

func (ws *workspace) ForEach(handler common.DirHandler) common.RunWith {
	return &runWith{ws, handler}
}

/*
Runs the command with dir being the absolute path for every folder defined in a workspace

Example:
> /home/user/Documents/
  - main.code-workspace: {"repo1", "repo2"}
  - repo1
  - repo2
  - repo3
    ...

Absolute Directories => ["/home/user/Documents/repo1", "/home/user/Documents/repo2"]
*/
func (rw *runWith) Absolute() error {
	return rw.dirs(dirsConfig{}, rw.handler)
}

/*
Runs the command with dir being the relative path for every folder defined in a workspace

Example:
> /home/user/Documents/
  - main.code-workspace: {"repo1", "repo2"}
  - repo1
  - repo2
  - repo3
    ...

Relative Directories => ["repo1", "repo2"]
*/
func (rw *runWith) Relative() error {
	return rw.dirs(dirsConfig{
		relative: true,
	}, rw.handler)
}

/*
Runs the command with dir being the absolute path for every folder defined in the workspace's directory

Example:
> /home/user/Documents/
  - main.code-workspace: {"repo1", "repo2"}
  - repo1
  - repo2
  - repo3
    ...

Root Directories => ["/home/user/Documents/repo1", "/home/user/Documents/repo2", "/home/user/Documents/repo3", ...]
*/
func (rw *runWith) Root() error {
	return rw.dirs(dirsConfig{
		all: true,
	}, rw.handler)
}

func (rw *runWith) dirs(c dirsConfig, f func(dir string) error) error {
	dirs, err := rw.ws.directories(c)
	if err != nil {
		return err
	}

	var errorString string

	errChan := make(chan error, len(dirs))
	var wg sync.WaitGroup
	wg.Add(len(dirs))

	for _, dir := range dirs {
		go func(_dir string) {
			defer wg.Done()
			err := f(_dir)
			if err != nil {
				errChan <- fmt.Errorf("In dir: %s :: %s", _dir, err.Error())
			}
		}(dir)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			errorString += "\n" + err.Error()
		}
	}

	if len(errorString) != 0 {
		return errors.New(errorString)
	}

	return nil
}
