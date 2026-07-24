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

var destroyCmd = &cobra.Command{
	Use:                   "destroy",
	Short:                 "Destroy the docker stack",
	GroupID:               "stack",
	DisableFlagsInUseLine: true,
	Run: func(c *cobra.Command, args []string) {
		out := c.OutOrStdout()

		cmd := destroyGen(out)
		basecmd.StackExecute(out, cmd, "Successfully destroyed the stack")
	},
}

func init() {
	StackCmd.AddCommand(destroyCmd)
}

func destroyGen(out io.Writer) []string {
	var projectName string = basecmd.ValidateProjectName(out)
	var dockerPath string = basecmd.ValidateDockerPath(out)

	fmt.Fprintf(out, text.Green("Project %s \n"), projectName)
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "DESTROY all project specific IMAGES, VOLUMES and NETWORKS")

	execArgs := fmt.Sprintf("docker compose -f %s down -v --rmi all", dockerPath)
	return []string{"sh", "-c", execArgs}
}
