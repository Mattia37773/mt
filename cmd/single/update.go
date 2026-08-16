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
	"path/filepath"
	"runtime"
	"strings"

	"github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/config"

	"github.com/mattia37773/mt/ui/form"
	"github.com/mattia37773/mt/ui/text"

	"github.com/spf13/cobra"
)

var skipInteractiveUpdate bool

var updateCmd = &cobra.Command{
	Use:                   "update",
	Short:                 "Updates the cli",
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
	// todo flag setting to nointeractivly say no
	updateCmd.Flags().BoolVar(
		&skipInteractiveUpdate,
		"confirm",
		false,
		"Skip interactive confirmation",
	)
	cmd.RootCmd.AddCommand(updateCmd)
}

func updateSingle(out io.Writer) error {
	if config.LatestVersion == "" || config.Version == config.LatestVersion {
		fmt.Fprintln(out, text.GlowPink(config.AppConfig.Logo))
		fmt.Fprintln(out, text.Green("The Newest Version is already installed"))
		return nil
	}

	switch buildmethod := config.BuildMethod; buildmethod {
	case "source":
		var confirm bool
		var err error = nil

		if skipInteractiveUpdate == false {
			confirm, err = form.Confirm("Do you want to pull the latest updates from GitHub?")
		} else {
			confirm = true
		}

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
		// TODO define how thats done
		cmd := `go install ` + config.AppConfig.ModulePath + `@latest`

		fmt.Fprintln(out, text.Green("Installed via go"))
		fmt.Fprintln(out, "Please update with the follwing commands")
		fmt.Fprintln(out, "clear the package cache: go clean -modcache")
		fmt.Fprintf(out, "install the update: %s\n", cmd)
	case "homebrew":
		fmt.Fprintln(out, text.Green("Installed via Homebrew."))
		fmt.Fprintln(out, "Please update with: brew upgrade mt")
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

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download update: HTTP %s", resp.Status)
	}

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

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("download failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download update: HTTP %s", resp.Status)
	}

	f, err := os.Create(zipPath)
	if err != nil {
		return fmt.Errorf("creation of zip file failed: %w", err)
	}

	_, err = io.Copy(f, resp.Body)
	f.Close()
	if err != nil {
		return fmt.Errorf("writing the zip file failed: %w", err)
	}

	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return fmt.Errorf("opening the zip failed: %w", err)
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

			outExe, err := os.Create(newExePath)
			if err != nil {
				rc.Close()
				return fmt.Errorf("%s", err)
			}

			_, err = io.Copy(outExe, rc)

			outExe.Close()
			rc.Close()

			if err != nil {
				return fmt.Errorf("%s", err)
			}

			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("no exe found in zip")
	}

	oldExePath := exePath + ".old"
	_ = os.Remove(oldExePath)
	err = os.Rename(exePath, oldExePath)
	if err != nil {
		return fmt.Errorf("failed to replace binary: %w", err)
	}

	err = os.Rename(newExePath, exePath)
	if err != nil {
		_ = os.Rename(oldExePath, exePath)
		return fmt.Errorf("failed to install new binary: %w", err)
	}

	_ = os.Remove(oldExePath)
	successfullyUpdated(out)
	return nil
}

func successfullyUpdated(out io.Writer) {
	fmt.Fprintln(out, text.GlowPink(config.AppConfig.Logo))
	fmt.Fprintf(out, text.Green("Successfully Updated the Cli to version: %s \n"), config.LatestVersion)
}
