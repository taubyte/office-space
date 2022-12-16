package runtime

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

/*
Wrap takes 1-2 handlers, [0] is the generic command and [1] is the dry command
if the dry command is not provided it simply calls the command
*/
func Wrap(handlers ...ActionFunc) cli.ActionFunc {
	var (
		handler ActionFunc
		dry     ActionFunc
	)
	switch len(handlers) {
	case 1:
		handler = handlers[0]
		dry = handlers[0]
	case 2:
		handler = handlers[0]
		dry = handlers[1]
	default:
		panic(fmt.Sprintf("Unexpected handler length: %d", len(handlers)))
	}

	return func(ctx *cli.Context) error {
		_runtime.Context = ctx

		if ctx.Bool("dry") == true {
			return dry(_runtime)
		}

		return handler(_runtime)
	}
}
