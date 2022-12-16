package workspace

import "fmt"

func (ws *workspace) List() ([]string, error) {
	file, err := ws.Read()
	if err != nil {
		return nil, fmt.Errorf("Reading workspace failed with: %s", err)
	}

	return file.List(), nil
}
