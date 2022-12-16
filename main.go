package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/pterm/pterm"
	"github.com/taubyte/office-space/cli"
)

func init() {
	pterm.SetDefaultOutput(os.Stderr)
}

func main() {
	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	app, err := cli.Initialize(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = app.RunContext(ctx, os.Args)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "air" {
			// Wait for air cleanup
			spinnerInfo, err := pterm.DefaultSpinner.Start("Cleaning up air...")
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)

			spinnerInfo.RemoveWhenDone = true
			spinnerInfo.Stop()
		}
	}
}
