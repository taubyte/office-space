package work

import (
	"github.com/pterm/pterm"
	"github.com/taubyte/office-space/runtime"
	. "github.com/taubyte/office-space/singletons"
)

func clean(ctx *runtime.Context) error {
	err := GoWork().Clean()
	if err != nil {
		return err
	}

	pterm.Success.Println("go.work is clean!")
	return nil
}
