package runtime

func (ctx *Context) Dry() bool {
	return ctx.Bool("dry")
}
