/*
Copyright © 2026 Matze
*/
package db

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/helper/basecmd"
	"github.com/mattia37773/mt/ui/text"
	"github.com/spf13/cobra"
)

var exportDbCmd = &cobra.Command{
	Use:                   "export",
	Short:                 "Export the local database",
	GroupID:               "db",
	DisableFlagsInUseLine: true,
	Run: func(c *cobra.Command, args []string) {
		exportDb(c.OutOrStdout(), args)
	},
}

func init() {
	DbCmd.AddCommand(exportDbCmd)
}

func exportDb(out io.Writer, args []string) {
	var projectName string = config.ProjectConfig.ProjectName
	var db config.DbConfig = config.GetDbConfig()

	basecmd.ExecuteDbCommand(out, db.Export, args)

	copy := exec.Command("docker", "cp", projectName+"-"+config.ProjectConfig.DB.ContainerName+":/tmp/"+projectName+db.Filetype, "./")
	copy.Output()

	rm := exec.Command("docker", "exec", projectName+"-"+config.ProjectConfig.DB.ContainerName, "rm", "-rf", "/tmp/"+projectName+db.Filetype)
	rm.Output()

	fmt.Fprint(out, text.Green("Created Dump Successfully\n"))
}
