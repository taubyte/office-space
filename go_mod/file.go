package go_mod

import (
	"fmt"
	"os"
	"path"

	"golang.org/x/mod/modfile"
)

func (f *goModFile) Initialize() error {
	data, err := os.ReadFile(path.Join(f.dir, GoModName))
	if err != nil {
		return fmt.Errorf("Reading `%s` failed with: %s", f.dir, err)
	}

	f.mod, err = modfile.Parse(GoModName, data, nil)
	if err != nil {
		return fmt.Errorf("Parsing `%s` failed with: %s", f.dir, err)
	}

	return nil
}

func (f *goModFile) File() *modfile.File {
	return f.mod
}
