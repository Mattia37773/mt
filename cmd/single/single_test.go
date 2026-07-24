package single

import (
	"bytes"
	"testing"

	"github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/tests/base"
	"github.com/stretchr/testify/assert"
)

func TestComposerCommandGenSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	composerCommand := composerGen(rootCmd.OutOrStdout(), []string{"install"})
	assert.Equal(t, []string{"exec", config.ProjectConfig.ProjectName + "-fpm", "sh", "-c", "composer install"}, composerCommand)
}

func TestRunCommandGenSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	runCommand := runGen(rootCmd.OutOrStdout(), []string{"ls", "--help"})
	assert.Equal(t, []string{"exec", config.ProjectConfig.ProjectName + "-fpm", "sh", "-c", " ls --help"}, runCommand)
}

func TestShellCommandGenWithoutUserSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	shellCommand := shellGen(rootCmd.OutOrStdout(), "", []string{"nginx"})
	assert.Equal(t, []string{"sh", "-c", "docker exec -it " + config.ProjectConfig.ProjectName + "-nginx $(docker exec " + config.ProjectConfig.ProjectName + "-nginx sh -c 'command -v bash || echo /bin/sh')"}, shellCommand)
}

func TestShellCommandGenWithUserSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	shellCommand := shellGen(rootCmd.OutOrStdout(), "www-data", []string{"fpm"})
	assert.Equal(t, []string{"sh", "-c", "docker exec -u www-data -it " + config.ProjectConfig.ProjectName + "-fpm $(docker exec " + config.ProjectConfig.ProjectName + "-fpm sh -c 'command -v bash || echo /bin/sh')"}, shellCommand)
}
