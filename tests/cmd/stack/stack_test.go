/*
Copyright © 2026 Matze
*/
package stack_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/mattia37773/mt/cmd"
	_ "github.com/mattia37773/mt/cmd/php"
	_ "github.com/mattia37773/mt/cmd/single"
	_ "github.com/mattia37773/mt/cmd/stack"
	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/helper/docker"
	base "github.com/mattia37773/mt/tests/basetest"
	"github.com/stretchr/testify/assert"
)

// TODO add os specific flags
type stackConfig struct {
	containers  []string
	ProjectPath string
}

type testCase struct {
	name   string
	config stackConfig
}

// TODO choose a much smaler test project ;=}
func TestStack(t *testing.T) {
	testCases := []testCase{
		{
			name: "Symfony",
			config: stackConfig{
				containers: []string{
					"db",
					"fpm",
					"nginx",
				},
				ProjectPath: "symfony",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name+" Start Stack", func(t *testing.T) {
			testStartCmd(t, tc)
		})

		t.Run(tc.name+" PS", func(t *testing.T) {
			testPsCmd(t, tc)
		})

		t.Run(tc.name+" Logs", func(t *testing.T) {
			testLogsCmd(t, tc)
		})

		t.Run(tc.name+" Logs With Container Flag", func(t *testing.T) {
			testLogsCmdSpecificContainer(t, tc)
		})

		t.Run(tc.name+" Restart Stack", func(t *testing.T) {
			testRestartCmd(t, tc)
		})

		t.Run(tc.name+" Stop Stack", func(t *testing.T) {
			testStopCmd(t, tc)
		})

		t.Run(tc.name+" Destroy Stack", func(t *testing.T) {
			testDestroyCmd(t, tc)
		})
	}
}

func testStartCmd(t *testing.T, tc testCase) {
	base.ChangeDir(t, tc.config.ProjectPath)

	// create the containernames
	expected := make([]string, len(tc.config.containers))
	for i, name := range tc.config.containers {
		expected[i] = config.ProjectConfig.ProjectName + "-" + name
	}

	rootCmd := cmd.RootCmd
	base.SilentCommand(t, rootCmd, []string{"stack", "start"})

	base.WaitForContainersRunning(t, expected, 3*time.Minute)

	for _, container := range expected {
		container := container
		t.Run("Check_Container_"+container, func(t *testing.T) {
			running := docker.IsContainerRunning(container)
			assert.True(t, running, "The Container '%s'should run but doesn't", container)
		})
	}
}

func testPsCmd(t *testing.T, tc testCase) {
	base.ChangeDir(t, tc.config.ProjectPath)

	// create the containernames
	expected := make([]string, len(tc.config.containers))
	for i, name := range tc.config.containers {
		expected[i] = config.ProjectConfig.ProjectName + "-" + name
	}

	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"stack", "ps"})

	err := base.ExecuteCommand(rootCmd)
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(buf.String())

	assert.Contains(t, cleanOutput, "NAME", "Output should contain column header 'NAME'")
	assert.Contains(t, cleanOutput, "STATUS", "Output should contain column header 'STATUS'")

	for _, container := range expected {
		assert.Contains(t, cleanOutput, container, "The 'ps' output should contain container '%s'", container)
	}
}

func testLogsCmd(t *testing.T, tc testCase) {
	base.ChangeDir(t, tc.config.ProjectPath)

	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"stack", "logs"})

	err := base.ExecuteCommand(rootCmd)
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(buf.String())

	for _, service := range tc.config.containers {
		assert.Contains(t, cleanOutput, service, "The logs output should contain entries from '%s'", service)
	}
}

func testLogsCmdSpecificContainer(t *testing.T, tc testCase) {
	base.ChangeDir(t, tc.config.ProjectPath)

	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"stack", "logs", config.ProjectConfig.Main.ContainerName})

	err := base.ExecuteCommand(rootCmd)
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(buf.String())

	for _, service := range []string{config.ProjectConfig.Main.ContainerName} {
		assert.Contains(t, cleanOutput, service, "The logs output should contain entries from '%s'", service)
	}
}

func testRestartCmd(t *testing.T, tc testCase) {
	base.ChangeDir(t, tc.config.ProjectPath)

	// create the containernames
	expected := make([]string, len(tc.config.containers))
	for i, name := range tc.config.containers {
		expected[i] = config.ProjectConfig.ProjectName + "-" + name
	}

	// Save start timestamps of all containers BEFORE restart
	timestampsBefore := make(map[string]string)
	for _, container := range expected {
		startTime := docker.GetContainerStartedAt(container)
		assert.NotEmpty(t, startTime, "Container '%s' must be running before test", container)
		timestampsBefore[container] = startTime
	}

	time.Sleep(1 * time.Second)

	buf := new(bytes.Buffer)
	rootCmd := cmd.RootCmd
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"stack", "restart"})

	err := base.ExecuteCommand(rootCmd)
	assert.NoError(t, err)

	for _, container := range expected {
		container := container
		t.Run("Check_Restart_"+container, func(t *testing.T) {
			startTimeAfter := docker.GetContainerStartedAt(container)
			assert.NotEmpty(t, startTimeAfter, "Container '%s' should still exist after restart", container)

			// Timestamp must be different/newer than before
			assert.NotEqual(t, timestampsBefore[container], startTimeAfter, "The container '%s' was not restarted", container)
		})
	}
}

func testStopCmd(t *testing.T, tc testCase) {
	base.ChangeDir(t, tc.config.ProjectPath)

	// create the containernames
	expected := make([]string, len(tc.config.containers))
	for i, name := range tc.config.containers {
		expected[i] = config.ProjectConfig.ProjectName + "-" + name
	}

	rootCmd := cmd.RootCmd
	base.SilentCommand(t, rootCmd, []string{"stack", "stop"})

	for _, container := range expected {
		container := container
		t.Run("Check_Container_DoesNotExist_"+container, func(t *testing.T) {
			status := docker.GetContainerStatus(container)

			assert.Empty(t, status, "The container '%s' should not exist, but status is '%s'", container, status)
		})
	}
}

func testDestroyCmd(t *testing.T, tc testCase) {
	base.ChangeDir(t, tc.config.ProjectPath)

	// create the containernames
	expected := make([]string, len(tc.config.containers))
	for i, name := range tc.config.containers {
		expected[i] = config.ProjectConfig.ProjectName + "-" + name
	}

	projectName := config.ProjectConfig.ProjectName

	rootCmd := cmd.RootCmd
	base.SilentCommand(t, rootCmd, []string{"stack", "destroy"})

	for _, container := range expected {
		container := container
		t.Run("Check_Container_Removed_"+container, func(t *testing.T) {
			status := docker.GetContainerStatus(container)
			assert.Empty(t, status, "Container '%s' should be removed, but still exists with status '%s'", container, status)
		})
	}

	t.Run("Check_Networks_Removed", func(t *testing.T) {
		networks := docker.GetNetworksByProject(projectName)
		assert.Empty(t, networks, "All networks for project '%s' should be removed", projectName)
	})

	t.Run("Check_Volumes_Removed", func(t *testing.T) {
		volumes := docker.GetVolumesByProject(projectName)
		assert.Empty(t, volumes, "All volumes for project '%s' should be removed", projectName)
	})

	t.Run("Check_Images_Removed", func(t *testing.T) {
		images := docker.GetImagesByProject(projectName)
		assert.Empty(t, images, "All images for project '%s' should be removed", projectName)
	})
}
