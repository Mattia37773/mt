/*
Copyright © 2026 Matze
*/
package single

import (
	"io"

	"github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/helper/basecmd"
	"github.com/spf13/cobra"
)

var composerCmd = &cobra.Command{
	Use:                   "composer [args]",
	Short:                 "Run Composer inside in a container",
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	RunE: func(c *cobra.Command, args []string) error {
		out := c.OutOrStdout()

		containerErr := basecmd.CheckContainerExits(out, config.ProjectConfig.ProjectName+"-"+config.ProjectConfig.Backend.ContainerName)
		if containerErr != nil {
			return containerErr
		}

		cmd, err := composerGen(out, args)
		if err != nil {
			return err
		}

		basecmd.ExecuteCommand(out, cmd)
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(composerCmd)
}

func composerGen(out io.Writer, args []string) ([]string, error) {
	return basecmd.ExecuteBackendCommand(out, "composer", args)
}
