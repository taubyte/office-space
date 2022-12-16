package singletons

import (
	"github.com/taubyte/office-space/common"
)

var SingletonGoMod common.GoMod

func GoMod() common.GoMod {
	if SingletonGoMod == nil {
		panic("GoMod is nil")
	}

	return SingletonGoMod
}
