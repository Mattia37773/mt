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

var psCmd = &cobra.Command{
	Use:                   "ps",
	Short:                 "Show stack status",
	GroupID:               "stack",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		psStack()
	},
}

func init() {
	StackCmd.AddCommand(psCmd)
}

func psStack() {
	var projectName string = config.GetProjectName()
	var dockerPath string = config.DockerPath()

	fmt.Printf(ui.Green("Project %s \n"), projectName)
	fmt.Println("")
	fmt.Println("Stack Status")

	cmd := exec.Command("docker", "compose", "-f", dockerPath, "ps")
	shell.ExecuteCommand(cmd)
}
