/*
Copyright © 2026 Matze
*/
package db

import (
	"fmt"
	"os/exec"

	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/functions/docker"
	"github.com/mattia37773/mt/functions/shell"
	"github.com/mattia37773/mt/functions/ui"

	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:                   "shell",
	Short:                 "Connect to the database",
	GroupID:               "db",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		shellDb(args)
	},
}

func init() {
	DbCmd.AddCommand(shellCmd)
}

func shellDb(args []string) {

	var projectName string = config.GetProjectName()
	var db config.DbConfig = config.GetDbConfig()

	fmt.Printf(ui.Green("Project %s \n"), projectName)
	fmt.Println("")
	docker.ContainerExists(projectName + "-" + db.Container)
	fmt.Printf("Opening a %s shell \n", db.Name)

	cmd := exec.Command("docker", "exec", "-it", projectName+"-"+db.Container, "sh", "-c", db.Shell)

	shell.ExecuteCommand(cmd)
}
