/*
Copyright © 2026 Matze
*/
package stack

import (
	"fmt"
	"io"
	"strings"

	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/helper/basecmd"
	"github.com/mattia37773/mt/ui/text"

	"github.com/spf13/cobra"
)

var followLogs bool

var logsCmd = &cobra.Command{
	Use:                   "logs [flags] [service]",
	Short:                 "Show container logs",
	GroupID:               "stack",
	DisableFlagsInUseLine: true,
	RunE: func(c *cobra.Command, args []string) error {
		var service string
		if len(args) > 0 {
			service = args[0]
		}

		cmd, err := logsStack(c.OutOrStdout(), followLogs, service)
		if err != nil {
			return err
		}
		basecmd.StackExecute(c.OutOrStdout(), cmd, "")
		return nil
	},
}

func init() {
	StackCmd.AddCommand(logsCmd)

	logsCmd.Flags().BoolVarP(&followLogs, "follow", "f", false, "Follow log output")
}

func logsStack(out io.Writer, follow bool, service string) ([]string, error) {

	projectName, err := basecmd.ValidateProjectName(out)
	if err != nil {
		return []string{}, err
	}

	dockerPath, err := basecmd.ValidateDockerPath(out)
	if err != nil {
		return []string{}, err
	}

	basecmd.CheckContainerExits(out, projectName+"-"+config.ProjectConfig.Main.ContainerName)

	fmt.Fprintf(out, text.Green("Project %s \n"), projectName)
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "Printing Logs")

	execArgs := []string{"docker", "compose", "-f", dockerPath, "logs"}

	if follow {
		execArgs = append(execArgs, "-f")
	}

	if service != "" {
		execArgs = append(execArgs, service)
	}
	fmt.Println([]string{"sh", "-c", strings.Join(execArgs, " ")})

	return []string{"sh", "-c", strings.Join(execArgs, " ")}, nil
}
