/*
Copyright © 2026 Matze
*/
package single

import (
	"io"

	"github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/helper/basecmd"
	"github.com/spf13/cobra"
)

var composerCmd = &cobra.Command{
	Use:                   "composer [args]",
	Short:                 "Run Composer inside in a container",
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	Run: func(c *cobra.Command, args []string) {
		out := c.OutOrStdout()

		cmd := composerGen(out, args)
		basecmd.ExecuteCommand(out, cmd)
	},
}

func init() {
	cmd.RootCmd.AddCommand(composerCmd)
}

func composerGen(out io.Writer, args []string) []string {
	cmd := basecmd.ExecuteBackendCommand(out, "composer", args)

	return cmd
}
