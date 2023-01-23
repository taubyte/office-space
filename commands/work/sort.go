package work

import (
	"github.com/taubyte/office-space/common"
	"github.com/taubyte/office-space/runtime"
	. "github.com/taubyte/office-space/singletons"
	"golang.org/x/exp/slices"
)

func sort(ctx *runtime.Context) error {
	err := Workspace().Edit(func(vs *common.VsWorkspace) error {
		folders := make([]string, len(vs.Folders))
		for idx, folder := range vs.Folders {
			folders[idx] = folder.Path
		}

		slices.Sort(folders)

		vs.Folders = make([]common.VsFolder, len(folders))
		for idx, folder := range folders {
			vs.Folders[idx] = common.VsFolder{
				Path: folder,
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return Workspace().OpenWithCode()
}
