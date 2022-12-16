package autocomplete

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/taubyte/office-space/runtime"
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:   "autocomplete",
		Usage:  fmt.Sprintf("Used with eval or in .bashrc for autocompletion"),
		Action: runtime.Wrap(command),
	}
}

func command(ctx *runtime.Context) error {
	basePath := filepath.Base(os.Args[0])

	fmt.Println(script + basePath)

	return nil
}
