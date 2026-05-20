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

var yarnCmd = &cobra.Command{
	Use:                   "yarn [args]",
	Short:                 "Run Yarn in the frontend container",
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		yarnSingle(args)
	},
}

func init() {

	cmd.RootCmd.AddCommand(yarnCmd)
}

func yarnSingle(args []string) {
	var projectName string = config.GetProjectName()
	var container string = config.FrontendContainer()

	fmt.Printf(ui.Green("Project %s \n"), projectName)
	fmt.Println("")

	docker.ContainerExists(projectName + "-" + container)
	docker.CommandExistsInContainer(projectName+"-"+container, "yarn")
	fmt.Printf("Running Yarn inside %s-%s \n", projectName, container)

	cmdStr := "yarn " + strings.Join(args, " ")
	cmd := exec.Command("docker", "exec", "-it", projectName+"-"+container, "sh", "-c", cmdStr)

	shell.ExecuteCommand(cmd)
}
