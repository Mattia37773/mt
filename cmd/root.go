/*
Copyright © 2026 Matze
*/
package cmd

import (
	"fmt"
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
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// gets the config data into a strucht
	config.ParseConfigFile()
	// error message for a not found command
	err := RootCmd.Execute()
	if err != nil {
		errorMessage(err)
		os.Exit(1)
	}

}

func init() {
	versionStyle(RootCmd)
	RootCmd.Version = config.AppConfig.Version

	// When the command is update
	// the message shouldn't show up
	if len(os.Args) == 2 {
		arg := os.Args[1:]
		argName := strings.Join(arg, "")
		if argName != "update" {
			showUpdateMessage(config.AppConfig.Version)
		}
	} else {
		showUpdateMessage(config.AppConfig.Version)
	}
}

func versionStyle(cmd *cobra.Command) {
	cmd.SetVersionTemplate(
		text.GlowPink("{{.Name}}") +
			" " +
			text.GlowPurple("version {{.Version}}") +
			"\n",
	)
}

func errorMessage(err error) {
	errorText := strings.Split(err.Error(), "\n")
	fmt.Println(text.Red("Error:"), strings.NewReplacer("[", "", "]", "").Replace(errorText[0]))

	if len(errorText) > 1 {
		fmt.Println(errorText[2])
		fmt.Println(errorText[3])
		fmt.Println()
	}

	fmt.Println(text.GlowPink("Try mt --help for usage."))
	os.Exit(1)
}

func showUpdateMessage(version string) {
	fmt.Println(config.Version)
	fmt.Println(config.AppConfig.Version)
	latestVersion := config.GetNewestCliVersionFunction(version)
	fmt.Println(latestVersion)
	if version < latestVersion {
		if version != "dev" {
			lines := []string{
				"A new update is available",
				"Current Version: " + config.AppConfig.Version,
				"Latest  Version: " + latestVersion,
			}
			ui.Border(lines)
		}
	}
}
