package go_mod

import (
	"golang.org/x/mod/modfile"
)

type goMod struct{}

type goModFile struct {
	path string
	mod  *modfile.File
}
