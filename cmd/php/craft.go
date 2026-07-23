/*
Copyright © 2026 Matze
*/
package php

import (
	"io"

	"github.com/mattia37773/mt/helper/basecmd"
	"github.com/spf13/cobra"
)

var craftCmd = &cobra.Command{
	Use:                   "craft [args]",
	Short:                 "Run the Craft CMS cli",
	GroupID:               "php",
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	Run: func(c *cobra.Command, args []string) {
		craft(c.OutOrStdout(), args)
	},
}

func init() {
	PhpCmd.AddCommand(craftCmd)
}

func craft(out io.Writer, args []string) {
	basecmd.ExecuteBackendCommand(out, "php craft", args)
}
