package singletons

import (
	"github.com/taubyte/office-space/common"
)

var SingletonWorkspace common.Workspace

func Workspace() common.Workspace {
	if SingletonWorkspace == nil {
		panic("Workspace is nil")
	}

	return SingletonWorkspace
}
