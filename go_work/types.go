package go_work

import (
	"path"
	"sync"

	"github.com/taubyte/office-space/common"
	"github.com/taubyte/office-space/env"
)

var _ common.GoWork = &goWork{}

type goWork struct {
	workFileLock sync.RWMutex
}

func Initialize() (w common.GoWork, err error) {
	w = &goWork{}

	return
}

func (w *goWork) Dir() string {
	return env.Get().WorkspaceDirectory()
}

func (w *goWork) Loc() string {
	return path.Join(w.Dir(), "go.work")
}
