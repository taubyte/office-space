package workspace

import (
	"sync"

	"github.com/taubyte/office-space/common"
)

type workspace struct {
	fileLock sync.RWMutex
}

type runWith struct {
	ws      *workspace
	handler common.DirHandler
}

type dirsConfig struct {
	relative bool
	all      bool
}
