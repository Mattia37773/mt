/*
Copyright © 2026 Matze
*/
package php

import (
	"io"

	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/helper/basecmd"

	"github.com/spf13/cobra"
)

var consoleCmd = &cobra.Command{
	Use:                   "console [args]",
	Short:                 "Run the Symfony console",
	GroupID:               "php",
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	RunE: func(c *cobra.Command, args []string) error {
		out := c.OutOrStdout()

		containerErr := basecmd.CheckContainerExits(out, config.ProjectConfig.ProjectName+"-"+config.ProjectConfig.Backend.ContainerName)
		if containerErr != nil {
			return containerErr
		}

		cmd, err := consoleGen(out, args)
		if err != nil {
			return err
		}

		basecmd.ExecuteCommand(out, cmd)
		return nil
	},
}

func init() {
	PhpCmd.AddCommand(consoleCmd)
}

func consoleGen(out io.Writer, args []string) ([]string, error) {
	return basecmd.ExecuteBackendCommand(out, "bin/console", args)
}
