package go_mod

import (
	"github.com/taubyte/office-space/common"
	"golang.org/x/mod/modfile"
)

var _ common.GoMod = &goMod{}

type goMod struct{}

var _ common.GoModFile = &goModFile{}

type goModFile struct {
	dir string
	mod *modfile.File
}
