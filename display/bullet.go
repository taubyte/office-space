package display

import (
	"os"
	"path"
	"path/filepath"

	"github.com/pterm/pterm"
)

func (d *Displayer) WroteDirectoriesInFile(wrote []string, file string, dir string) {
	if d.verbose == false {
		return
	}

	bullets := make([]pterm.BulletListItem, len(wrote))

	for idx, folder := range wrote {
		bullets[idx] = pterm.BulletListItem{
			Text:        folder,
			TextStyle:   pterm.NewStyle(pterm.FgYellow),
			BulletStyle: pterm.NewStyle(pterm.FgYellow),
			Bullet:      ">",
		}
	}

	var relativePath string
	cwd, err := os.Getwd()
	if err != nil {
		return
	} else {
		filePath := path.Join(dir, file)

		relativePath, err = filepath.Rel(cwd, filePath)
		if err != nil {
			relativePath = filePath
		}
	}

	filename := pterm.LightMagenta(relativePath)

	if len(bullets) > 0 {
		// Ignoring error as it is always nil
		bulletList, _ := pterm.DefaultBulletList.WithItems(bullets).Srender()

		pterm.Success.Printfln("Wrote relative directories to %s: \n%s", filename, bulletList)
	} else {
		pterm.Success.Printfln("Wrote empty %s", relativePath)
	}
}
