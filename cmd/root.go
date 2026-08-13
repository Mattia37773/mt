/*
Copyright © 2026 Matze
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/ui"
	"github.com/mattia37773/mt/ui/text"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:           "mt [options] [command]",
	Short:         "A tool for managing projects",
	Long:          `This is a tool to manage local project.`,
	SilenceErrors: true,
	SilenceUsage:  true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// parse the config file
		err := config.ParseConfigFile()
		if err != nil {
			return err
		}

		// When the command is update
		// the message shouldn't show up
		if len(os.Args) == 2 {
			arg := os.Args[1:]
			argName := strings.Join(arg, "")
			if argName != "update" {
				showUpdateMessage(cmd.OutOrStdout())
			}
		} else {
			showUpdateMessage(cmd.OutOrStdout())
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// error message for a not found command
	err := RootCmd.Execute()
	if err != nil {
		errorMessage(err)
		os.Exit(1)
	}

}

func init() {
	versionStyle(RootCmd)
	RootCmd.Version = config.Version
	//RootCmd.CompletionOptions.DisableDefaultCmd = true
}

func versionStyle(cmd *cobra.Command) {
	cmd.SetVersionTemplate(
		text.GlowPink("{{.Name}}") +
			" " +
			text.GlowPurple("version "+config.Version) +
			"\n",
	)
}

func errorMessage(err error) {
	errorText := strings.Split(err.Error(), "\n")

	fmt.Println(text.Red("Error:"), errorText[0])

	for _, line := range errorText[1:] {
		fmt.Println(line)
	}

	fmt.Println()
	fmt.Println(text.GlowPink("Try mt --help for usage."))
	os.Exit(1)
}

func showUpdateMessage(out io.Writer) {
	latestVersion := config.LatestVersion
	version := config.Version

	if version < latestVersion {
		if version != "dev" {
			lines := []string{
				"A new update is available",
				"Current Version: " + version,
				"Latest  Version: " + latestVersion,
			}
			ui.Border(out, lines)
		}
	}
}
