/*
Copyright © 2026 Matze
*/
package single

import (
	"os/exec"
	"runtime"
	"testing"

	_ "github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/config"
	_ "github.com/mattia37773/mt/config"
	base "github.com/mattia37773/mt/tests/basetest"
	"github.com/stretchr/testify/assert"
)

func TestUpdateFromSource(t *testing.T) {
	base.ChangeDirToBin(t)

	binary := "./source-install"
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}

	versionCmd := exec.Command(binary, "--version")
	versionOut, err := versionCmd.CombinedOutput()
	assert.NoError(t, err)

	assert.Contains(t, base.StripANSI(string(versionOut)),
		"mt version v1.0.0",
	)

	updateCmd := exec.Command(binary, "update", "--confirm")
	updateOut, err := updateCmd.CombinedOutput()
	assert.NoError(t, err)

	assert.Contains(t, base.StripANSI(string(updateOut)),
		"Successfully Updated the Cli to version: ",
	)

	updateAfterCmd := exec.Command(binary, "update", "--confirm")
	updateAfterOut, err := updateAfterCmd.CombinedOutput()
	assert.NoError(t, err)

	assert.Contains(t, base.StripANSI(string(updateAfterOut)),
		"The Newest Version is already installed",
	)
}

func TestUpdateGo(t *testing.T) {
	base.ChangeDirToBin(t)

	binary := "./go-install"
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}

	updateCmd := exec.Command(binary, "update", "--confirm")
	updateOut, err := updateCmd.CombinedOutput()
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(string(updateOut))
	assert.Contains(t, cleanOutput, "Installed via go")
	assert.Contains(t, cleanOutput, "Please update with the follwing commands")
	assert.Contains(t, cleanOutput, "clear the package cache: go clean -modcache")
	assert.Contains(t, cleanOutput, "install the update: go install "+config.AppConfig.ModulePath+"@latest")
}

func TestUpdateHomebrew(t *testing.T) {
	base.ChangeDirToBin(t)

	binary := "./homebrew-install"
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}

	updateCmd := exec.Command(binary, "update", "--confirm")
	updateOut, err := updateCmd.CombinedOutput()
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(string(updateOut))
	assert.Contains(t, cleanOutput, "Installed via Homebrew.")
	assert.Contains(t, cleanOutput, "Please update with: brew upgrade mt")
}
