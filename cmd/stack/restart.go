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

var restartCmd = &cobra.Command{
	Use:                   "restart",
	Short:                 "Restart the Docker Stack",
	GroupID:               "stack",
	DisableFlagsInUseLine: true,
	Run: func(c *cobra.Command, args []string) {
		out := c.OutOrStdout()

		cmd := restartGen(out)
		basecmd.StackExecute(out, cmd, "Restarted the stack")
	},
}

func init() {
	StackCmd.AddCommand(restartCmd)
}

func restartGen(out io.Writer) []string {
	var projectName string = basecmd.ValidateProjectName(out)
	var dockerPath string = basecmd.ValidateDockerPath(out)

	fmt.Fprintf(out, text.Green("Project %s \n"), projectName)
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "Restart the docker stack")

	execArgs := fmt.Sprintf("docker compose -f %s restart", dockerPath)
	return []string{"sh", "-c", execArgs}
}
