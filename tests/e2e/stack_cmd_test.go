package e2e

import (
	"bytes"
	"os/exec"
	"testing"
	"time"

	"github.com/mattia37773/mt/cmd"
	_ "github.com/mattia37773/mt/cmd/php"
	_ "github.com/mattia37773/mt/cmd/single"
	_ "github.com/mattia37773/mt/cmd/stack"
	"github.com/mattia37773/mt/config"
	base "github.com/mattia37773/mt/helper/basetest"
	"github.com/mattia37773/mt/helper/docker"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	// t.Run("Build", func(t *testing.T) {
	// 	testBuildCmd(t)
	// })

	// t.Run("Start", func(t *testing.T) {
	// 	testStartCmd(t)
	// })

	// t.Run("Restart", func(t *testing.T) {
	// 	testRestartCmd(t)
	// })

	// t.Run("Ps", func(t *testing.T) {
	// 	testPsCmd(t)
	// })

	// t.Run("Logs", func(t *testing.T) {
	// 	testLogsCmd(t)
	// })

	// t.Run("Fpm Logs", func(t *testing.T) {
	// 	testLogsCmdSpecificContainer(t)
	// })

	// t.Run("Stop", func(t *testing.T) {
	// 	testStopCmd(t)
	// })

	// t.Run("StartAgain", func(t *testing.T) {
	// 	testStartCmd(t)
	// })

	// t.Run("Destory", func(t *testing.T) {
	// 	testDestroyCmd(t)
	// })
}

func testBuildCmd(t *testing.T) {
	base.ChangeDirToSymfony(t)

	rootCmd := cmd.RootCmd
	base.SilentCommand(t, rootCmd, []string{"stack", "build"})

	expectedImages := []string{
		config.ProjectConfig.ProjectName + "-fpm:latest",
	}

	for _, img := range expectedImages {
		t.Run("Check_Image_"+img, func(t *testing.T) {
			exists := docker.ImageExists(img)
			assert.True(t, exists, "The dockerimage '%s' but doesn't exist", img)
			// remove image again
			cmd := exec.Command("docker", "image", "rm", "-f", img)
			cmd.Run()
		})
	}

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

	for _, container := range expectedContainers {
		container := container
		t.Run("Check_Container_"+container, func(t *testing.T) {
			running := docker.IsContainerRunning(container)
			assert.True(t, running, "The Container '%s'should run but doesn't", container)
		})
	}
}

func testPsCmd(t *testing.T) {
	base.ChangeDirToSymfony(t)

	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"stack", "ps"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(buf.String())

	assert.Contains(t, cleanOutput, "NAME", "Output should contain column header 'NAME'")
	assert.Contains(t, cleanOutput, "STATUS", "Output should contain column header 'STATUS'")

	expectedContainers := []string{
		config.ProjectConfig.ProjectName + "-db",
		config.ProjectConfig.ProjectName + "-fpm",
		config.ProjectConfig.ProjectName + "-mailpit",
		config.ProjectConfig.ProjectName + "-nginx",
		config.ProjectConfig.ProjectName + "-phpmyadmin",
	}

	for _, container := range expectedContainers {
		assert.Contains(t, cleanOutput, container, "The 'ps' output should contain container '%s'", container)
	}
}

func testLogsCmd(t *testing.T) {
	base.ChangeDirToSymfony(t)

	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"stack", "logs"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(buf.String())

	expectedServices := []string{
		"fpm",
		"db",
		"nginx",
	}

	for _, service := range expectedServices {
		assert.Contains(t, cleanOutput, service, "The logs output should contain entries from '%s'", service)
	}
}

func testLogsCmdSpecificContainer(t *testing.T) {
	base.ChangeDirToSymfony(t)

	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"stack", "logs", "fpm"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(buf.String())

	expectedServices := []string{
		"fpm",
	}

	for _, service := range expectedServices {
		assert.Contains(t, cleanOutput, service, "The logs output should contain entries from '%s'", service)
	}
}

func testRestartCmd(t *testing.T) {
	base.ChangeDirToSymfony(t)

	expectedContainers := []string{
		config.ProjectConfig.ProjectName + "-db",
		config.ProjectConfig.ProjectName + "-fpm",
		config.ProjectConfig.ProjectName + "-mailpit",
		config.ProjectConfig.ProjectName + "-nginx",
		config.ProjectConfig.ProjectName + "-phpmyadmin",
	}

	// 1. Save start timestamps of all containers BEFORE restart
	timestampsBefore := make(map[string]string)
	for _, container := range expectedContainers {
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

	err := rootCmd.Execute()
	assert.NoError(t, err)

	for _, container := range expectedContainers {
		container := container
		t.Run("Check_Restart_"+container, func(t *testing.T) {
			startTimeAfter := docker.GetContainerStartedAt(container)
			assert.NotEmpty(t, startTimeAfter, "Container '%s' should still exist after restart", container)

			// Timestamp must be different/newer than before
			assert.NotEqual(t, timestampsBefore[container], startTimeAfter, "The container '%s' was not restarted", container)
		})
	}
}

func testStopCmd(t *testing.T) {
	base.ChangeDirToSymfony(t)
	t.Log(config.ProjectConfig.ProjectName)

	rootCmd := cmd.RootCmd
	base.SilentCommand(t, rootCmd, []string{"stack", "stop"})

	expectedContainers := []string{
		config.ProjectConfig.ProjectName + "-db",
		config.ProjectConfig.ProjectName + "-fpm",
		config.ProjectConfig.ProjectName + "-mailpit",
		config.ProjectConfig.ProjectName + "-nginx",
		config.ProjectConfig.ProjectName + "-phpmyadmin",
	}

	for _, container := range expectedContainers {
		container := container
		t.Run("Check_Container_DoesNotExist_"+container, func(t *testing.T) {
			status := docker.GetContainerStatus(container)

			assert.Empty(t, status, "The container '%s' should not exist, but status is '%s'", container, status)
		})
	}
}

func testDestroyCmd(t *testing.T) {
	base.ChangeDirToSymfony(t)

	projectName := config.ProjectConfig.ProjectName

	rootCmd := cmd.RootCmd
	base.SilentCommand(t, rootCmd, []string{"stack", "destroy"})

	expectedContainers := []string{
		projectName + "-db",
		projectName + "-fpm",
		projectName + "-mailpit",
		projectName + "-nginx",
		projectName + "-phpmyadmin",
	}

	for _, container := range expectedContainers {
		container := container
		t.Run("Check_Container_Removed_"+container, func(t *testing.T) {
			status := docker.GetContainerStatus(container)
			assert.Empty(t, status, "Container '%s' should be removed, but still exists with status '%s'", container, status)
		})
	}

	t.Run("Check_Networks_Removed", func(t *testing.T) {
		// Sucht nach Docker-Netzwerken, die mit dem Projektnamen matchen
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
