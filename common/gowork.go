package common

import "golang.org/x/mod/modfile"

type GoWork interface {
	Read() (*modfile.WorkFile, error)
	Write(file *modfile.WorkFile) error
	Edit(func(*modfile.WorkFile) error) error

	Remove()
	Clean() error
	Init() error

	AddUse(dir string) error
	RemoveUse(dir string) error

	RelativeTo(dir string) (string, error)
}
