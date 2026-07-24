package php

import (
	"bytes"
	"testing"

	"github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/tests/base"
	"github.com/stretchr/testify/assert"
)

func TestConsoleCommandGenSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	consoleCommand := consoleGen(rootCmd.OutOrStdout(), []string{"doctrine:migrations:migrate", "-n"})
	assert.Equal(t, []string{"exec", config.ProjectConfig.ProjectName + "-fpm", "sh", "-c", "bin/console doctrine:migrations:migrate -n"}, consoleCommand)
}

func TestCraftCommandGenSuccess(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	craftCommand := craftGen(rootCmd.OutOrStdout(), []string{"clear-caches/all"})
	assert.Equal(t, []string{"exec", config.ProjectConfig.ProjectName + "-fpm", "sh", "-c", "php craft clear-caches/all"}, craftCommand)
}
