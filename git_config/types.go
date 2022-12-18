package git_config

import (
	"github.com/taubyte/office-space/common"
)

var _ common.GitConfig = &gitConfig{}

type gitConfig struct{
	path string
	repositoryformatversion int
	filemode 				bool
	bare 					bool
	logallrefupdates 		bool
	ignorecase				bool
	precomposeunicode		bool
	remotes					[]GithubRemote
	branches				[]GithubBranch
}

type GithubRemote struct {
	url string
	fetch string
}

type GithubBranch struct {
	remote string
	merge string
}

