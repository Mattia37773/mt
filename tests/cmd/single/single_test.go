/*
Copyright © 2026 Matze
*/
package single_test

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/mattia37773/mt/cmd"
	_ "github.com/mattia37773/mt/cmd/db"
	_ "github.com/mattia37773/mt/cmd/php"
	_ "github.com/mattia37773/mt/cmd/stack"
	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/helper/docker"
	base "github.com/mattia37773/mt/tests/basetest"
	"github.com/stretchr/testify/assert"
)

func TestSingles(t *testing.T) {
	t.Run("Start", func(t *testing.T) {
		testStartCmd(t)
		time.Sleep(30 * time.Second)
	})

	t.Run("Composer help", func(t *testing.T) {
		testComposerHelp(t)
	})

	t.Run("Run with --help argument", func(t *testing.T) {
		testRunWithHelp(t)
	})

	t.Run("Symfony console test", func(t *testing.T) {
		testConsole(t)
	})

	t.Run("Destory Everything", func(t *testing.T) {
		testDestroy(t)
	})

	//*
	//// * Negative tests
	//*

	t.Run("Shell without container argument", func(t *testing.T) {
		testShellWithoutContainerArgument(t)
	})

	t.Run("Shell without flag argument", func(t *testing.T) {
		testShellWithoutFlagArgument(t)
	})

	//*
	//// * not running tests
	//*

	t.Run("Console without container", func(t *testing.T) {
		testConsoleContainerNotRunning(t)
	})

	t.Run("Craft without container", func(t *testing.T) {
		testCraftContainerNotRunning(t)
	})

	t.Run("Composer without container", func(t *testing.T) {
		testComposerContainerNotRunning(t)
	})

	t.Run("Run without container", func(t *testing.T) {
		testRunContainerNotRunning(t)
	})
}

func testStartCmd(t *testing.T) {
	base.ChangeDirToSymfony(t)

	rootCmd := cmd.RootCmd
	base.SilentCommand(t, rootCmd, []string{"stack", "start"})

	expectedContainers := []string{
		config.ProjectConfig.ProjectName + "-db",
		config.ProjectConfig.ProjectName + "-fpm",
		config.ProjectConfig.ProjectName + "-mailpit",
		config.ProjectConfig.ProjectName + "-nginx",
		config.ProjectConfig.ProjectName + "-phpmyadmin",
	}

	base.WaitForContainersRunning(t, expectedContainers, 3*time.Minute)

	for _, container := range expectedContainers {
		container := container
		t.Run("Check_Container_"+container, func(t *testing.T) {
			running := docker.IsContainerRunning(container)
			assert.True(t, running, "The Container '%s'should run but doesn't", container)
		})
	}
}

func testComposerHelp(t *testing.T) {
	base.ChangeDirToSymfony(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"composer", "--help"})

	errExec := base.ExecuteCommand(rootCmd)
	assert.NoError(t, errExec)
	cleanOutput := base.StripANSI(buf.String())

	lastLine := base.GetLineFromStringReversed(cleanOutput, 1)
	assert.Contains(t, lastLine, "/usr/local/bin/composer list --raw")
}

func testRunWithHelp(t *testing.T) {
	base.ChangeDirToSymfony(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"run", "ls", "--help"})
	errExec := base.ExecuteCommand(rootCmd)
	assert.NoError(t, errExec)
	cleanOutput := base.StripANSI(buf.String())

	lastLine := base.GetLineFromStringReversed(cleanOutput, 1)
	assert.Contains(t, lastLine, "or available locally via: info '(coreutils) ls invocation'")
}

func testConsole(t *testing.T) {
	base.ChangeDirToSymfony(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"php", "console"})
	errExec := base.ExecuteCommand(rootCmd)
	assert.NoError(t, errExec)
	cleanOutput := base.StripANSI(buf.String())

	assert.Contains(t, cleanOutput, "Fatal error: Uncaught LogicException: Dependencies are missing. Try running \"composer install\". in /app/bin/console:")
}

func testDestroy(t *testing.T) {
	base.ChangeDirToSymfony(t)
	db := config.GetDbConfig()

	// // delete vendor directory
	// vendorErr := os.RemoveAll("vendor")
	// assert.Nil(t, vendorErr)

	dbDumpErr := os.RemoveAll(config.ProjectConfig.ProjectName + db.Filetype)
	assert.Nil(t, dbDumpErr)

	rootCmd := cmd.RootCmd
	base.SilentCommand(t, rootCmd, []string{"stack", "destroy"})
}

//*
//// * Negative tests
//*

func testShellWithoutContainerArgument(t *testing.T) {
	base.ChangeDirToSymfony(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"shell"})
	errExec := base.ExecuteCommand(rootCmd)

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"Missing argument CONTAINER",
	)
}

func testShellWithoutFlagArgument(t *testing.T) {
	base.ChangeDirToSymfony(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"shell", "fpm", "-u"})
	errExec := base.ExecuteCommand(rootCmd)

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"flag needs an argument: 'u' in -u",
	)
}

//*
//// * not running tests
//*

func testConsoleContainerNotRunning(t *testing.T) {
	base.ChangeDirToSymfony(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"php", "console", "symfony > laravel"})
	errExec := base.ExecuteCommand(rootCmd)
	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The Container: "+config.ProjectConfig.ProjectName+
			"-"+config.ProjectConfig.Backend.ContainerName+
			" doesn't exist",
		"Did you forget to run mt stack start?",
	)

}

func testCraftContainerNotRunning(t *testing.T) {
	base.ChangeDirToSymfony(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"php", "craft", "5 is worse then 4"})
	errExec := base.ExecuteCommand(rootCmd)

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The Container: "+config.ProjectConfig.ProjectName+`-`+config.ProjectConfig.Backend.ContainerName+" doesn't exist",
		"Did you forget to run mt stack start?",
	)
}

func testComposerContainerNotRunning(t *testing.T) {
	base.ChangeDirToSymfony(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"composer", "blabla"})
	errExec := base.ExecuteCommand(rootCmd)

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The Container: "+config.ProjectConfig.ProjectName+`-`+config.ProjectConfig.Backend.ContainerName+" doesn't exist",
		"Did you forget to run mt stack start?",
	)
}

func testRunContainerNotRunning(t *testing.T) {
	base.ChangeDirToSymfony(t)
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"run", "i hate testing"})
	errExec := base.ExecuteCommand(rootCmd)

	assert.Error(t, errExec)
	assert.Contains(t, errExec.Error(),
		"The Container: "+config.ProjectConfig.ProjectName+`-`+config.ProjectConfig.Main.ContainerName+" doesn't exist",
		"Did you forget to run mt stack start?",
	)
}
