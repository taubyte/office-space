package env

func Get() getter {
	return getter{}
}

func (getter) WorkspaceName() string {
	return engine.get(WorkspaceNameEnvKey, DefaultWorkspaceName)
}

func (getter) WorkspaceExt() string {
	return engine.get(WorkspaceExtEnvKey, DefaultWorkspaceExt)
}

func (getter) WorkspaceDirectory() string {
	return engine.get(WorkspaceDirectoryEnvKey, "")
}

func (getter) GitPrefix() string {
	return engine.get(GitPrefixEnvKey, "")
}
