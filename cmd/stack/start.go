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
	RunE: func(c *cobra.Command, args []string) error {
		out := c.OutOrStdout()

		cmd, err := startGen(out)
		if err != nil {
			return err
		}

		err = basecmd.StackExecute(out, cmd, "Successfully started the stack")
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	StackCmd.AddCommand(StartCmd)
}

func startGen(out io.Writer) ([]string, error) {
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
	fmt.Fprintln(out, "Starting the docker stack")

	execArgs := fmt.Sprintf("docker compose -f %s up -d", dockerPath)
	return []string{"sh", "-c", execArgs}, nil
}
