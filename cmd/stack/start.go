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

var StartCmd = &cobra.Command{
	Use:                   "start",
	Short:                 "Start the docker stack",
	GroupID:               "stack",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {

		startStack()
	},
}

func init() {
	StackCmd.AddCommand(StartCmd)
}

func startStack() {
	var projectName string = config.GetProjectName()
	var dockerPath string = config.DockerPath()

	fmt.Printf(ui.Green("Project %s \n"), projectName)
	fmt.Println("")
	fmt.Println("Starting the docker stack")

	cmd := exec.Command("docker", "compose", "-f", dockerPath, "up", "-d")
	shell.ExecuteCommand(cmd)

	fmt.Println(ui.Green("Successfully started the stack"))
}
