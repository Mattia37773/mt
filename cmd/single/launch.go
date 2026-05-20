/*
Copyright © 2026 Matze
*/
package single

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/functions/env"
	"github.com/mattia37773/mt/functions/shell"
	"github.com/mattia37773/mt/functions/ui"

	"github.com/spf13/cobra"
)

var launchCmd = &cobra.Command{
	Use:                   "launch",
	Short:                 "Launch the website in the browser",
	Aliases:               []string{"open"},
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		launchSingle()
	},
}

func init() {
	cmd.RootCmd.AddCommand(launchCmd)
}

func launchSingle() {
	var projectName string = config.GetProjectName()
	var primaryUrl string = env.Get("PRIMARY_SITE_URL")

	fmt.Printf(ui.Green("Project %s \n"), projectName)
	fmt.Println()

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", primaryUrl)
	case "windows":
		cmd = exec.Command("start", primaryUrl)
	default:
		cmd = exec.Command("xdg-open", primaryUrl)
	}

	shell.ExecuteCommand(cmd)

	fmt.Printf(ui.Green("Opening %s in your browser\n"), primaryUrl)
}
