package workspace

import "github.com/taubyte/office-space/common"

func (w *workspace) Edit(handler common.WorkSpaceHandler) error {
	w.fileLock.Lock()
	defer w.fileLock.Unlock()

	ws, err := w.read()
	if err != nil {
		return err
	}

	err = handler(&ws)
	if err != nil {
		return err
	}

	return w.write(ws)
}
