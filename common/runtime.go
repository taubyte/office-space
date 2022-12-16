package common

import (
	"io/fs"

	"github.com/urfave/cli/v2"
)

// Sent with a method that changes current directory, call to return to the
// previous directory, equivalent to `cd -`
type ReturnToDirMethod func() error

type Runtime interface {
	Execute(name string, args ...string) error
	DisplayExec(name string, args ...string)
	ExecuteInDir(dir, name string, args ...string) error
	ExecuteCapture(name string, args ...string) (out, errOut string, err error)
	ExecuteCaptureInDir(dir, name string, args ...string) (out, errOut string, err error)
	WriteFile(name string, data []byte, perm fs.FileMode) error
	RemoveAll(path string) error
	Remove(path string) error
	Chdir(dir string) (r ReturnToDirMethod, err error)
	RemoveIfExist(fileName string) (removed bool)

	// Exported CLI commands
	Bool(string) bool

	Dry() bool

	SetContext(ctx *cli.Context)
}
