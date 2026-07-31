package stack

import (
	"bytes"
	"testing"

	"github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/config"
	_ "github.com/mattia37773/mt/helper/basecmd"
	base "github.com/mattia37773/mt/helper/basetest"

	"github.com/stretchr/testify/assert"
)

func TestBuildCommandGenSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	buildCommand, _ := buildGen(rootCmd.OutOrStdout())
	assert.Equal(t, []string{"sh", "-c", "docker compose -f " + config.ProjectConfig.Paths.DockerCompose + " build --no-cache"}, buildCommand)
}

func TestStartCommandGenSuccess(t *testing.T) {

	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	startCommand, _ := startGen(rootCmd.OutOrStdout())
	assert.Equal(t, []string{"sh", "-c", "docker compose -f " + config.ProjectConfig.Paths.DockerCompose + " up -d"}, startCommand)
}

func TestDestroyCommandGenSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	destroyCommand, _ := destroyGen(rootCmd.OutOrStdout())
	assert.Equal(t, []string{"sh", "-c", "docker compose -f " + config.ProjectConfig.Paths.DockerCompose + " down -v --rmi all"}, destroyCommand)
}

func TestPsCommandGenSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	psCommand, _ := psGen(rootCmd.OutOrStdout())
	assert.Equal(t, []string{"sh", "-c", "docker compose -f " + config.ProjectConfig.Paths.DockerCompose + " ps"}, psCommand)
}

func TestRestartCommandGenSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	restartCommand, _ := restartGen(rootCmd.OutOrStdout())
	assert.Equal(t, []string{"sh", "-c", "docker compose -f " + config.ProjectConfig.Paths.DockerCompose + " restart"}, restartCommand)
}

func TestStopCommandGenSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	stopCommand, _ := stopGen(rootCmd.OutOrStdout())
	assert.Equal(t, []string{"sh", "-c", "docker compose -f " + config.ProjectConfig.Paths.DockerCompose + " down"}, stopCommand)
}
