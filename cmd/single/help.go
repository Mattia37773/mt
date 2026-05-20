/*
Copyright © 2026 Matze
*/
package single

import (
	"fmt"
	"strings"

	"github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/functions/ui"

	"github.com/spf13/cobra"
)

func init() {
	helpCommandNotFound()
	helpStyle()
}

func helpCommandNotFound() {
	cmd.RootCmd.SetHelpCommand(&cobra.Command{
		Use:   "help [command]",
		Short: "Shows the help text for an command",
		Run: func(c *cobra.Command, args []string) {
			if len(args) > 0 {
				foundCmd, _, err := cmd.RootCmd.Find(args)
				if err != nil || foundCmd == nil {
					fmt.Printf(ui.Red("Error: "))
					fmt.Printf("unknown command \"%s\" for \"mt\"", args[0])
					fmt.Println()
					fmt.Println(ui.GlowPink("Try mt --help"))
					return
				}
				foundCmd.Help()
				return
			}
		},
	})
}

func helpStyle() {
	cmd.RootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Printf(ui.GlowPink("Usage: "))
		fmt.Println(ui.GlowPurple(cmd.UseLine()))

		visibleCommands := 0
		for _, c := range cmd.Commands() {
			if !c.Hidden {
				visibleCommands++
			}
		}

		if cmd.Short != "" {
			fmt.Println()
			fmt.Println(ui.GlowPink(cmd.Short))
		}

		if visibleCommands > 0 {
			fmt.Println()
			fmt.Println(ui.GlowPink(ui.Bold("Commands:")))

			maxLen := 0
			for _, c := range cmd.Commands() {
				if !c.Hidden {
					if len(c.Name()) > maxLen {
						maxLen = len(c.Name())
					}
				}
			}

			for _, c := range cmd.Commands() {
				if !c.Hidden {
					name := ui.Purple(c.Name())
					padding := strings.Repeat(" ", maxLen-len(c.Name()))

					fmt.Printf("  %s%s  %s\n",
						name,
						padding,
						c.Short,
					)
				}
			}
		}

		if len(cmd.Flags().FlagUsages()) > 0 {
			fmt.Println()
			fmt.Println(ui.GlowPink("Flags:"))
			fmt.Print(ui.Lime(cmd.Flags().FlagUsages()))
		}
	})
}
