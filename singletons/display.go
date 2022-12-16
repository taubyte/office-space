package singletons

import (
	"github.com/taubyte/office-space/common"
)

var SingletonDisplay common.Displayer

func Display() common.Displayer {
	if SingletonDisplay == nil {
		panic("Display is nil")
	}

	return SingletonDisplay
}
