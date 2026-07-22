/*
Copyright © 2026 Matze
*/
// add tests
package config

import (
	"runtime/debug"
	"strings"
)

var AppConfig = struct {
	Version       string
	BuildMethod   string
	ModulePath    string
	GithubUrl     string
	GithubBaseApi string
	Logo          string
}{
	Version:       getCurrentVersion(), //  TODO get versin somewhgere sys.GetCurrentVersion()
	BuildMethod:   "source",            // TODO override during buildtome
	ModulePath:    "github.com/mattia37773/mt",
	GithubUrl:     "https://github.com/mattia37773/mt",
	GithubBaseApi: "https://api.github.com/repos/mattia37773/mt",
	Logo: `
███╗   ███╗ ████████╗
████╗ ████║ ╚══██╔══╝
██╔████╔██║    ██║
██║╚██╔╝██║    ██║
██║ ╚═╝ ██║    ██║
╚═╝     ╚═╝    ╚═╝
`,
}

func getCurrentVersion() string {

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "dev"
	}
	v := info.Main.Version
	v = strings.Split(v, "+")[0]

	return v
}
