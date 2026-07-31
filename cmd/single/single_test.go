/*
Copyright © 2026 Matze
*/
package single

import (
	"bytes"
	"os"
	"testing"

	"github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/config"
	base "github.com/mattia37773/mt/helper/basetest"
	"github.com/stretchr/testify/assert"
)

func TestComposerCommandGenSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	composerCommand, _ := composerGen(rootCmd.OutOrStdout(), []string{"install"})
	assert.Equal(t, []string{"exec", config.ProjectConfig.ProjectName + "-fpm", "sh", "-c", "composer install"}, composerCommand)
}

func TestRunCommandGenSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	runCommand, _ := runGen(rootCmd.OutOrStdout(), []string{"ls", "--help"})
	assert.Equal(t, []string{"exec", config.ProjectConfig.ProjectName + "-fpm", "sh", "-c", "ls --help"}, runCommand)
}

func TestShellCommandGenWithoutUserSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	shellCommand, _ := shellGen(rootCmd.OutOrStdout(), "", []string{"nginx"})

	assert.Equal(t, []string{"sh", "-c", "docker exec -it " + config.ProjectConfig.ProjectName + "-nginx $(docker exec " + config.ProjectConfig.ProjectName + "-nginx sh -c 'command -v bash || echo /bin/sh')"}, shellCommand)
}

func TestShellCommandGenWithUserSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	shellCommand, _ := shellGen(rootCmd.OutOrStdout(), "www-data", []string{"fpm"})
	assert.Equal(t, []string{"sh", "-c", "docker exec -u www-data -it " + config.ProjectConfig.ProjectName + "-fpm $(docker exec " + config.ProjectConfig.ProjectName + "-fpm sh -c 'command -v bash || echo /bin/sh')"}, shellCommand)
}

func TestConfigCommandGen(t *testing.T) {
	base.ChangeDirToDefaultConfig(t)
	rootCmd := cmd.RootCmd
	os.Remove(".mt.yaml")

	// ! This doesn't work for some reason
	// it doesnt generate the file
	// buf := new(bytes.Buffer)
	// rootCmd.SetOut(buf)
	// rootCmd.SetErr(buf)

	// rootCmd.SetArgs([]string{"config"})

	// errExec := rootCmd.Execute()
	// assert.NoError(t, errExec)

	errGenerate := generateConfigFile(rootCmd.OutOrStdout())
	assert.NoError(t, errGenerate)

	assert.FileExists(t, ".mt.yaml")

	base.AssertFilesEqual(t, "default.yaml", ".mt.yaml")
	os.Remove(".mt.yaml")
}
