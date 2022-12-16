package work

import (
	"fmt"

	"github.com/taubyte/office-space/common"
	"github.com/taubyte/office-space/runtime"
	. "github.com/taubyte/office-space/singletons"
)

func build(ctx *runtime.Context) error {
	work, err := GoWork().Read()
	if err != nil {
		return fmt.Errorf("Reading GoWork failed with: %s", err)
	}

	err = Workspace().Edit(func(ws *common.VsWorkspace) error {
		ws.Folders = make([]common.VsFolder, len(work.Use))

		for idx, using := range work.Use {
			ctx.DisplayExec(fmt.Sprintf("%s < `%s`", Workspace().FileName(), using.Path))

			ws.Folders[idx] = common.VsFolder{
				Path: using.Path,
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("Editing workspace failed with: %s", err)
	}

	return Workspace().OpenWithCode()
}
