/*
Copyright © 2026 Matze
*/
package single

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/functions/ui"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:                   "update",
	Short:                 "Updates the cli",
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		updateSingle()
	},
}

func init() {
	cmd.RootCmd.AddCommand(updateCmd)
}

func updateSingle() {
	if config.Version == config.GetNewestVersion() {
		fmt.Println(ui.GlowPink(config.Logo))
		fmt.Println(ui.Green("The Newest Version is already installed"))
		os.Exit(0)
	}

	switch buildmethod := config.BuildMethod; buildmethod {
	case "github":
		confirm, err := ui.Confirm("Do you want to pull the latest updates from GitHub?")
		if err != nil {
			fmt.Print(ui.Red("Error: %s\n"), err)
			os.Exit(1)
		}
		if confirm == true {
			githubReleasesUpdate()
		} else {
			fmt.Println("Continuing without updating.")
		}
	case "source":
		confirm, err := ui.Confirm("You compiled this project yourself.\nDo you want to pull the latest updates from GitHub?")
		if err != nil {
			fmt.Print(ui.Red("Error: %s\n"), err)
			os.Exit(1)
		}
		if confirm == true {
			githubReleasesUpdate()
		} else {
			fmt.Println("Continuing without updating.")
		}
	case "go":
		cmd := `go install -ldflags "-X ` + config.ModulePath + `/config.BuildMethod=go" ` + config.ModulePath + `@latest`

		fmt.Println(ui.Green("Installed via go"))
		fmt.Println("Please update with the follwing commands")
		fmt.Println("clear the package cache: go clean -modcache")
		fmt.Printf("install the update: \"%s\"\n", cmd)

	case "homebrew":
		fmt.Println(ui.Green("Installed via Homebrew."))
		fmt.Println("Please update with \"brew upgrade mt\".")
	default:
		fmt.Println("There is currently not an option for this build method.")
		fmt.Println("This is only possible by manually modifiy the source code")
	}
}

func githubReleasesUpdate() {
	switch userOs := runtime.GOOS; userOs {
	case "darwin":
		url := getUrl(".tar.gz")
		updateMacLinux(url)
	case "linux":
		url := getUrl(".tar.gz")
		updateMacLinux(url)
	case "windows":
		url := getUrl(".zip")
		updateWindows(url)
	default:
		fmt.Printf(ui.Red("Sorry but your operation system is not supported"))
		fmt.Printf("You can check the source code out and add support for your operatingsystem ;)")
		os.Exit(0)
	}
}

func getUrl(archive string) string {
	os := runtime.GOARCH
	if os == "amd64" {
		os = "x86_64"
	}

	url := config.GithubUrl + "/releases/download/" + config.GetNewestVersion() + "/mt_" + strings.TrimPrefix(config.GetNewestVersion(), "v") + "_" + strings.Title(runtime.GOOS) + "_" + os + archive
	return url
}

func updateMacLinux(url string) {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Download
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Gzip
	gzr, err := gzip.NewReader(resp.Body)
	if err != nil {
		fmt.Println("gzip error:", err)
		return
	}
	defer gzr.Close()

	// Tar
	tr := tar.NewReader(gzr)
	var tmpPath string

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		name := filepath.Base(header.Name)

		if name == "mt" || name == "mt.exe" {
			dir := filepath.Dir(exePath)
			tmpFile, err := os.CreateTemp(dir, "mt-*")
			if err != nil {
				if runtime.GOOS == "linux" {
					if os.Geteuid() != 0 {
						fmt.Println(("You need to run this command with sudo to update"))
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
				return
			}
			defer tmpFile.Close()

			// write binary
			io.Copy(tmpFile, tr)
			tmpPath = tmpFile.Name()
			break
		}
	}

	if tmpPath == "" {
		fmt.Println("Binary not found")
		return
	}

	os.Chmod(tmpPath, 0755)

	// Replace
	err = os.Rename(tmpPath, exePath)
	if err != nil {
		fmt.Println(ui.Red("Error during Installation:"+"%s"), err)
		return
	}

	successfullyUpdated()
}

func updateWindows(url string) {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("exe error:", err)
		return
	}

	tmpDir := os.TempDir()

	zipPath := filepath.Join(tmpDir, "mt-update.zip")
	newExePath := filepath.Join(tmpDir, "mt-new.exe")
	helperPath := filepath.Join(tmpDir, "mt-update.bat")

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("download failed:", err)
		return
	}
	defer resp.Body.Close()

	f, err := os.Create(zipPath)
	if err != nil {
		fmt.Println("create zip failed:", err)
		return
	}

	_, err = io.Copy(f, resp.Body)
	f.Close()
	if err != nil {
		fmt.Println("write zip failed:", err)
		return
	}

	r, err := zip.OpenReader(zipPath)
	if err != nil {
		fmt.Println("zip open failed:", err)
		return
	}
	defer r.Close()

	found := false

	for _, file := range r.File {
		name := filepath.Base(file.Name)

		if strings.HasSuffix(name, ".exe") {
			rc, err := file.Open()
			if err != nil {
				fmt.Println(err)
				return
			}

			out, err := os.Create(newExePath)
			if err != nil {
				rc.Close()
				fmt.Println(err)
				return
			}

			_, err = io.Copy(out, rc)

			out.Close()
			rc.Close()

			if err != nil {
				fmt.Println(err)
				return
			}

			found = true
			break
		}
	}

	if !found {
		fmt.Println("no exe found in zip")
		return
	}

	// CMD helper
	script := fmt.Sprintf(`
@echo off
timeout /t 2 /nobreak > nul

:loop
tasklist | find /i "mt.exe" > nul
if not errorlevel 1 (
    timeout /t 1 > nul
    goto loop
)

move /Y "%s" "%s"
start "" "%s"
`, newExePath, exePath, exePath)

	err = os.WriteFile(helperPath, []byte(script), 0644)
	if err != nil {
		fmt.Println("helper write failed:", err)
		return
	}

	cmd := exec.Command("cmd", "/C", helperPath)
	err = cmd.Start()
	if err != nil {
		fmt.Println("helper start failed:", err)
		return
	}
	successfullyUpdated()
}

func successfullyUpdated() {
	fmt.Println(ui.GlowPink(config.Logo))
	fmt.Printf(ui.Green("Successfully Updated the Cli to version: %s \n"), config.GetNewestVersion())
	os.Exit(0)
}
