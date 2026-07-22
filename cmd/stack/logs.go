/*
Copyright © 2026 Matze
*/
package stack

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/helper/shell"
	"github.com/mattia37773/mt/ui/text"

	"github.com/spf13/cobra"
)

var followLogs bool

var logsCmd = &cobra.Command{
	Use:                   "logs [service]",
	Short:                 "Show container logs",
	GroupID:               "stack",
	DisableFlagsInUseLine: true,
	Run: func(c *cobra.Command, args []string) {
		var service string
		if len(args) > 0 {
			service = args[0]
		}

		logsStack(c.OutOrStdout(), followLogs, service)
	},
}

func init() {
	StackCmd.AddCommand(logsCmd)

	logsCmd.Flags().BoolVarP(&followLogs, "follow", "f", false, "Follow log output")
}

func logsStack(out io.Writer, follow bool, service string) {
	var projectName string = config.ProjectConfig.ProjectName
	var dockerPath string = config.ProjectConfig.Paths.DockerCompose

	fmt.Fprintf(out, text.Green("Project %s \n"), projectName)
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "Printing Logs")

	cmdArgs := []string{"compose", "-f", dockerPath, "logs"}

	if follow {
		cmdArgs = append(cmdArgs, "-f")
	}

	if service != "" {
		cmdArgs = append(cmdArgs, service)
	}

	cmd := exec.Command("docker", cmdArgs...)
	cmd.Stdout = out
	cmd.Stderr = out

	err := shell.ExecuteCommand(cmd)
	if err != nil {
		fmt.Fprint(out, text.Red("Something went wrong: "))
		fmt.Fprintln(out, err)
		os.Exit(1)
	}
}
