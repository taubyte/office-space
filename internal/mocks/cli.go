package mocks

import (
	"context"
	"os"
	"path"

	"github.com/taubyte/office-space/cli"
	"github.com/taubyte/office-space/common"
	"github.com/taubyte/office-space/env"
	. "github.com/taubyte/office-space/singletons"
)

const TestDirectory = "./_test_env"

func ResetTestDir() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	testDir := path.Join(cwd, TestDirectory)
	os.Mkdir(testDir, 0777)
	env.Set().WorkspaceDirectory(testDir)
	return testDir, nil
}

func CLI() (*mockCLI, error) {
	testDir, err := ResetTestDir()
	if err != nil {
		return nil, err
	}

	app, err := cli.Initialize(context.Background())
	if err != nil {
		return nil, err
	}

	return &mockCLI{
		Dir:     testDir,
		Runtime: Runtime(),
		app:     app,
	}, nil
}

type mockCLI struct {
	common.Runtime

	Dir string

	app *cli.App
}

func (m *mockCLI) Close() {
	os.RemoveAll(m.Dir)
}

func (*mockCLI) Run(args ...string) error {
	os.Args = append(os.Args[0:1], args...)

	// Do not open with code
	os.Args = append(os.Args, "--nocode")

	cli, err := cli.Initialize(context.Background())
	if err != nil {
		return err
	}

	return cli.Run(os.Args)
}
