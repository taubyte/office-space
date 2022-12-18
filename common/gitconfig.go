package common

import "net/url"

type GitConfig interface {
	Open(string) (GitConfigParser, error)
}

type GitConfigParser interface {
	Remote() (*url.URL, error)
}
