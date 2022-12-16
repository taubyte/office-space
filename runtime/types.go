package runtime

import (
	"github.com/urfave/cli/v2"
)

// A Function which takes the runtime context for use in generating cli commands
type ActionFunc func(ctx *Context) error

// The runtime context which handles the --dry boolean and displays "commands" being executed
type Context struct {
	*cli.Context
}
