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
		composerSingle(c.OutOrStdout(), args)
	},
}

func init() {
	cmd.RootCmd.AddCommand(composerCmd)
}

func composerSingle(out io.Writer, args []string) {

	basecmd.ExecuteBackendCommand(out, "composer", args)
}
