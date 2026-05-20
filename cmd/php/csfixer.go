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

var csfixerCmd = &cobra.Command{
	Use:                   "csfixer [args]",
	Short:                 "Run the PHP coding standards fixer",
	Aliases:               []string{"cs", "fixer", "cs-fixer"},
	GroupID:               "php",
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		csfixerSymfony(args)
	},
}

func init() {
	PhpCmd.AddCommand(csfixerCmd)
}

func csfixerSymfony(args []string) {
	var projectName string = config.GetProjectName()
	var phpContainer string = config.BackendContainer()

	fmt.Printf(ui.Green("Project %s \n"), projectName)
	fmt.Println("")
	docker.ContainerExists(projectName + "-" + phpContainer)
	docker.FileExistsInContainer(projectName+"-"+phpContainer, "vendor/bin/php-cs-fixer")
	fmt.Printf("Running PHP CsFixer  in %s-%s \n", projectName, phpContainer)

	cmdStr := "vendor/bin/php-cs-fixer " + strings.Join(args, " ")
	cmd := exec.Command("docker", "exec", "-it", projectName+"-"+phpContainer, "sh", "-c", cmdStr)

	shell.ExecuteCommand(cmd)
}
