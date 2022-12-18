package common

import (
	"github.com/bigkevmcd/go-configparser"
)

type GitConfig interface {
	Open(string) (*configparser.ConfigParser, error)
}
