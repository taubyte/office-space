package cli

import (
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

// slices "bitbucket.org/taubyte/utils/slices/string"
func slicesContains(slice []string, value string) bool {
	for _, s := range slice {
		if value == s {
			return true
		}
	}
	return false
}

// Check default to have run be the default command
// Ex, `go run . git`
// Returns arguments [run git]
// `go run . issue`
// Returns [issue]
func structureDefaultCommand() bool {
	args := os.Args

	var defaultCommand bool
	if len(args) > 1 {
		for _, cmd := range validCommands {
			if args[1] == cmd {
				return true
			}
		}
	}

	// Append run if it is not a default or help command
	if IsHelpCommand() == false {
		os.Args = append([]string{args[0], "run"}, args[1:]...)
	}

	return defaultCommand
}

// https://github.com/urfave/cli/issues/427#issuecomment-712292636
func movePostfixOptions(args []string, validFlags []string) []string {
	var idx = 1
	the_args := make([]string, 0)
	for {
		if idx >= len(args) {
			break
		}

		if args[idx][0] == '-' && slicesContains(validFlags, args[idx]) {
			if !strings.Contains(args[idx], "=") {
				idx++
			}
		} else {
			// add to args accumulator
			the_args = append(the_args, args[idx])

			// remove from real args list
			new_args := make([]string, 0)
			new_args = append(new_args, args[:idx]...)
			new_args = append(new_args, args[idx+1:]...)
			args = new_args
			idx--
		}

		idx++
	}

	// append extracted arguments to the real args
	return append(args, the_args...)
}

/*
Gets a string slice from cli.flags so that they can be moved to their required
position from anywhere.
This way `asd work d --dry` and `asd --dry work d` are both valid commands.
*/
func getValidFlags(flags []cli.Flag) []string {
	valid := make([]string, 0)
	for _, flag := range flags {
		names := make([]string, 0)
		for _, name := range append(flag.Names()) {
			names = append(names, []string{
				"-" + name,
				"--" + name,
			}...)
		}
		valid = append(valid, names...)
	}

	return append(valid)
}
