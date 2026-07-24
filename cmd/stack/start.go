/*
Copyright © 2026 Matze
*/
package stack

import (
	"fmt"
	"io"

	"github.com/mattia37773/mt/helper/basecmd"
	"github.com/mattia37773/mt/ui/text"
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:                   "start",
	Short:                 "Start the docker stack",
	GroupID:               "stack",
	DisableFlagsInUseLine: true,
	Run: func(c *cobra.Command, args []string) {
		out := c.OutOrStdout()

		cmd := startGen(out)
		basecmd.StackExecute(out, cmd, "Successfully started the stack")
	},
}

func init() {
	StackCmd.AddCommand(StartCmd)
}

func startGen(out io.Writer) []string {
	var projectName string = basecmd.ValidateProjectName(out)
	var dockerPath string = basecmd.ValidateDockerPath(out)

	fmt.Fprintf(out, text.Green("Project %s \n"), projectName)
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "Starting the docker stack")

	execArgs := fmt.Sprintf("docker compose -f %s up -d", dockerPath)
	return []string{"sh", "-c", execArgs}
}
