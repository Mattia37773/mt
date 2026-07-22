/*
Copyright © 2026 Matze
*/
package stack

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/helper/shell"
	"github.com/mattia37773/mt/ui/text"

	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:                   "destroy",
	Short:                 "Destroy the docker stack",
	GroupID:               "stack",
	DisableFlagsInUseLine: true,
	Run: func(c *cobra.Command, args []string) {
		destroyStack(c.OutOrStdout())
	},
}

func init() {
	StackCmd.AddCommand(destroyCmd)
}

func destroyStack(out io.Writer) {
	var projectName string = config.ProjectConfig.ProjectName
	var dockerPath string = config.ProjectConfig.Paths.DockerCompose

	fmt.Fprintf(out, text.Green("Project %s \n"), projectName)
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "DESTROY all project specific IMAGES, VOLUMES and NETWORKS")

	cmd := exec.Command("docker", "compose", "-f", dockerPath, "down", "-v", "--rmi", "all")
	err := shell.ExecuteCommand(cmd)
	if err != nil {
		fmt.Fprint(out, text.Red("Something went wrong: "))
		fmt.Fprintln(out, err)
		os.Exit(1)
	}

	fmt.Fprintln(out, text.Green("Successfully destroyed the stack"))
}
