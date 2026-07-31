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
	RunE: func(c *cobra.Command, args []string) error {
		out := c.OutOrStdout()

		cmd, err := restartGen(out)
		if err != nil {
			return err
		}

		basecmd.StackExecute(out, cmd, "Restarted the stack")
		return nil
	},
}

func init() {
	StackCmd.AddCommand(restartCmd)
}

func restartGen(out io.Writer) ([]string, error) {
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
	fmt.Fprintln(out, "Restart the docker stack")

	execArgs := fmt.Sprintf("docker compose -f %s restart", dockerPath)
	return []string{"sh", "-c", execArgs}, nil
}
