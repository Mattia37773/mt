/*
Copyright © 2026 Matze
*/
package db

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/functions/docker"
	"github.com/mattia37773/mt/functions/file"
	"github.com/mattia37773/mt/functions/shell"
	"github.com/mattia37773/mt/functions/ui"

	"github.com/spf13/cobra"
)

var importDbCmd = &cobra.Command{
	Use:                   "import [flags]",
	Short:                 "Import a database dump",
	GroupID:               "db",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var projectName string = config.GetProjectName()
		var db config.DbConfig = config.GetDbConfig()

		// File Input
		var importFile string
		importFile, _ = cmd.Flags().GetString("file")
		if importFile == "" {
			// Picks the file
			var err error
			importFile, err = ui.FilePicker(".")
			if err != nil {
				fmt.Printf("\033[31mError: %s\033[0m \n", err)
				os.Exit(1)
			}

		}
		fmt.Println(importFile)

		file.Exists(importFile)

		importDb(importFile, projectName, db)
		return nil
	},
}

func init() {
	importDbCmd.Flags().StringP("file", "f", "", "File to import into the Database")
	DbCmd.AddCommand(importDbCmd)

	importDbCmd.RegisterFlagCompletionFunc("file", func(
		cmd *cobra.Command,
		args []string,
		toComplete string,
	) ([]string, cobra.ShellCompDirective) {

		return nil, cobra.ShellCompDirectiveDefault
	})
}

func importDb(file string, projectName string, db config.DbConfig) {
	fmt.Printf(ui.Green("Project %s \n"), projectName)
	fmt.Println("")
	docker.ContainerExists(projectName + "-" + db.Container)

	copy := exec.Command("docker", "cp", "./"+file, projectName+"-"+db.Container+":/tmp/"+projectName+db.Filetype)
	shell.ExecuteCommandOnlyErrors(copy)

	cmd := exec.Command("docker", "exec", "-i", projectName+"-"+db.Container+"", "bash", "-pc", db.Import)
	shell.ExecuteCommandOnlyErrors(cmd)

	rm := exec.Command("docker", "exec", projectName+"-"+db.Container, "rm", "-rf", "/tmp/blueprint.gzip")
	rm.Output()

	fmt.Printf(ui.Green("Imported the Database dump into %s\n"), projectName)
}
