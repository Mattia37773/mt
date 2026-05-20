/*
Copyright © 2026 Matze
*/
package single

import (
	"fmt"
	"os/exec"

	"github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/functions/docker"
	"github.com/mattia37773/mt/functions/shell"
	"github.com/mattia37773/mt/functions/ui"

	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:                   "shell [container] [args]",
	Short:                 "Open a shell in a container",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("Error: Missing argument CONTAINER")
		}

		user, _ := cmd.Flags().GetString("user")

		shellSingle(args, user)
		return nil
	},
}

func init() {
	shellCmd.Flags().StringP("user", "u", "", "User to Connect to the Container")
	cmd.RootCmd.AddCommand(shellCmd)
}

func shellSingle(args []string, user string) {
	var projectName string = config.GetProjectName()
	var container string = args[0]
	var cmd *exec.Cmd

	fmt.Printf(ui.Green("Project %s \n"), projectName)
	fmt.Println("")

	docker.ContainerExists(projectName + "-" + container)
	fmt.Printf("Opening Shell %s in %s-%s \n", "bash", projectName, container)

	if user != "" {
		cmd = exec.Command("sh", "-c", "docker exec -u "+user+" -it "+projectName+"-"+container+" $(docker exec "+projectName+"-"+container+" sh -c 'command -v bash || echo /bin/sh')")

	} else {
		cmd = exec.Command("sh", "-c", "docker exec -it "+projectName+"-"+container+" $(docker exec "+projectName+"-"+container+" sh -c 'command -v bash || echo /bin/sh')")
	}

	shell.ExecuteCommand(cmd)
}
