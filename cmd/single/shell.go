/*
Copyright © 2026 Matze
*/
package single

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/helper/docker"
	"github.com/mattia37773/mt/helper/shell"

	"github.com/mattia37773/mt/ui/text"

	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:                   "shell [container] [args]",
	Short:                 "Open a shell in a container",
	DisableFlagsInUseLine: true,
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("Error: Missing argument CONTAINER")
		}

		user, _ := c.Flags().GetString("user")

		shellSingle(c.OutOrStdout(), args, user)
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(shellCmd)
	shellCmd.Flags().StringP("user", "u", "", "User to Connect to the Container")
}

func shellSingle(out io.Writer, args []string, user string) {
	var projectName string = config.ProjectConfig.ProjectName
	var container string = args[0]
	var cmd *exec.Cmd

	fmt.Fprintf(out, text.Green("Project %s \n"), projectName)
	fmt.Fprintln(out, "")

	docker.ContainerExists(projectName + "-" + container)
	fmt.Fprintf(out, "Opening Shell %s in %s-%s \n", "bash", projectName, container)

	if user != "" {
		cmd = exec.Command("sh", "-c", "docker exec -u "+user+" -it "+projectName+"-"+container+" $(docker exec "+projectName+"-"+container+" sh -c 'command -v bash || echo /bin/sh')")

	} else {
		cmd = exec.Command("sh", "-c", "docker exec -it "+projectName+"-"+container+" $(docker exec "+projectName+"-"+container+" sh -c 'command -v bash || echo /bin/sh')")
	}

	shell.ExecuteCommand(cmd)
}
