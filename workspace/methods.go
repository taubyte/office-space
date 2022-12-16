package workspace

import (
	"os"
	"path"

	"github.com/taubyte/office-space/env"
	. "github.com/taubyte/office-space/singletons"
)

func (ws *workspace) Stat() (err error) {
	_, err = os.Stat(ws.Loc())

	return
}

func (ws *workspace) Delete() error {
	return Runtime().Remove(ws.Loc())
}

func (ws *workspace) Loc() string {
	return path.Join(ws.Dir(), ws.FileName())
}

func (w *workspace) directories(config dirsConfig) (dirs []string, err error) {
	if config.all {
		return w.alldirs(config)
	}

	ws, err := w.Read()
	if err != nil {
		return nil, err
	}

	dirs = make([]string, 0)
	for _, folder := range ws.Folders {
		if config.relative == true {
			dirs = append(dirs, folder.Path)
		} else {
			dirs = append(dirs, path.Join(w.Dir(), folder.Path))
		}
	}

	return
}

func (w *workspace) relativeDirs() (dirs []string, err error) {
	return w.dirs(dirsConfig{relative: true})
}

func (w *workspace) alldirs(config dirsConfig) ([]string, error) {
	dirs, err := os.ReadDir(w.Dir())
	if err != nil {
		return nil, err
	}

	_dirs := make([]string, 0)
	for _, dir := range dirs {
		if dir.IsDir() {
			if config.relative == true {
				_dirs = append(_dirs, dir.Name())
			} else {
				_dirs = append(_dirs, path.Join(w.Dir(), dir.Name()))
			}
		}
	}

	return _dirs, nil
}

func (w *workspace) dirs(config dirsConfig) (dirs []string, err error) {
	if config.all {
		return w.alldirs(config)
	}

	ws, err := w.Read()
	if err != nil {
		return nil, err
	}

	dirs = make([]string, 0)
	for _, folder := range ws.Folders {
		if config.relative == true {
			dirs = append(dirs, folder.Path)
		} else {
			dirs = append(dirs, path.Join(w.Dir(), folder.Path))
		}

	}

	return
}

func (w *workspace) OpenWithCode() error {
	if Runtime().Bool("nocode") == true {
		return nil
	}

	return Runtime().Execute("code", w.Loc())
}

func (ws *workspace) FileName() string {
	return ws.Name() + ws.Ext()
}

func (ws *workspace) Dir() string {
	return env.Get().WorkspaceDirectory()
}

func (ws *workspace) Name() string {
	return env.Get().WorkspaceName()
}

func (ws *workspace) Ext() string {
	return env.Get().WorkspaceExt()
}
