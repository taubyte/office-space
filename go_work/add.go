package go_work

import (
	"fmt"

	"golang.org/x/mod/modfile"

	. "github.com/taubyte/office-space/singletons"
)

func (w *goWork) AddUse(dir string) error {
	// Confirm the directory has a valid go.mod
	modFile, err := GoMod().Open(dir)
	if err != nil {
		return err
	}

	err = w.Edit(func(wf *modfile.WorkFile) error {
		for _, use := range wf.Use {
			if use.Path == dir {
				return fmt.Errorf("Already using `%s`", dir)
			}
		}

		relativePath, err := w.RelativeTo(dir)
		if err != nil {
			return err
		}

		Runtime().DisplayExec("go", "work", "use", relativePath)

		if Runtime().Dry() {
			return nil
		}

		return wf.AddUse(relativePath, modFile.File().Syntax.Name)
	})
	if err != nil {
		return err
	}

	return nil
}
