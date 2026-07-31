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
	RunE: func(c *cobra.Command, args []string) error {
		out := c.OutOrStdout()

		cmd, err := destroyGen(out)
		if err != nil {
			return err
		}
		basecmd.StackExecute(out, cmd, "Successfully destroyed the stack")
		return nil
	},
}

func init() {
	StackCmd.AddCommand(destroyCmd)
}

func destroyGen(out io.Writer) ([]string, error) {
	projectName, err := basecmd.ValidateProjectName(out)
	if err != nil {
		return []string{}, err
	}

	dockerPath, err := basecmd.ValidateDockerPath(out)
	if err != nil {
		return []string{}, err
	}

	fmt.Fprintf(out, text.Green("Project %s \n"), projectName)
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "DESTROY all project specific IMAGES, VOLUMES and NETWORKS")

	execArgs := fmt.Sprintf("docker compose -f %s down -v --rmi all", dockerPath)
	return []string{"sh", "-c", execArgs}, nil
}
