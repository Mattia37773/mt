/*
Copyright © 2026 Matze
*/
package php

import (
	"io"

	"github.com/mattia37773/mt/helper/basecmd"

	"github.com/spf13/cobra"
)

var consoleCmd = &cobra.Command{
	Use:                   "console [args]",
	Short:                 "Run the Symfony console",
	GroupID:               "php",
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	Run: func(c *cobra.Command, args []string) {
		consoleSymfony(c.OutOrStdout(), args)
	},
}

func init() {
	PhpCmd.AddCommand(consoleCmd)
}

func consoleSymfony(out io.Writer, args []string) {
	basecmd.ExecuteBackendCommand(out, "bin/console", args)
}
