package issue

import (
	"fmt"

	"github.com/taubyte/office-space/runtime"
)

func getBranchPrefix(ctx *runtime.Context) (branchPrefix string, err error) {
	branchPrefix = ctx.Args().First()
	if len(branchPrefix) == 0 {
		err = fmt.Errorf("Must provide a branch prefix, ex: `asd issue <prefix>`")
	}

	return
}
