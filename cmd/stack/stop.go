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

var stopCmd = &cobra.Command{
	Use:                   "stop",
	Short:                 "Stop the local stack",
	GroupID:               "stack",
	DisableFlagsInUseLine: true,
	Run: func(c *cobra.Command, args []string) {
		out := c.OutOrStdout()

		cmd := stopGen(out)
		basecmd.StackExecute(out, cmd, "Stack stopped")
	},
}

func init() {
	StackCmd.AddCommand(stopCmd)
}

func stopGen(out io.Writer) []string {
	var projectName string = basecmd.ValidateProjectName(out)
	var dockerPath string = basecmd.ValidateDockerPath(out)

	fmt.Fprintf(out, text.Green("Project %s \n"), projectName)
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "Stopping the docker stack")

	execArgs := fmt.Sprintf("docker compose -f %s down", dockerPath)
	return []string{"sh", "-c", execArgs}
}
