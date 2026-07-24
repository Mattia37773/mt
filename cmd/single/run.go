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

var runCmd = &cobra.Command{
	Use:                   "run [commands] [args]",
	Short:                 "Run a Command inside the Backend Container",
	Aliases:               []string{"exec"},
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	Run: func(c *cobra.Command, args []string) {
		out := c.OutOrStdout()

		cmd := runGen(out, args)
		basecmd.ExecuteCommand(out, cmd)
	},
}

func init() {
	cmd.RootCmd.AddCommand(runCmd)
}

func runGen(out io.Writer, args []string) []string {
	return basecmd.ExecuteMainCommand(out, "", args)
}
