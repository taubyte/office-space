package common

import "golang.org/x/mod/modfile"

type GoMod interface {
	Open(dir string) (GoModFile, error)
}

type GoModFile interface {
	File() *modfile.File
	DropReplace(packageName, packageVersion string) error
}
