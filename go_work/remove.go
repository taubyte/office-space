package go_work

import (
	. "github.com/taubyte/office-space/singletons"
	"golang.org/x/mod/modfile"
)

func (w *goWork) RemoveUse(dir string) error {
	// Confirm the directory has a valid go.mod
	_, err := GoMod().Open(dir)
	if err != nil {
		return err
	}

	// Remove if it already exists
	err = w.Edit(func(wf *modfile.WorkFile) error {
		relativePath, err := w.RelativeTo(dir)
		if err != nil {
			return err
		}

		var found bool
		for _, use := range wf.Use {
			if use.Path == relativePath {
				found = true
			}
		}

		if found == true {
			Runtime().DisplayExec("go", "work", "edit", "-dropreplace="+dir)

			return wf.DropUse(relativePath)
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
