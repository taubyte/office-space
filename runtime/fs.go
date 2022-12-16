package runtime

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/taubyte/office-space/common"
)

func (c *Context) WriteFile(name string, data []byte, perm fs.FileMode) error {
	c.DisplayExec("touch", []string{name}...)
	if c.Bool("dry") == true {
		return nil
	}

	return os.WriteFile(name, data, perm)
}

func (c *Context) RemoveAll(path string) error {
	c.DisplayExec("rm", []string{"-r", path}...)
	if c.Bool("dry") == true {
		return nil
	}

	return os.RemoveAll(path)
}

func (c *Context) Remove(path string) error {
	c.DisplayExec("rm", []string{path}...)
	if c.Bool("dry") == true {
		return nil
	}

	return os.Remove(path)
}

func (c *Context) Chdir(dir string) (r common.ReturnToDirMethod, err error) {
	c.DisplayExec("cd", dir)

	wd, _ := os.Getwd()
	r = func() error {
		_, err := c.Chdir(wd)
		return err
	}

	err = os.Chdir(dir)
	if err != nil {
		err = fmt.Errorf("chdir failed with: %s", err)
		return
	}

	return
}

func (c *Context) RemoveIfExist(fileName string) (removed bool) {
	_, err := os.Stat(fileName)
	if err == nil {
		err = c.Execute("rm", fileName)
		if err == nil {
			return true
		}
	}

	return false
}
