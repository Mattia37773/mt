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
		shellDb(c.OutOrStdout(), args)
	},
}

func init() {
	DbCmd.AddCommand(shellCmd)
}

func shellDb(out io.Writer, args []string) {
	var db config.DbConfig = config.GetDbConfig()

	cmd := basecmd.ExecuteDbCommand(out, db.Shell, args)
	basecmd.ExecuteCommand(out, cmd)
}
