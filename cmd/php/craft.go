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

var craftCmd = &cobra.Command{
	Use:                   "craft [args]",
	Short:                 "Run the Craft CMS cli",
	GroupID:               "php",
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	RunE: func(c *cobra.Command, args []string) error {
		out := c.OutOrStdout()

		containerErr := basecmd.CheckContainerExits(out, config.ProjectConfig.ProjectName+"-"+config.ProjectConfig.Backend.ContainerName)
		if containerErr != nil {
			return containerErr
		}

		cmd, err := craftGen(out, args)
		if err != nil {
			return err
		}

		basecmd.ExecuteCommand(out, cmd)
		return nil
	},
}

func init() {
	PhpCmd.AddCommand(craftCmd)
}

func craftGen(out io.Writer, args []string) ([]string, error) {
	return basecmd.ExecuteBackendCommand(out, "php craft", args)
}
