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

var stopCmd = &cobra.Command{
	Use:                   "stop",
	Short:                 "Stop the local stack",
	GroupID:               "stack",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		stopStack()
	},
}

func init() {
	StackCmd.AddCommand(stopCmd)
}

func stopStack() {
	var projectName string = config.GetProjectName()
	var dockerPath string = config.DockerPath()

	fmt.Printf(ui.Green("Project %s \n"), projectName)
	fmt.Println("")
	fmt.Println("Stopping the docker stack")

	cmd := exec.Command("docker", "compose", "-f", dockerPath, "down")
	shell.ExecuteCommand(cmd)

	fmt.Println(ui.Green("Stack stopped"))
}
