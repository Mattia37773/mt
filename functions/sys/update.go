/*
Copyright © 2026 Matze
*/
package sys

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"runtime/debug"
	"strings"

	"github.com/mattia37773/mt/functions/ui"
)

func GetNewestCliVersionFunction(repoUrl string, currentVersion string) string {
	response, err := http.Get(repoUrl + "/releases/latest")

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

func GetCurrentVersion() string {

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "dev"
	}
	v := info.Main.Version
	v = strings.Split(v, "+")[0]

	return v
}

func ShowUpdateMessage(version string, repoUrl string) {
	latestVersion := GetNewestCliVersionFunction(repoUrl, version)
	if version < latestVersion {
		if version != "dev" {
			lines := []string{
				"A new update is available",
				"Current Version: " + version,
				"Latest  Version: " + latestVersion,
			}
			ui.Border(lines)
		}
	}
}
