package singletons

import (
	"github.com/taubyte/office-space/common"
)

var SingletonGitConfig common.GitConfig

//One singleton for the git config
func GitConfig() common.GitConfig {
	if SingletonGitConfig == nil {
		panic("GitConfig is nil")
	}
	return SingletonGitConfig
}
