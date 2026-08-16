/*
Copyright © 2026 Matze
*/
package db

import (
	"fmt"
	"io"
	"os"

	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/helper/basecmd"
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
		out := c.OutOrStdout()

		// File Input
		var importFile string
		importFile, _ = c.Flags().GetString("file")
		if importFile == "" {
			// Picks the file
			var err error
			importFile, err = form.FilePicker(".")
			if err != nil {
				return fmt.Errorf("%s", err)
			}

		}

		if importFile == "file" && len(args) > 0 {
			importFile = args[0]
		}
		_, err := os.Stat(importFile)
		if err != nil {
			return fmt.Errorf("The to be imported file doesn't exist: %s \nPlease choose a other file", importFile)
		}

		fmt.Fprintln(out, importFile)

		importErr := importExecute(out, importFile)
		if importErr != nil {
			return importErr
		}

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

func importExecute(out io.Writer, file string) error {
	projectName, projectNameErr := basecmd.ValidateProjectName(out)
	if projectNameErr != nil {
		return projectNameErr
	}

	_, containerErr := docker.ContainerExits(out, projectName+"-"+config.ProjectConfig.DB.ContainerName)
	if containerErr != nil {
		return containerErr
	}

	fmt.Fprintln(out, "")
	fmt.Fprintf(out, text.Green("Project %s \n"), projectName)

	// copy
	copy := copyImportGen(out, file)
	copyCmdExecErr := shell.ExecuteCommandOnlyErrors(copy)
	if copyCmdExecErr != nil {
		return copyCmdExecErr
	}

	// import
	importArgs := importGen(out, file)
	importExecErr := shell.ExecuteCommandOnlyErrors(importArgs)
	if importExecErr != nil {
		return importExecErr
	}

	// delete
	rmArgs := removeImportGen(out)
	removeExecErr := shell.ExecuteCommandOnlyErrors(rmArgs)
	if removeExecErr != nil {
		return removeExecErr
	}

	fmt.Fprintf(out, text.Green("Imported the Database dump into %s\n"), projectName)

	return nil
}

func importGen(out io.Writer, file string) []string {
	var projectName string = config.ProjectConfig.ProjectName
	db, _ := config.GetDbConfig()
	containerName := projectName + "-" + config.ProjectConfig.DB.ContainerName

	return []string{"docker", "exec", "-i", containerName, "bash", "-pc", db.Import}
}

func copyImportGen(out io.Writer, file string) []string {
	var projectName string = config.ProjectConfig.ProjectName
	db, _ := config.GetDbConfig()

	containerName := projectName + "-" + config.ProjectConfig.DB.ContainerName

	targetPath := fmt.Sprintf("%s:/tmp/%s%s", containerName, projectName, db.Filetype)

	return []string{"docker", "cp", "./" + file, targetPath}
}

func removeImportGen(out io.Writer) []string {
	var projectName string = config.ProjectConfig.ProjectName
	db, _ := config.GetDbConfig()
	containerName := projectName + "-" + config.ProjectConfig.DB.ContainerName

	targetFile := fmt.Sprintf("/tmp/%s%s", projectName, db.Filetype)

	return []string{"docker", "exec", containerName, "rm", "-rf", targetFile}
}
