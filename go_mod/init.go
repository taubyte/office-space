package go_mod

import (
	"github.com/taubyte/office-space/common"
)

func Initialize() (m common.GoMod, err error) {
	m = &goMod{}
	return
}
