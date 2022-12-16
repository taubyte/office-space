package common

type RunWith interface {
	/* Runs the command with dir being the absolute path for every folder defined in a workspace */
	Absolute() error

	/* Runs the command with dir being the relative path for every folder defined in a workspace */
	Relative() error

	/* Runs the command with dir being the absolute path for every folder defined in the workspace's directory */
	Root() error
}

type DirHandler func(dir string) error
type WorkSpaceHandler func(vs *VsWorkspace) error

type Workspace interface {
	Stat() error
	Delete() error

	/* Defines a runWith for running commands asynchronously */
	ForEach(DirHandler) RunWith
	Loc() string
	Dir() string
	FileName() string

	Read() (VsWorkspace, error)
	Write(vs VsWorkspace) error
	Edit(WorkSpaceHandler) error

	New() error
	OpenWithCode() error

	AddUse(dir string) error
	RemoveUse(dir string) error

	List() ([]string, error)
	RelativeTo(dir string) (string, error)
}
