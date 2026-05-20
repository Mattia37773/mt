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

var restartCmd = &cobra.Command{
	Use:                   "restart",
	Short:                 "Restart the Docker Stack",
	GroupID:               "stack",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		restartStack()
	},
}

func init() {
	StackCmd.AddCommand(restartCmd)
}

func restartStack() {
	var projectName string = config.GetProjectName()
	var dockerPath string = config.DockerPath()

	fmt.Printf(ui.Green("Project %s \n"), projectName)
	fmt.Println("")
	fmt.Println("Restart the docker stack")

	cmd := exec.Command("docker", "compose", "-f", dockerPath, "restart")
	shell.ExecuteCommand(cmd)

	fmt.Println(ui.Green("Restarted the stack"))
}
