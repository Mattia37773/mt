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

var phpstanCmd = &cobra.Command{
	Use:                   "phpstan [args]",
	Short:                 "Run phpStan analysis ",
	Aliases:               []string{"stan"},
	GroupID:               "php",
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		phpstanSymfony(args)
	},
}

func init() {
	PhpCmd.AddCommand(phpstanCmd)
}

func phpstanSymfony(args []string) {
	var projectName string = config.GetProjectName()
	var phpContainer string = config.BackendContainer()

	fmt.Printf(ui.Green("Project %s \n"), projectName)
	fmt.Println("")
	docker.ContainerExists(projectName + "-" + phpContainer)
	docker.FileExistsInContainer(projectName+"-"+phpContainer, "vendor/bin/phpstan")
	fmt.Printf("Running phpstan in %s-%s \n", projectName, phpContainer)

	cmdStr := " vendor/bin/phpstan " + strings.Join(args, " ")
	cmd := exec.Command("docker", "exec", "-it", projectName+"-"+phpContainer, "sh", "-c", cmdStr)

	shell.ExecuteCommand(cmd)
}
