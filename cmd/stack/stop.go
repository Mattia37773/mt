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

var stopCmd = &cobra.Command{
	Use:                   "stop",
	Short:                 "Stop the local stack",
	GroupID:               "stack",
	DisableFlagsInUseLine: true,
	Run: func(c *cobra.Command, args []string) {
		stopStack(c.OutOrStdout())
	},
}

func init() {
	StackCmd.AddCommand(stopCmd)
}

func stopStack(out io.Writer) {
	// todo add valdation for those two
	var projectName string = config.ProjectConfig.ProjectName
	var dockerPath string = config.ProjectConfig.Paths.DockerCompose

	fmt.Fprintf(out, text.Green("Project %s \n"), projectName)
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "Stopping the docker stack")

	cmd := exec.Command("docker", "compose", "-f", dockerPath, "down")
	err := shell.ExecuteCommand(cmd)
	if err != nil {
		fmt.Fprint(out, text.Red("Something went wrong: "))
		fmt.Fprintln(out, err)
		os.Exit(1)
	}

	fmt.Fprintln(out, text.Green("Stack stopped"))
}
