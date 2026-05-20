/*
Copyright © 2026 Matze
*/
package stack

import (
	"fmt"

	"github.com/mattia37773/mt/cmd"

	"github.com/spf13/cobra"
)

var StackCmd = &cobra.Command{
	Use:                   "stack [command]",
	Short:                 "Manage the local docker stack",
	Long:                  "This command helps you to manage the local Docker stack",
	DisableFlagsInUseLine: false,
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
	cmd.RootCmd.AddCommand(StackCmd)

	StackCmd.AddGroup(&cobra.Group{
		ID:    "stack",
		Title: "Commands",
	})
}
