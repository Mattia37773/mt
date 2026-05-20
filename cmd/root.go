/*
Copyright © 2026 Matze
*/
package cmd

import (
	"os"
	"strings"

	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/functions/sys"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:           "mt [options] [command]",
	Short:         "A tool for managing projects",
	Long:          `This is a tool to manage local project.`,
	SilenceErrors: true,
	SilenceUsage:  true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		sys.Error(err)
		os.Exit(1)
	}

}

func init() {
	sys.VersionStyle(RootCmd)
	RootCmd.Version = config.Version

	// When the command is update
	// the message shouldn't show up
	if len(os.Args) == 2 {
		arg := os.Args[1:]
		argName := strings.Join(arg, "")
		if argName != "update" {
			sys.ShowUpdateMessage(config.Version, config.GithubBaseApi)
		}
	} else {
		sys.ShowUpdateMessage(config.Version, config.GithubBaseApi)
	}
}
