/*
Copyright © 2026 Matze
*/
package config

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"runtime/debug"
)

// config that can be overridden by ldflags
var (
	Version                = getCurrentVersion()
	OverrideCurrentVersion = ""
	BuildMethod            = "source"
	LatestVersion          = getNewestCliVersionFunction(Version)
	Environment            = "dev"
)

// static config
var AppConfig = struct {
	BuildMethod   string
	ModulePath    string
	GithubUrl     string
	GithubBaseApi string
	Logo          string
}{
	BuildMethod:   BuildMethod,
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

func init() {
	getCurrentVersion()
}
func getCurrentVersion() string {

	if OverrideCurrentVersion != "" {
		return OverrideCurrentVersion
	}

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "dev"
	}

	return info.Main.Version
}

func getNewestCliVersionFunction(currentVersion string) string {
	response, err := http.Get(AppConfig.GithubBaseApi + "/releases/latest")

	// Needed for an error if the
	// device isn't connected with internet
	if err != nil {
		return currentVersion
	}

	dir := os.TempDir()

	// The github api has harsh rate limiting
	// it returns nil when that limit is hit
	// So nothing gets printed when the limit is reached
	if response == nil || response.StatusCode != 403 {
		var data map[string]interface{}
		json.NewDecoder(response.Body).Decode(&data)
		latestVersion, _ := data["tag_name"].(string)

		// Write the newest version to a file
		file, _ := os.Create(dir + "mt-version")
		defer file.Close()
		file.WriteString(latestVersion)

		return latestVersion
	} else {
		_, err := os.Stat(dir + "mt-version")
		if !errors.Is(err, os.ErrNotExist) {
			fileversion, _ := os.ReadFile(dir + "mt-version")
			return string(fileversion)
		}

		return currentVersion
	}
}
