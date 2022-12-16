package go_work

import (
	"os"

	"golang.org/x/mod/modfile"
)

func (w *goWork) read() (*modfile.WorkFile, error) {
	data, err := os.ReadFile(w.Loc())
	if err != nil {
		err = w.Init()
		if err != nil {
			return nil, err
		}

		data, err = os.ReadFile(w.Loc())
		if err != nil {
			return nil, err
		}
	}

	return modfile.ParseWork("go.work", data, nil)
}

func (w *goWork) Read() (*modfile.WorkFile, error) {
	w.workFileLock.RLock()
	defer w.workFileLock.RUnlock()

	return w.read()
}
