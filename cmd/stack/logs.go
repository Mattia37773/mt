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

var logsCmd = &cobra.Command{
	Use:                   "logs",
	Short:                 "Show container logs",
	GroupID:               "stack",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		logsStack()
	},
}

func init() {
	StackCmd.AddCommand(logsCmd)
}

func logsStack() {
	var projectName string = config.GetProjectName()
	var dockerPath string = config.DockerPath()

	fmt.Printf(ui.Green("Project %s \n"), projectName)
	fmt.Println("")
	fmt.Println("Printing Logs")

	cmd := exec.Command("docker", "compose", "-f", dockerPath, "logs", "-f")
	shell.ExecuteCommand(cmd)
}
