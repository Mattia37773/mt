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
	"github.com/mattia37773/mt/functions/shell"
	"github.com/mattia37773/mt/functions/ui"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:                   "run [commands] [args]",
	Short:                 "Run a Command inside the Backend Container",
	Aliases:               []string{"exec"},
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		runSingle(args)
	},
}

func init() {
	cmd.RootCmd.AddCommand(runCmd)
}

func runSingle(args []string) {
	var projectName string = config.GetProjectName()
	var container string = config.BackendContainer()

	cmdStr := strings.Join(args, " ")

	fmt.Printf(ui.Green("Project %s \n"), projectName)
	fmt.Println("")
	fmt.Printf("Running %s in %s-%s \n", cmdStr, projectName, container)

	cmd := exec.Command("docker", "exec", "-it", projectName+"-"+container, "sh", "-c", cmdStr)

	shell.ExecuteCommand(cmd)
}
