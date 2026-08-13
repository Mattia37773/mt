/*
Copyright © 2026 Matze
*/
package single

import (
	"bytes"
	"os/exec"
	"runtime"
	"testing"

	"github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/config"
	_ "github.com/mattia37773/mt/config"
	base "github.com/mattia37773/mt/tests/basetest"
	"github.com/stretchr/testify/assert"
)

func TestUpdateFromSource(t *testing.T) {
	base.ChangeDirToBin(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	binary := "./source-install"
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}

	// updateRejectCmd := exec.Command("../../bin/source-install", "update", "--confirm=false")
	// updateRejectOut, _ := updateRejectCmd.Output()

	// assert.Contains(t, base.StripANSI(string(updateRejectOut)),
	// 	"mt version v1.0.0",
	// )

	versionCmd := exec.Command(binary, "--version")
	versionOut, _ := versionCmd.Output()

	assert.Contains(t, base.StripANSI(string(versionOut)),
		"mt version v1.0.0",
	)

	updateCmd := exec.Command(binary, "update", "--confirm")
	updateOut, _ := updateCmd.Output()

	assert.Contains(t, base.StripANSI(string(updateOut)),
		"Successfully Updated the Cli to version: ",
	)

	updateAfterCmd := exec.Command(binary, "update", "--confirm")
	updateAfterOut, _ := updateAfterCmd.Output()

	assert.Contains(t, base.StripANSI(string(updateAfterOut)),
		"The Newest Version is already installed",
	)
}

func TestUpdateGo(t *testing.T) {
	base.ChangeDirToBin(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	binary := "./go-install"
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}

	updateCmd := exec.Command(binary, "update", "--confirm")
	updateOut, _ := updateCmd.Output()

	assert.Contains(t, base.StripANSI(string(updateOut)),
		"Installed via go",
		"Please update with the follwing commands",
		"clear the package cache: go clean -modcache",
		"install the update: go install "+config.AppConfig.ModulePath+"@latest",
	)
}

func TestUpdateHomebrew(t *testing.T) {
	base.ChangeDirToBin(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	binary := "./homebrew-install"
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}

	updateCmd := exec.Command(binary, "update", "--confirm")
	updateOut, _ := updateCmd.Output()

	assert.Contains(t, base.StripANSI(string(updateOut)),
		"Installed via Homebrew.",
		"Please update with: brew upgrade mt",
	)
}
