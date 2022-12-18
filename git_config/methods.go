package git_config

import (
	"errors"

	"github.com/bigkevmcd/go-configparser"
	"github.com/taubyte/office-space/common"
)

/*
Open method will take path string that's passed in and initialize and return a configuration parser instance or an error
*/
func (w *gitConfig) Open(path string) (common.GitConfigParser, error) {
	if len(path) != 0 { // If a path was provided as expected
		// Attempt to create a new configuration file parser
		f, err := configparser.NewConfigParserFromFile(path)
		if err != nil { // If we encounter an error
			return nil, err // Return nil and the error
		}

		return &gitConfigParser{f}, nil // Otherwise return the configparser and nil for the error
	} else { // Otherwise
		// Return nil for the parser and an error explaining the issue with the path
		return nil, errors.New("You must include a valid path when calling Open() method on a Github Configuration instance.")
	}
}
