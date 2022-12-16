package singletons

import (
	"github.com/taubyte/office-space/common"
)

var SingletonGoWork common.GoWork

func GoWork() common.GoWork {
	if SingletonGoWork == nil {
		panic("GoWork is nil")
	}

	return SingletonGoWork
}
