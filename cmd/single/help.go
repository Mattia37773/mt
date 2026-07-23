/*
Copyright © 2026 Matze
*/
package single

import (
	"fmt"
	"strings"

	"github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/ui/text"

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
			out := c.OutOrStdout()
			if len(args) > 0 {
				foundCmd, _, err := cmd.RootCmd.Find(args)
				if err != nil || foundCmd == nil {
					fmt.Fprint(out, text.Red("Error: "))
					fmt.Fprintf(out, "unknown command \"%s\" for \"mt\"", args[0])
					fmt.Fprintln(out)
					fmt.Fprintln(out, text.GlowPink("Try mt --help"))
					return
				}
				// foundCmd.SetOut(out)
				// foundCmd.SetErr(out)
				foundCmd.Help()
				return
			}
		},
	})
}

func helpStyle() {
	cmd.RootCmd.SetHelpFunc(func(c *cobra.Command, args []string) {
		out := c.OutOrStdout()

		fmt.Fprintln(out, text.GlowPink("Usage: ")+text.GlowPurple(c.UseLine()))

		visibleCommands := 0
		for _, sub := range c.Commands() {
			if !sub.Hidden {
				visibleCommands++
			}
		}

		if c.Short != "" {
			fmt.Fprintln(out)
			fmt.Fprintln(out, text.GlowPink(c.Short))
		}

		if visibleCommands > 0 {
			fmt.Fprintln(out)
			fmt.Fprintln(out, text.GlowPink(text.Bold("Commands:")))

			maxLen := 0
			for _, sub := range c.Commands() {
				if !sub.Hidden {
					if len(sub.Name()) > maxLen {
						maxLen = len(sub.Name())
					}
				}
			}

			for _, sub := range c.Commands() {
				if !sub.Hidden {
					name := text.Purple(sub.Name())
					padding := strings.Repeat(" ", maxLen-len(sub.Name()))

					fmt.Fprintf(out, "  %s%s  %s\n",
						name,
						padding,
						sub.Short,
					)
				}
			}
		}

		if len(c.Flags().FlagUsages()) > 0 {
			fmt.Fprintln(out)
			fmt.Fprintln(out, text.GlowPink("Flags:"))
			fmt.Fprint(out, text.Lime(c.Flags().FlagUsages()))
		}
	})
}
