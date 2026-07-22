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

var StartCmd = &cobra.Command{
	Use:                   "start",
	Short:                 "Start the docker stack",
	GroupID:               "stack",
	DisableFlagsInUseLine: true,
	Run: func(c *cobra.Command, args []string) {
		startStack(c.OutOrStdout())
	},
}

func init() {
	StackCmd.AddCommand(StartCmd)
}

func startStack(out io.Writer) {
	// todo add valdation for those two
	var projectName string = config.ProjectConfig.ProjectName
	var dockerPath string = config.ProjectConfig.Paths.DockerCompose

	fmt.Fprintf(out, text.Green("Project %s \n"), projectName)
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "Starting the docker stack")

	cmd := exec.Command("docker", "compose", "-f", dockerPath, "up", "-d")
	err := shell.ExecuteCommand(cmd)
	if err != nil {
		fmt.Fprint(out, text.Red("Something went wrong: "))
		fmt.Fprintln(out, err)
		os.Exit(1)
	}

	fmt.Fprintln(out, text.Green("Successfully started the stack"))
}
