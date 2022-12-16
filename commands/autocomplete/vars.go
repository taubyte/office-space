package autocomplete

import (
	_ "embed"
)

//go:embed bash_autocomplete.sh
var script string
