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

var craftCmd = &cobra.Command{
	Use:                   "craft [args]",
	Short:                 "Run the Craft CMS cli",
	GroupID:               "php",
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		craft(args)
	},
}

func init() {
	PhpCmd.AddCommand(craftCmd)
}

func craft(args []string) {
	var projectName string = config.GetProjectName()
	var container string = config.BackendContainer()

	fmt.Printf(ui.Green("Project %s \n"), projectName)
	fmt.Println("")
	docker.ContainerExists(projectName + "-" + container)
	docker.FileExistsInContainer(projectName+"-"+container, "craft")
	fmt.Printf("Running Craft inside %s-%s \n", projectName, container)

	cmdStr := "php craft " + strings.Join(args, " ")
	cmd := exec.Command("docker", "exec", "-it", projectName+"-"+container, "sh", "-c", cmdStr)

	shell.ExecuteCommand(cmd)
}
