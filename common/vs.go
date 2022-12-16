package common

//	{
//		"folders": [
//				{
//						"path": "dreamland-cli"
//				}
//		],
//		"settings": {}
//	}
type VsWorkspace struct {
	Folders  []VsFolder             `json:"folders"`
	Settings map[string]interface{} `json:"settings"`
}

func (ws VsWorkspace) List() []string {
	if len(ws.Folders) == 0 {
		return []string{}
	}

	files := make([]string, len(ws.Folders))
	for idx, folder := range ws.Folders {
		files[idx] = folder.Path
	}

	return files
}

type VsFolder struct {
	Path string `json:"path"`
}
