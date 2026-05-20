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

var composerCmd = &cobra.Command{
	Use:                   "composer [args]",
	Short:                 "Run Composer inside in a container",
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		composerSingle(args)
	},
}

func init() {
	cmd.RootCmd.AddCommand(composerCmd)
}

func composerSingle(args []string) {
	var projectName string = config.GetProjectName()
	var container string = config.BackendContainer()

	fmt.Printf(ui.Green("Project %s \n"), projectName)
	fmt.Println("")
	docker.ContainerExists(projectName + "-" + container)
	docker.CommandExistsInContainer(projectName+"-"+container, "composer")

	fmt.Printf("Running Composer inside %s-%s \n", projectName, container)
	cmdStr := "composer " + strings.Join(args, " ")
	cmd := exec.Command("docker", "exec", "-it", projectName+"-"+container, "sh", "-c", cmdStr)

	shell.ExecuteCommand(cmd)
}
