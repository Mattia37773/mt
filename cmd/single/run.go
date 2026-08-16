/*
Copyright © 2026 Matze
*/
package single

import (
	"io"

	"github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/helper/basecmd"
	"github.com/mattia37773/mt/helper/docker"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:                   "run [commands] [args]",
	Short:                 "Run a Command inside the Backend Container",
	Aliases:               []string{"exec"},
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	RunE: func(c *cobra.Command, args []string) error {
		out := c.OutOrStdout()

		projectName, projectNameErr := basecmd.ValidateProjectName(out)
		if projectNameErr != nil {
			return projectNameErr
		}

		_, containerErr := docker.ContainerExits(out, projectName+"-"+config.ProjectConfig.Main.ContainerName)
		if containerErr != nil {
			return containerErr
		}

		cmd, err := runGen(out, args)
		if err != nil {
			return err
		}

		basecmd.ExecuteCommand(out, cmd)
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(runCmd)
}

func runGen(out io.Writer, args []string) ([]string, error) {
	return basecmd.ExecuteMainCommand(out, args[0], args[1:])
}
