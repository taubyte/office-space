package runtime

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func (c *Context) Execute(name string, args ...string) error {
	c.DisplayExec(name, args...)

	if c.Bool("dry") == true {
		return nil
	}

	cmd := exec.Command(name, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("Start command failed with: %s", err)
	}

	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("Command %s, %v failed with: %s", name, args, err)
	}

	return nil
}

func (c *Context) DisplayExec(name string, args ...string) {
	commandStr := name
	for _, arg := range args {
		commandStr += " " + arg
	}

	fmt.Println("$", commandStr)
}

func (c *Context) ExecuteInDir(dir, name string, args ...string) error {
	c.DisplayExec("cd", dir)
	c.DisplayExec(name, args...)

	if c.Bool("dry") == true {
		return nil
	}

	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("Start command failed with: %s", err)
	}

	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("Command %s, %v failed with: %s", name, args, err)
	}

	return nil
}

// TODO return reader/writer
func (c *Context) ExecuteCapture(name string, args ...string) (out, errOut string, err error) {
	c.DisplayExec(name, args...)

	if c.Bool("dry") == true {
		return
	}

	cmd := exec.Command(name, args...)
	var _out bytes.Buffer
	var _errOut bytes.Buffer
	cmd.Stdout = &_out
	cmd.Stderr = &_errOut
	defer func() {
		out = _out.String()
		errOut = _errOut.String()
	}()

	err = cmd.Start()
	if err != nil {
		err = fmt.Errorf("Start command failed with: %s", err)
	}
	err = cmd.Wait()

	return
}

// TODO return reader/writer
func (c *Context) ExecuteCaptureInDir(dir, name string, args ...string) (out, errOut string, err error) {
	c.DisplayExec("cd", dir)
	c.DisplayExec(name, args...)

	if c.Bool("dry") == true {
		return
	}

	cmd := exec.Command(name, args...)
	cmd.Dir = dir

	var _out bytes.Buffer
	var _errOut bytes.Buffer
	cmd.Stdout = &_out
	cmd.Stderr = &_errOut
	defer func() {
		out = _out.String()
		errOut = _errOut.String()
	}()

	err = cmd.Start()
	if err != nil {
		err = fmt.Errorf("Start command failed with: %s", err)
		return
	}
	err = cmd.Wait()

	return
}
