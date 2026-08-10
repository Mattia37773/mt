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

	"github.com/mattia37773/mt/ui/form"
	"github.com/mattia37773/mt/ui/text"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:                   "update",
	Short:                 "Updates the cli",
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	RunE: func(c *cobra.Command, args []string) error {
		err := updateSingle(c.OutOrStdout())
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(updateCmd)
}

func updateSingle(out io.Writer) error {
	if config.Version == config.LatestVersion {
		fmt.Fprintln(out, text.GlowPink(config.AppConfig.Logo))
		fmt.Fprintln(out, text.Green("The Newest Version is already installed"))
		return nil
	}

	switch buildmethod := config.BuildMethod; buildmethod {
	case "github":
		confirm, err := form.Confirm("Do you want to pull the latest updates from GitHub?")
		if err != nil {
			return fmt.Errorf("%s", err)
		}
		if confirm == true {
			githubReleasesUpdate(out)
		} else {
			fmt.Fprintln(out, "Continuing without updating.")
		}
	case "source":
		confirm, err := form.Confirm("You compiled this project yourself.\nDo you want to pull the latest updates from GitHub?")
		if err != nil {
			return fmt.Errorf("%s", err)
		}
		if confirm == true {
			err := githubReleasesUpdate(out)
			if err != nil {
				return err
			}
		} else {
			fmt.Fprintln(out, "Continuing without updating.")
		}
		return nil
	case "go":
		cmd := `go install -ldflags "-X ` + config.AppConfig.ModulePath + `/config.BuildMethod=go" ` + config.AppConfig.ModulePath + `@latest`

		fmt.Fprintln(out, text.Green("Installed via go"))
		fmt.Fprintln(out, "Please update with the follwing commands")
		fmt.Fprintln(out, "clear the package cache: go clean -modcache")
		fmt.Fprintf(out, "install the update: \"%s\"\n", cmd)
	case "homebrew":
		fmt.Fprintln(out, text.Green("Installed via Homebrew."))
		fmt.Fprintln(out, "Please update with \"brew upgrade mt\".")
	default:
		fmt.Fprintln(out, "There is currently not an option for this build method.")
		fmt.Fprintln(out, "This is only possible by manually modifiy the source code")
	}
	return nil
}

func githubReleasesUpdate(out io.Writer) error {
	switch userOs := runtime.GOOS; userOs {
	case "darwin":
		url := getUrl(".tar.gz")
		err := updateMacLinux(out, url)

		if err != nil {
			return err
		}
	case "linux":
		url := getUrl(".tar.gz")
		err := updateMacLinux(out, url)

		if err != nil {
			return err
		}
	case "windows":
		url := getUrl(".zip")
		err := updateWindows(out, url)

		if err != nil {
			return err
		}
	default:
		fmt.Fprintln(out, text.Red("Sorry but your operation system is not supported"))
		fmt.Fprintf(out, "You can check the source code out and add support for your operatingsystem ;)")
	}
	return nil
}

func getUrl(archive string) string {
	os := runtime.GOARCH
	if os == "amd64" {
		os = "x86_64"
	}

	url := config.AppConfig.GithubUrl + "/releases/download/" + config.LatestVersion + "/mt_" + strings.TrimPrefix(config.LatestVersion, "v") + "_" + strings.Title(runtime.GOOS) + "_" + os + archive
	return url
}

func updateMacLinux(out io.Writer, url string) error {
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	// Download
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defer resp.Body.Close()

	// Gzip
	gzr, err := gzip.NewReader(resp.Body)
	if err != nil {
		return fmt.Errorf("%s", err)
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
			return fmt.Errorf("%s", err)
		}
		name := filepath.Base(header.Name)

		if name == "mt" || name == "mt.exe" {
			dir := filepath.Dir(exePath)
			tmpFile, err := os.CreateTemp(dir, "mt-*")
			if err != nil {
				if runtime.GOOS == "linux" {
					if os.Geteuid() != 0 {
						fmt.Fprintln(out, "You need to run this command with sudo to update")
					} else {
						return fmt.Errorf("%s", err)
					}
				} else {
					return fmt.Errorf("%s", err)
				}
				return fmt.Errorf("%s", err)
			}
			defer tmpFile.Close()

			// write binary
			io.Copy(tmpFile, tr)
			tmpPath = tmpFile.Name()
			break
		}
	}

	if tmpPath == "" {
		return fmt.Errorf("Binary not found", err)
	}

	os.Chmod(tmpPath, 0755)

	// Replace
	err = os.Rename(tmpPath, exePath)
	if err != nil {
		return fmt.Errorf("Somehting went wrint durint the installation: %s", err)
	}

	successfullyUpdated(out)
	return nil
}

func updateWindows(out io.Writer, url string) error {
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	tmpDir := os.TempDir()

	zipPath := filepath.Join(tmpDir, "mt-update.zip")
	newExePath := filepath.Join(tmpDir, "mt-new.exe")
	helperPath := filepath.Join(tmpDir, "mt-update.bat")

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("download failed", err)
	}
	defer resp.Body.Close()

	f, err := os.Create(zipPath)
	if err != nil {
		return fmt.Errorf("creation of zip file failed", err)
	}

	_, err = io.Copy(f, resp.Body)
	f.Close()
	if err != nil {
		return fmt.Errorf("writing the zip file failed", err)
	}

	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return fmt.Errorf("opening the zip failed", err)
	}
	defer r.Close()

	found := false

	for _, file := range r.File {
		name := filepath.Base(file.Name)

		if strings.HasSuffix(name, ".exe") {
			rc, err := file.Open()
			if err != nil {
				return fmt.Errorf("%s", err)
			}

			out, err := os.Create(newExePath)
			if err != nil {
				rc.Close()
				return fmt.Errorf("%s", err)
			}

			_, err = io.Copy(out, rc)

			out.Close()
			rc.Close()

			if err != nil {
				return fmt.Errorf("%s", err)
			}

			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("no exe found in zip", found)
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
		return fmt.Errorf("helper write failed: %s", err)
	}

	cmd := exec.Command("cmd", "/C", helperPath)
	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("helper start failed: %s", err)
	}
	successfullyUpdated(out)
	return nil
}

func successfullyUpdated(out io.Writer) {
	fmt.Fprintln(out, text.GlowPink(config.AppConfig.Logo))
	fmt.Fprintf(out, text.Green("Successfully Updated the Cli to version: %s \n"), config.LatestVersion)
}
