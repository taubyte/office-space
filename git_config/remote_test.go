package git_config

import (
	_ "embed"
	"testing"

	"github.com/taubyte/office-space/internal/mocks"
)

func TestRemote(t *testing.T) {
	
	// Initialize
	c := &gitConfig{}

	// Mocks CLI
	ctx, err := mocks.CLI()
	if err != nil {
		t.Error(err)
		return
	}

	defer ctx.Close() // Comment this line out to see generated file structure

	// Create a fake workspace with some repos
	err = ctx.FakeWorkspace("repo1", "repo2", "repo3", "repo4")
	if err != nil {
		t.Error(err)
		return
	}

	// Create the branches
	branch1 := "TP-1_some_1branch_stuff"
	branch2 := "TP-2_some_2branch_stuff"
	branch3 := "TP-3_some_3branch_stuff"
	branch4 := "TP-4_some_4branch_stuff"

	// Create some test data
	testData := map[string][]string{
		"repo1": {branch1, branch2, branch3},
		"repo2": {branch1, branch2, branch3},
		"repo3": {branch1, branch2, branch3},
		"repo4": {branch1, branch4},
	}

	// Create a map to hold the paths for the repos
	paths := map[string]string{}

	// Go through the test data
	for name, branches := range testData { 
		// Create a fake module with the branches
		paths[name], err = ctx.FakeModuleWithBranches(name, branches...)
		if err != nil {
			return
		}

		// Add the remote for the new branches
		err = ctx.ExecuteInDir(paths[name], "git", "remote", "add", "origin", "https://github.com/fake/fake.git")
		if err != nil {
			return
		}
	}

	// Iterate through all the paths
	for _,path := range paths {
		// Get a parser by opening the config file for the current branch path
		parser, err := c.Open(path + "/.git/config")
		if err != nil {
			t.Error(err)
			return
		}

		// Grab the remote url from the .git/config
		url, err := parser.Remote()
		if err != nil {
			t.Error(err)
			return
		}

		// Check if it matches our junk data
		if url.String() != "https://github.com/fake/fake.git" {	
			// If it doesn't then show an error for this test
			t.Error("Git Configuration Test : Expected 'https://github.com/fake/fake.git', Got: ",url)
			return
		}
	}
}
