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
		out := c.OutOrStdout()

		cmd := craftGen(out, args)
		basecmd.ExecuteCommand(out, cmd)
	},
}

func init() {
	PhpCmd.AddCommand(craftCmd)
}

func craftGen(out io.Writer, args []string) []string {
	return basecmd.ExecuteBackendCommand(out, "php craft", args)
}
