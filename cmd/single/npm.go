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

var npmCmd = &cobra.Command{
	Use:                   "npm [args]",
	Short:                 "Run npm in the frontend container",
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		npmSingle(args)
	},
}

func init() {
	cmd.RootCmd.AddCommand(npmCmd)
}

func npmSingle(args []string) {
	var projectName string = config.GetProjectName()
	var container string = config.FrontendContainer()

	fmt.Printf(ui.Green("Project %s \n"), projectName)
	fmt.Println("")
	docker.ContainerExists(projectName + "-" + container)
	docker.CommandExistsInContainer(projectName+"-"+container, "npm")
	fmt.Printf("Running Npm in %s-%s \n", projectName, container)

	cmdStr := "npm " + strings.Join(args, " ")
	cmd := exec.Command("docker", "exec", "-it", projectName+"-"+container, "sh", "-c", cmdStr)

	shell.ExecuteCommand(cmd)
}
