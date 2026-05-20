/*
Copyright © 2026 Matze
*/
package single

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/functions/docker"
	"github.com/mattia37773/mt/functions/shell"
	"github.com/mattia37773/mt/functions/ui"

	"github.com/spf13/cobra"
)

var bunCmd = &cobra.Command{
	Use:                   "bun [args]",
	Short:                 "Run Bun inside in a container",
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		bunSingle(args)
	},
}

func init() {
	cmd.RootCmd.AddCommand(bunCmd)
}

func bunSingle(args []string) {
	var projectName string = config.GetProjectName()
	var container string = config.FrontendContainer()

	fmt.Printf(ui.Green("Project %s \n"), projectName)
	fmt.Println("")
	docker.ContainerExists(projectName + "-" + container)
	docker.CommandExistsInContainer(projectName+"-"+container, "bun")
	fmt.Printf("Running Bun inside %s-%s \n", projectName, container)

	cmdStr := "bun " + strings.Join(args, " ")
	cmd := exec.Command("docker", "exec", "--interactive", "--tty", projectName+"-"+container, "sh", "-c", cmdStr)

	shell.ExecuteCommand(cmd)
}
