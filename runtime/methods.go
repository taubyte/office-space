package runtime

import (
	"github.com/urfave/cli/v2"
)

func (ctx *Context) SetContext(_ctx *cli.Context) {
	ctx.Context = _ctx
}
