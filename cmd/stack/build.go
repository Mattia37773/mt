/*
Copyright © 2026 Matze
*/
package stack

import (
	"fmt"
	"os/exec"

	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/functions/shell"
	"github.com/mattia37773/mt/functions/ui"
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
	var projectName string = config.GetProjectName()
	var dockerPath string = config.DockerPath()

	fmt.Printf(ui.Green("Project %s \n"), projectName)
	fmt.Println("")
	fmt.Println("Building the docker stack")

	cmd := exec.Command("docker", "compose", "-f", dockerPath, "build", "--no-cache")
	shell.ExecuteCommand(cmd)

	fmt.Println(ui.Green("Successfully built the stack"))
}
