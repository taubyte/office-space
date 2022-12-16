package work

import (
	"github.com/pterm/pterm"
	"github.com/taubyte/office-space/runtime"
	. "github.com/taubyte/office-space/singletons"
)

func delete(ctx *runtime.Context) error {
	GoWork().Remove()

	pterm.Success.Println("go.work is gone!")
	return nil
}
