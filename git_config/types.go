package git_config

import (
	"github.com/bigkevmcd/go-configparser"
	"github.com/taubyte/office-space/common"
)

var _ common.GitConfig = &gitConfig{}

type gitConfig struct {
}

var _ common.GitConfigParser = &gitConfigParser{}

type gitConfigParser struct {
	parser *configparser.ConfigParser
}
