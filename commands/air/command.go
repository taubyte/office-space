package air

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/taubyte/office-space/runtime"
	"github.com/urfave/cli/v2"
)

// TODO embed air itself as a library
// Attempted but failed due to air on import effecting os.Args
func Command() *cli.Command {
	return &cli.Command{
		Name:   "air",
		Usage:  "Starts and removes on exit a golang air context for testing",
		Action: runtime.Wrap(command),
	}
}

func command(ctx *runtime.Context) (err error) {
	var specificTest string
	osArgs := ctx.Args().Slice()
	if len(osArgs) > 0 {
		if len(osArgs) > 1 {
			return fmt.Errorf("Got args %v, only expected one, ex: `asd air <TestToRun>`", osArgs)
		}
		specificTest = osArgs[0]
	}

	config := &airConfig{}
	err = toml.Unmarshal(EmptyAirToml, config)
	if err != nil {
		return
	}

	args := strings.Split(config.Build.Cmd, " ")
	if len(specificTest) > 0 {
		args = append(args, []string{"--run", specificTest}...)
	}

	config.Build.Cmd = strings.Join(args, " ")
	newTomlBytes, err := toml.Marshal(config)
	if err != nil {
		return
	}

	err = ctx.WriteFile(".air.toml", newTomlBytes, 0644)
	if err != nil {
		return
	}

	// Capture Ctrl+c to clear .air.toml and tmp directory
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cleanUp(ctx)
	}()

	return ctx.Execute("air")
}

func cleanUp(ctx *runtime.Context) {
	err := ctx.RemoveAll("tmp")
	if err != nil {
		log.Fatalf("Failed removing tmp: %v", err)
	}

	err = ctx.Remove(".air.toml")
	if err != nil {
		log.Fatalf("Failed removing .air.toml: %v", err)
	}
}
