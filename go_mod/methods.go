package go_mod

import (
	"fmt"

	"github.com/taubyte/office-space/common"
	. "github.com/taubyte/office-space/singletons"
	"golang.org/x/mod/modfile"
)

func (w *goMod) Open(dir string) (common.GoModFile, error) {
	f := &goModFile{
		dir: dir,
	}

	err := f.Initialize()
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (w *goModFile) DropReplace(packageName string, packageVersion string) error {
	Runtime().DisplayExec(w.File().Module.Mod.String(), "drop", "replace", packageName)

	err := w.File().DropReplace(packageName, packageVersion)
	if err != nil {
		return fmt.Errorf("dropping replace failed with: %s", err)
	}

	return Runtime().WriteFile(w.dir, modfile.Format(w.File().Syntax), 0777)
}
