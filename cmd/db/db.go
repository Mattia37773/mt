/*
Copyright © 2026 Matze
*/
package db

import (
	"fmt"

	"github.com/mattia37773/mt/cmd"

	"github.com/spf13/cobra"
)

var DbCmd = &cobra.Command{
	Use:                   "db [command] [flags]",
	Short:                 "Manage the local database in the db container",
	DisableFlagsInUseLine: true,
	SilenceErrors:         false,
	SilenceUsage:          true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help()
		}
		return fmt.Errorf("No such command %s ", args)
	},
}

func init() {
	cmd.RootCmd.AddCommand(DbCmd)

	DbCmd.AddGroup(&cobra.Group{
		ID:    "db",
		Title: "Commands",
	})
}
