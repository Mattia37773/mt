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

var phpunitCmd = &cobra.Command{
	Use:                   "phpunit [args]",
	Short:                 "Run PHPUnit tests",
	Aliases:               []string{"unit"},
	GroupID:               "php",
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		phpunitSymfony(args)
	},
}

func init() {
	PhpCmd.AddCommand(phpunitCmd)
}

func phpunitSymfony(args []string) {
	var projectName string = config.GetProjectName()
	var phpContainer string = config.BackendContainer()

	fmt.Printf(ui.Green("Project %s \n"), projectName)
	fmt.Println("")
	docker.ContainerExists(projectName + "-" + phpContainer)
	docker.FileExistsInContainer(projectName+"-"+phpContainer, "vendor/bin/phpunit")
	fmt.Printf("Running PHPUnit  in %s-%s \n", projectName, phpContainer)

	cmdStr := "vendor/bin/phpunit " + strings.Join(args, " ")
	cmd := exec.Command("docker", "exec", "-it", projectName+"-"+phpContainer, "sh", "-c", cmdStr)

	shell.ExecuteCommand(cmd)
}
