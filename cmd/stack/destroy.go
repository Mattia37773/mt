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

var destroyCmd = &cobra.Command{
	Use:                   "destroy",
	Short:                 "A brief description of your command",
	GroupID:               "stack",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		destroyStack()
	},
}

func init() {
	StackCmd.AddCommand(destroyCmd)
}

func destroyStack() {
	var projectName string = config.GetProjectName()
	var dockerPath string = config.DockerPath()

	fmt.Printf(ui.Green("Project %s \n"), projectName)
	fmt.Println("")
	fmt.Println("DESTROY all project specific IMAGES, VOLUMES and NETWORKS")

	cmd := exec.Command("docker", "compose", "-f", dockerPath, "down", "-v", "--rmi", "all")
	shell.ExecuteCommand(cmd)

	fmt.Println(ui.Green("Successfully destroyed the stack"))
}
