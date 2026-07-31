/*
Copyright © 2026 Matze
*/
package single

import (
	"fmt"
	"io"
	"os"

	"github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/ui/text"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:                   "config",
	Short:                 "Generate the default config file",
	Aliases:               []string{"init"},
	DisableFlagsInUseLine: true,
	RunE: func(c *cobra.Command, args []string) error {
		out := c.OutOrStdout()

		err := generateConfigFile(out)
		if err != nil {
			return err
		}

		fmt.Fprintln(out, text.Green("Configfile generated"))
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(configCmd)
}

func generateConfigFile(out io.Writer) error {
	_, err := os.Stat(".mt.yaml")
	if !os.IsNotExist(err) {
		return fmt.Errorf("Error: Config file already exists")

	}

	file, err := os.Create(".mt.yaml")
	if err != nil {
		return fmt.Errorf("Something went wrong with creating the config File")
	}

	defer file.Close()

	_, err = file.WriteString(configDefaultContent())
	if err != nil {
		return fmt.Errorf("Error: Something went wrong with creating the config File")
	}
	return nil
}

// Default config file content
func configDefaultContent() string {

	return `
# =============================================================================
# MT CLI Tool Configuration
# =============================================================================

# this should be the container prefix
projectName: projectname
paths:
	# docker compose path
    dockerCompose: docker/docker-compose.yaml
	# env file to load secrets. They can be referenced here like this: ${MYSQL_ROOT_PASSWORD}
    env: docker/.env
db:
    containerName: "db"
    provider: "mysql"
    name: "appDb"
    password: "password"
    user: "uDb"
backend:
    containerName: fpm
frontend:
    containerName: fpm
# container for the run command
main:
    containerName: fpm	

`
}
