package air

import (
	"regexp"
	"time"
)

// Copied from https://github.com/cosmtrek/air/runner/config.go
type airConfig struct {
	Root        string    `toml:"root"`
	TmpDir      string    `toml:"tmp_dir"`
	TestDataDir string    `toml:"testdata_dir"`
	Build       cfgBuild  `toml:"build"`
	Color       cfgColor  `toml:"color"`
	Log         cfgLog    `toml:"log"`
	Misc        cfgMisc   `toml:"misc"`
	Screen      cfgScreen `toml:"screen"`
}

type cfgBuild struct {
	Cmd              string        `toml:"cmd"`
	Bin              string        `toml:"bin"`
	FullBin          string        `toml:"full_bin"`
	ArgsBin          []string      `toml:"args_bin"`
	Log              string        `toml:"log"`
	IncludeExt       []string      `toml:"include_ext"`
	ExcludeDir       []string      `toml:"exclude_dir"`
	IncludeDir       []string      `toml:"include_dir"`
	ExcludeFile      []string      `toml:"exclude_file"`
	ExcludeRegex     []string      `toml:"exclude_regex"`
	ExcludeUnchanged bool          `toml:"exclude_unchanged"`
	FollowSymlink    bool          `toml:"follow_symlink"`
	Delay            int           `toml:"delay"`
	StopOnError      bool          `toml:"stop_on_error"`
	SendInterrupt    bool          `toml:"send_interrupt"`
	KillDelay        time.Duration `toml:"kill_delay"`
	regexCompiled    []*regexp.Regexp
}

type cfgLog struct {
	AddTime bool `toml:"time"`
}

type cfgColor struct {
	Main    string `toml:"main"`
	Watcher string `toml:"watcher"`
	Build   string `toml:"build"`
	Runner  string `toml:"runner"`
	App     string `toml:"app"`
}

type cfgMisc struct {
	CleanOnExit bool `toml:"clean_on_exit"`
}

type cfgScreen struct {
	ClearOnRebuild bool `toml:"clear_on_rebuild"`
}
