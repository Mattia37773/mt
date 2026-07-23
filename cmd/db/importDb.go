/*
Copyright © 2026 Matze
*/
package db

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/helper/docker"
	"github.com/mattia37773/mt/helper/shell"
	"github.com/mattia37773/mt/ui/form"
	"github.com/mattia37773/mt/ui/text"
	"github.com/spf13/cobra"
)

var importDbCmd = &cobra.Command{
	Use:                   "import [flags]",
	Short:                 "Import a database dump",
	GroupID:               "db",
	DisableFlagsInUseLine: true,
	RunE: func(c *cobra.Command, args []string) error {
		var projectName string = config.ProjectConfig.ProjectName
		var db config.DbConfig = config.GetDbConfig()
		out := c.OutOrStdout()

		// File Input
		var importFile string
		importFile, _ = c.Flags().GetString("file")
		if importFile == "" {
			// Picks the file
			var err error
			importFile, err = form.FilePicker(".")
			if err != nil {
				fmt.Fprintf(out, "\033[31mError: %s\033[0m \n", err)
				os.Exit(1)
			}

		}
		fmt.Fprintln(out, importFile)

		importDb(out, importFile, projectName, db)
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

func importDb(out io.Writer, file string, projectName string, db config.DbConfig) {
	fmt.Fprintf(out, text.Green("Project %s \n"), projectName)
	fmt.Fprintln(out, "")
	docker.ContainerExists(projectName + "-" + config.ProjectConfig.DB.ContainerName)

	copy := exec.Command("docker", "cp", "./"+file, projectName+"-"+config.ProjectConfig.DB.ContainerName+":/tmp/"+projectName+db.Filetype)
	shell.ExecuteCommandOnlyErrors(copy)

	cmd := exec.Command("docker", "exec", "-i", projectName+"-"+config.ProjectConfig.DB.ContainerName+"", "bash", "-pc", db.Import)
	shell.ExecuteCommandOnlyErrors(cmd)

	rm := exec.Command("docker", "exec", projectName+"-"+config.ProjectConfig.DB.ContainerName, "rm", "-rf", "/tmp/blueprint.gzip")
	rm.Output()

	fmt.Fprintf(out, text.Green("Imported the Database dump into %s\n"), projectName)
}
