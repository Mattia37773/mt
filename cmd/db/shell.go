/*
Copyright © 2026 Matze
*/
package db

import (
	"io"

	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/helper/basecmd"

	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:                   "shell",
	Short:                 "Connect to the database",
	GroupID:               "db",
	DisableFlagsInUseLine: true,
	Run: func(c *cobra.Command, args []string) {
		out := c.OutOrStdout()

		cmd := shellGen(out, args)
		basecmd.ExecuteCommand(out, cmd)
	},
}

func init() {
	DbCmd.AddCommand(shellCmd)
}

func shellGen(out io.Writer, args []string) []string {
	return basecmd.ExecuteDbCommand(out, config.GetDbConfig().Shell, args)
}
