package work

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
	"github.com/taubyte/office-space/runtime"
	. "github.com/taubyte/office-space/singletons"
)

func Work(ctx *runtime.Context, goWorkOff bool) error {
	if goWorkOff == true {
		err := os.Setenv("GOWORK", "off")
		if err != nil {
			pterm.Error.Printfln("Setting [`GOWORK`] failed with: %s", err)
		}
	}

	err := clean(ctx)
	if err != nil {
		return fmt.Errorf("clean failed with: %s", err)
	}

	Display().SetVerbose(false)
	defer Display().SetVerbose(true)

	err = Workspace().ForEach(func(dir string) error {
		_, err := GoMod().Open(dir)
		if err != nil {
			pterm.Info.Printfln("Dir: `%s` is not a go package", dir)

			return nil
		}

		return GoWork().AddUse(dir)
	}).Absolute()
	if err != nil {
		return err
	}

	if goWorkOff == true {
		err = os.Unsetenv("GOWORK")
		if err != nil {
			pterm.Error.Printfln("Unsetting [`GOWORK`] failed with: %s", err)
		}
	} else {
		// Only open with code if doing a normal Work,  let other side handle as it's being called from a method
		return Workspace().OpenWithCode()
	}

	return nil
}
