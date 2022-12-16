package runtime

import (
	"errors"

	"github.com/taubyte/office-space/common"
	"github.com/taubyte/office-space/singletons"
)

var _runtime *Context

func Initialize() error {
	_runtime = &Context{nil}
	singletons.SingletonRuntime = _runtime

	return nil
}

func Get() (common.Runtime, error) {
	if _runtime == nil {
		return nil, errors.New("Runtime not initialized")
	}

	return _runtime, nil
}
