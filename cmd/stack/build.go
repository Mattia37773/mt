/*
Copyright © 2026 Matze
*/
package stack

import (
	"fmt"
	"os/exec"

	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/helper/shell"

	"github.com/mattia37773/mt/ui/text"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:                   "build",
	Short:                 "Build the Docker stack",
	GroupID:               "stack",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		buildStack()
	},
}

func init() {
	StackCmd.AddCommand(buildCmd)
}

func buildStack() {
	var projectName string = config.ProjectConfig.ProjectName
	var dockerPath string = config.ProjectConfig.Paths.DockerCompose

	fmt.Printf(text.Green("Project %s \n"), projectName)
	fmt.Println("")
	fmt.Println("Building the docker stack")

	cmd := exec.Command("docker", "compose", "-f", dockerPath, "build", "--no-cache")
	shell.ExecuteCommand(cmd)

	fmt.Println(text.Green("Successfully built the stack"))
}
