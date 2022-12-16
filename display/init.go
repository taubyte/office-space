package display

import "github.com/taubyte/office-space/common"

func Initialize() (d common.Displayer, err error) {
	d = &Displayer{
		verbose: true,
	}
	return
}
