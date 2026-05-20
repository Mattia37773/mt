/*
Copyright © 2026 Matze
*/
package config

import (
	"github.com/mattia37773/mt/functions/sys"
)

// Current Cli Version
var Version string = sys.GetCurrentVersion()

// Defines the Method the Sourcecode got compiled
// depending on the Result the update message shows a different way to update
// and the update command works differently
var BuildMethod string = "source"

// Github Repo
var ModulePath string = "github.com/mattia37773/mt"
var GithubUrl string = "https://" + ModulePath
var GithubBaseApi string = "https://api.github.com/repos/mattia37773/mt"

func GetNewestVersion() string {
	return sys.GetNewestCliVersionFunction(GithubBaseApi, Version)
}

var Logo string = `
███╗   ███╗ ████████╗
████╗ ████║ ╚══██╔══╝
██╔████╔██║    ██║
██║╚██╔╝██║    ██║
██║ ╚═╝ ██║    ██║
╚═╝     ╚═╝    ╚═╝
`
