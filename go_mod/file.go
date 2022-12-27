package go_mod

import (
	"fmt"
	"os"

	"golang.org/x/mod/modfile"
)

func (f *goModFile) Initialize() error {
	data, err := os.ReadFile(f.path)
	if err != nil {
		return fmt.Errorf("Reading `%s` failed with: %s", f.path, err)
	}

	f.mod, err = modfile.Parse(GoModName, data, nil)
	if err != nil {
		return fmt.Errorf("Parsing `%s` failed with: %s", f.path, err)
	}

	return nil
}

func (f *goModFile) File() *modfile.File {
	return f.mod
}
