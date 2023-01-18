package work

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/pterm/pterm"
	"github.com/taubyte/office-space/runtime"
	. "github.com/taubyte/office-space/singletons"
	"github.com/urfave/cli/v2"
)

func updateCommands() []*cli.Command {
	files, err := Workspace().List()
	if err != nil {
		return nil
	}

	commands := []*cli.Command{}
	for _, file := range files {
		basePath := filepath.Base(file)

		commands = append(commands, &cli.Command{
			Name: basePath,
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "no-git",
					Usage: "Does not pull or checkout master of the package being updated",
				},
			},
			Action: runtime.Wrap(update(file)),
		})
	}

	return commands
}

func update(filePath string) runtime.ActionFunc {
	return func(ctx *runtime.Context) error {
		absPath := path.Join(Workspace().Dir(), filePath)

		fileMod, err := GoMod().Open(absPath)
		if err != nil {
			return fmt.Errorf("Opening go.mod of `%s` failed with: %s", absPath, err)
		}
		packageName := fileMod.File().Module.Mod.Path

		err = Workspace().ForEach(func(dir string) error {
			if dir == absPath {
				return nil
			} else {
				mod, err := GoMod().Open(dir)
				if err != nil {
					pterm.Info.Printfln("Dir: `%s` is not a go package", dir)

					return nil
				}
				ctx.DisplayExec(mod.File().Module.Mod.String(), "drop", "replace", packageName)

				err = mod.DropReplace(packageName, "")
				if err != nil {
					return fmt.Errorf("drop replace of %s in %s failed with: %s", packageName, dir, err)
				}

				return nil
			}
		}).Absolute()
		if err != nil {
			return err
		}
		err = GoWork().RemoveUse(absPath)
		if err != nil {
			return err
		}

		err = Workspace().RemoveUse(absPath)
		if err != nil {
			return err
		}

		if ctx.Bool("no-git") == false {
			// TODO checkout HEAD?
			err = ctx.ExecuteInDir(absPath, "git", "checkout", "main")
			if err != nil {
				err = ctx.ExecuteInDir(absPath, "git", "checkout", "master")
				if err != nil {
					return err
				}
			}

			err = ctx.ExecuteInDir(absPath, "git", "pull")
			if err != nil {
				return err
			}
		}

		err = Workspace().ForEach(func(dir string) error {
			mod, err := GoMod().Open(dir)
			if err != nil {
				pterm.Info.Printfln("Dir: `%s` is not a go package", dir)

				return nil
			}
			err = mod.DropReplace(packageName, "")
			if err != nil {
				return fmt.Errorf("dropping replace of %s in %s failed with: %s", packageName, dir, err)
			}

			err = ctx.ExecuteInDir(dir, "go", "get", packageName)
			if err != nil {
				return fmt.Errorf("go get %s failed with: %s", packageName, err)
			}

			err = ctx.ExecuteInDir(dir, "go", "mod", "tidy")
			if err != nil {
				return fmt.Errorf("tidy %s failed with: %s", dir, err)
			}

			return nil
		}).Absolute()
		if err != nil {
			return err
		}

		return Work(ctx, false)
	}
}
