/*
Copyright © 2026 Matze
*/
package php

import (
	"fmt"

	"github.com/mattia37773/mt/cmd"

	"github.com/spf13/cobra"
)

var PhpCmd = &cobra.Command{
	Use:                   "php [command] [args]",
	Short:                 "Run some popular PHP tasks",
	Aliases:               []string{"symfony", "laravel"},
	DisableFlagsInUseLine: true,
	SilenceErrors:         false,
	SilenceUsage:          false,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help()
		}
		return fmt.Errorf("No such command %s ", args)
	},
}

func init() {
	cmd.RootCmd.AddCommand(PhpCmd)

	PhpCmd.AddGroup(&cobra.Group{
		ID:    "php",
		Title: "Commands",
	})
}
