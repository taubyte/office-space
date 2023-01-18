package git_config

import "net/url"

func (p *gitConfigParser) Remote() (*url.URL, error) {
	remoteUrl, err := p.parser.Get(`remote "origin"`, "url")
	if err != nil {
		return nil, err
	}

	return url.Parse(remoteUrl)
}
