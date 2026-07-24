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

var psCmd = &cobra.Command{
	Use:                   "ps",
	Short:                 "Show stack status",
	GroupID:               "stack",
	DisableFlagsInUseLine: true,
	Run: func(c *cobra.Command, args []string) {
		out := c.OutOrStdout()

		cmd := psGen(out)
		basecmd.StackExecute(out, cmd, "")
	},
}

func init() {
	StackCmd.AddCommand(psCmd)
}

func psGen(out io.Writer) []string {
	var projectName string = basecmd.ValidateProjectName(out)
	var dockerPath string = basecmd.ValidateDockerPath(out)

	fmt.Fprintf(out, text.Green("Project %s \n"), projectName)
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "Stack Status")

	execArgs := fmt.Sprintf("docker compose -f %s ps", dockerPath)

	return []string{"sh", "-c", execArgs}
}
