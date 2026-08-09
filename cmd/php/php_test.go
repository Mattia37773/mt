package php

import (
	"bytes"
	"testing"

	"github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/config"
	base "github.com/mattia37773/mt/tests/basetest"
	"github.com/stretchr/testify/assert"
)

func TestConsoleCommandGenSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	consoleCommand, _ := consoleGen(rootCmd.OutOrStdout(), []string{"doctrine:migrations:migrate", "-n"})
	assert.Equal(t, []string{"exec", config.ProjectConfig.ProjectName + "-fpm", "sh", "-c", "bin/console doctrine:migrations:migrate -n"}, consoleCommand)
}

func TestCraftCommandGenSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	craftCommand, _ := craftGen(rootCmd.OutOrStdout(), []string{"clear-caches/all"})
	assert.Equal(t, []string{"exec", config.ProjectConfig.ProjectName + "-fpm", "sh", "-c", "php craft clear-caches/all"}, craftCommand)
}
