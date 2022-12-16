package go_work

import (
	"os"

	. "github.com/taubyte/office-space/singletons"
	"golang.org/x/mod/modfile"
)

func (w *goWork) write(file *modfile.WorkFile) (err error) {
	file.Cleanup()

	if Runtime().Dry() == false {
		err = os.WriteFile(w.Loc(), modfile.Format(file.Syntax), 0777)
	}

	if err == nil {
		files := make([]string, len(file.Use))
		for idx, use := range file.Use {
			files[idx] = use.Path
		}

		Display().WroteDirectoriesInFile(files, "go.work", w.Dir())
	}

	return
}

func (w *goWork) Write(file *modfile.WorkFile) error {
	w.workFileLock.Lock()
	defer w.workFileLock.Unlock()

	return w.write(file)
}
