package git_config

import (
	_ "embed"
	"fmt"
	"testing"
)

func TestRemote(t *testing.T) {
	// Initialize

	c := &gitConfig{}

	parser, err := c.Open("test_git_config")
	if err != nil {
		t.Error(err)
		return
	}

	url, err := parser.Remote()
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("REMOTE: ", url)
}
