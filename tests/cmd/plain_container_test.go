/*
Copyright © 2026 Matze
*/
package cmd_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/mattia37773/mt/cmd"
	_ "github.com/mattia37773/mt/cmd/db"
	_ "github.com/mattia37773/mt/cmd/php"
	_ "github.com/mattia37773/mt/cmd/single"
	_ "github.com/mattia37773/mt/cmd/stack"
	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/helper/docker"
	base "github.com/mattia37773/mt/tests/basetest"
	"github.com/stretchr/testify/assert"
)

func TestPlainContainer(t *testing.T) {
	t.Run("Start Debian Stack", func(t *testing.T) {
		testStartDebianStack(t)
	})

	//*
	//// * db
	//*

	t.Run("Mysql shell command not found", func(t *testing.T) {
		testMysqlShellCommandNotFound(t)
	})

	t.Run("Mysql export command not found", func(t *testing.T) {
		testMysqlExportCommandNotFound(t)
	})

	//*
	//// * php
	//*

	t.Run("Console command not found", func(t *testing.T) {
		testConsoleCommandNotFound(t)
	})

	t.Run("Craft command not found", func(t *testing.T) {
		testCraftCommandNotFound(t)
	})

	// *
	//// * single
	// *

	t.Run("Compser command not found", func(t *testing.T) {
		testComposerCommandNotFound(t)
	})

	t.Run("Run command not found", func(t *testing.T) {
		testRunCommandNotFound(t)
	})

	// *
	//// * destroy
	// *

	t.Run("Destroy the container", func(t *testing.T) {
		testDestroyPlainContainer(t)
	})

}

func testStartDebianStack(t *testing.T) {
	base.ChangeDirToDefault(t)

	rootCmd := cmd.RootCmd
	base.SilentCommand(t, rootCmd, []string{"stack", "start"})

	expectedContainers := []string{
		config.ProjectConfig.ProjectName + "-" + config.ProjectConfig.Main.ContainerName,
	}

	base.WaitForContainersRunning(t, expectedContainers, 3*time.Minute)

	for _, container := range expectedContainers {
		t.Run("Check_Container_"+container, func(t *testing.T) {
			running := docker.IsContainerRunning(container)
			assert.True(t, running, "The Container '%s'should run but doesn't", container)
		})
	}
}

//*
//// * db
//*

func testMysqlShellCommandNotFound(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"db", "shell"})
	errExec := base.ExecuteCommand(rootCmd)

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The command: mysql -u"+config.ProjectConfig.DB.User+" -p"+config.ProjectConfig.DB.Password+" isn't available inside the "+config.ProjectConfig.DB.ContainerName,
	)
}

func testMysqlExportCommandNotFound(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"db", "export"})
	errExec := base.ExecuteCommand(rootCmd)

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The command: mysqldump -u"+config.ProjectConfig.DB.User+" -p"+config.ProjectConfig.DB.Password+" "+config.ProjectConfig.DB.Name+" > /tmp/nixy.sql isn't available inside the "+config.ProjectConfig.DB.ContainerName,
	)
}

//*
//// * php
//*

func testConsoleCommandNotFound(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"php", "console"})
	errExec := base.ExecuteCommand(rootCmd)

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The command: bin/console isn't available inside the "+config.ProjectConfig.DB.ContainerName,
	)
}

func testCraftCommandNotFound(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"php", "craft"})
	errExec := base.ExecuteCommand(rootCmd)

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The command: php craft isn't available inside the "+config.ProjectConfig.DB.ContainerName,
	)
}

//*
//// * single
//*

func testComposerCommandNotFound(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"composer"})
	errExec := base.ExecuteCommand(rootCmd)

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The command: composer isn't available inside the "+config.ProjectConfig.DB.ContainerName,
	)
}

func testRunCommandNotFound(t *testing.T) {
	base.ChangeDirToDefault(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"run", "olive"})
	errExec := base.ExecuteCommand(rootCmd)
	assert.NoError(t, errExec)
	cleanOutput := base.StripANSI(buf.String())

	lastLine := base.GetLineFromStringReversed(cleanOutput, 1)
	assert.Contains(t, lastLine, "sh: 1: olive: not found")
}

func testDestroyPlainContainer(t *testing.T) {
	base.ChangeDirToDefault(t)

	rootCmd := cmd.RootCmd
	base.SilentCommand(t, rootCmd, []string{"stack", "destroy"})
}
