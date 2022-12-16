package env

func Set() setter {
	return setter{}
}

func (setter) WorkspaceName(name string) {
	engine.set(WorkspaceNameEnvKey, name)
}

func (setter) WorkspaceExt(ext string) {
	engine.set(WorkspaceExtEnvKey, ext)
}

func (setter) WorkspaceDirectory(dir string) {
	engine.set(WorkspaceDirectoryEnvKey, dir)
}

func (setter) GitPrefix(prefix string) {
	engine.set(GitPrefixEnvKey, prefix)
}
