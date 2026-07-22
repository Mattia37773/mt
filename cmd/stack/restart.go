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

var restartCmd = &cobra.Command{
	Use:                   "restart",
	Short:                 "Restart the Docker Stack",
	GroupID:               "stack",
	DisableFlagsInUseLine: true,
	Run: func(c *cobra.Command, args []string) {
		restartStack(c.OutOrStdout())
	},
}

func init() {
	StackCmd.AddCommand(restartCmd)
}

func restartStack(out io.Writer) {
	// todo add valdation for those two
	var projectName string = config.ProjectConfig.ProjectName
	var dockerPath string = config.ProjectConfig.Paths.DockerCompose

	fmt.Fprintf(out, text.Green("Project %s \n"), projectName)
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "Restart the docker stack")

	cmd := exec.Command("docker", "compose", "-f", dockerPath, "restart")
	err := shell.ExecuteCommand(cmd)
	if err != nil {
		fmt.Fprint(out, text.Red("Something went wrong: "))
		fmt.Fprintln(out, err)
		os.Exit(1)
	}

	fmt.Fprintln(out, text.Green("Restarted the stack"))
}
