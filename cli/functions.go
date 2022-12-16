package cli

import "os"

var helpOptions = []string{
	"help", "-help", "--help",
	"h", "-h", "--h",
}

var bashCompletionOptions = []string{
	"-generate-bash-completion", "--generate-bash-completion",
}

func IsHelpCommand() bool {
	for _, arg := range os.Args {
		if slicesContains(append(helpOptions, bashCompletionOptions...), arg) {
			return true
		}
	}

	return false
}
