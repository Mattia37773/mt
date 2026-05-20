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

var exportDbCmd = &cobra.Command{
	Use:                   "export",
	Short:                 "Export the local database",
	GroupID:               "db",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		exportDb(args)
	},
}

func init() {
	DbCmd.AddCommand(exportDbCmd)
}

func exportDb(args []string) {
	var projectName string = config.GetProjectName()
	var db config.DbConfig = config.GetDbConfig()

	fmt.Printf(ui.Green("Project %s \n"), projectName)
	fmt.Println("")
	docker.ContainerExists(projectName + "-" + db.Container)

	cmd := exec.Command("docker", "exec", "-i", projectName+"-"+db.Container, "bash", "-c", db.Export)
	shell.ExecuteCommandOnlyErrors(cmd)

	copy := exec.Command("docker", "cp", projectName+"-"+db.Container+":/tmp/"+projectName+db.Filetype, "./")
	copy.Output()

	rm := exec.Command("docker", "exec", projectName+"-"+db.Container, "rm", "-rf", "/tmp/"+projectName+db.Filetype)
	rm.Output()

	fmt.Printf(ui.Green("Created Dump Successfully\n"))
}
