package display

type Displayer struct {
	verbose bool
}

func (d *Displayer) SetVerbose(v bool) {
	d.verbose = v
}
