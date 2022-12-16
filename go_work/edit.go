package go_work

import "golang.org/x/mod/modfile"

func (w *goWork) Edit(handler func(*modfile.WorkFile) error) error {
	w.workFileLock.Lock()
	defer w.workFileLock.Unlock()

	workFile, err := w.read()
	if err != nil {
		return err
	}

	err = handler(workFile)
	if err != nil {
		return err
	}

	return w.write(workFile)
}
