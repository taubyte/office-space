package air

import (
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
	config := &airConfig{}
	err = toml.Unmarshal(EmptyAirToml, config)
	if err != nil {
		return
	}

	argsSlice := ctx.Args().Slice()
	args := strings.Split(config.Build.Cmd, " ")
	if len(argsSlice) > 0 {
		if strings.HasPrefix(argsSlice[0], "-") == true {
			args = append(args, argsSlice...)
		} else {
			args = append(args, append([]string{"--run"}, argsSlice...)...)
		}
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
