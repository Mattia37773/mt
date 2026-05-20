/*
Copyright © 2026 Matze
*/
package single

import (
	"fmt"
	"os"

	"github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/functions/ui"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:                   "config",
	Short:                 "This command generates a config file for this project",
	Aliases:               []string{"init"},
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		configSingle(args)
	},
}

func init() {
	cmd.RootCmd.AddCommand(configCmd)
}

func configSingle(args []string) {
	generateConfigFile()

	fmt.Println(ui.Green("Configfile generated"))
}

func generateConfigFile() {
	_, err := os.Stat(".mt.yaml")
	if !os.IsNotExist(err) {
		fmt.Printf(ui.Red("Error: Config file already exists\n"))
		os.Exit(1)
	}

	file, err := os.Create(".mt.yaml")
	if err != nil {
		fmt.Printf(ui.Red("Error: Something went wrong with creating the config File\n"))
		os.Exit(1)
	}

	defer file.Close()

	_, err = file.WriteString(configDefaultContent())
	if err != nil {
		fmt.Printf(ui.Red("Error: Something went wrong with creating the config File\n"))
		os.Exit(1)
	}
}

// Default config file content
func configDefaultContent() string {

	return `
# =============================================================================
# MT CLI Tool Configuration
# =============================================================================

# Docker Compose Project Configuration
projectCompose:
    # paths relative to project root
    paths:
        dockerCompose: docker/docker-compose.yaml
        env: .env
    db:
        containerName: "db"
        type: "mysql"
        name: "appDb"
        password: "password"
        user: "uDb"
    backend:
        containerName: fpm
    frontend:
        containerName: fpm		

`
}
