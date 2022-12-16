package singletons

import "github.com/taubyte/office-space/common"

var SingletonRuntime common.Runtime

func Runtime() common.Runtime {
	if SingletonRuntime == nil {
		panic("Runtime is nil")
	}

	return SingletonRuntime
}
