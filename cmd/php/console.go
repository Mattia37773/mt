/*
Copyright © 2026 Matze
*/
package php

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/functions/docker"
	"github.com/mattia37773/mt/functions/shell"
	"github.com/mattia37773/mt/functions/ui"

	"github.com/spf13/cobra"
)

var consoleCmd = &cobra.Command{
	Use:                   "console [args]",
	Short:                 "Run the Symfony console",
	GroupID:               "php",
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		consoleSymfony(args)
	},
}

func init() {
	PhpCmd.AddCommand(consoleCmd)
}

func consoleSymfony(args []string) {
	var projectName string = config.GetProjectName()
	var phpContainer string = config.BackendContainer()

	fmt.Printf(ui.Green("Project %s \n "), projectName)
	fmt.Println("")
	docker.ContainerExists(projectName + "-" + phpContainer)
	docker.FileExistsInContainer(projectName+"-"+phpContainer, "bin/console")
	fmt.Printf("Running Symfony Console in %s-%s \n", projectName, phpContainer)

	cmdStr := "bin/console " + strings.Join(args, " ")
	cmd := exec.Command("docker", "exec", "-it", projectName+"-"+phpContainer, "bash", "-c", cmdStr)

	shell.ExecuteCommand(cmd)
}
