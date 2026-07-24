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

var BuildCmd = &cobra.Command{
	Use:                   "build",
	Short:                 "Build the Docker stack",
	GroupID:               "stack",
	DisableFlagsInUseLine: true,
	Run: func(c *cobra.Command, args []string) {
		out := c.OutOrStdout()

		cmd := buildGen(out)
		basecmd.StackExecute(out, cmd, "Successfully built the stack")
	},
}

func init() {
	StackCmd.AddCommand(BuildCmd)
}

func buildGen(out io.Writer) []string {
	var projectName string = basecmd.ValidateProjectName(out)
	var dockerPath string = basecmd.ValidateDockerPath(out)

	fmt.Fprintf(out, text.Green("Project %s \n"), projectName)
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "Building the docker stack")

	execArgs := fmt.Sprintf("docker compose -f %s build --no-cache", dockerPath)
	return []string{"sh", "-c", execArgs}
}
