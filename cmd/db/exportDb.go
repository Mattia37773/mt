/*
Copyright © 2026 Matze
*/
package db

import (
	"fmt"
	"io"

	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/helper/basecmd"
	"github.com/mattia37773/mt/helper/docker"
	"github.com/mattia37773/mt/helper/shell"
	"github.com/mattia37773/mt/ui/text"
	"github.com/spf13/cobra"
)

var exportDbCmd = &cobra.Command{
	Use:                   "export",
	Short:                 "Export the local database",
	GroupID:               "db",
	DisableFlagsInUseLine: true,
	RunE: func(c *cobra.Command, args []string) error {
		out := c.OutOrStdout()

		err := exportExecute(out, args)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	DbCmd.AddCommand(exportDbCmd)
}

func exportExecute(out io.Writer, args []string) error {
	projectName, projectNameErr := basecmd.ValidateProjectName(out)
	if projectNameErr != nil {
		return projectNameErr
	}

	_, containerErr := docker.ContainerExits(out, projectName+"-"+config.ProjectConfig.DB.ContainerName)
	if containerErr != nil {
		return containerErr
	}

	cmd, exportGenError := exportGen(out, args)
	if exportGenError != nil {
		return exportGenError
	}

	dockerCmd := append([]string{"docker"}, cmd...)
	exportGenExecErr := shell.ExecuteCommandOnlyErrors(dockerCmd)
	if exportGenExecErr != nil {
		return exportGenExecErr
	}

	copyArgs := copyExportGen(out)
	copyCmdExecErr := shell.ExecuteCommandOnlyErrors(copyArgs)
	if copyCmdExecErr != nil {
		return copyCmdExecErr
	}

	removeArgs := removeExportGen(out)
	removeExecErr := shell.ExecuteCommandOnlyErrors(removeArgs)
	if removeExecErr != nil {
		return removeExecErr
	}

	fmt.Fprint(out, text.Green("Created Dump Successfully\n"))

	return nil
}

func exportGen(out io.Writer, args []string) ([]string, error) {
	db, dbError := config.GetDbConfig()
	if dbError != nil {
		return []string{}, dbError
	}
	return basecmd.ExecuteDbCommand(out, db.Export, args)
}

func copyExportGen(out io.Writer) []string {
	var projectName string = config.ProjectConfig.ProjectName
	db, _ := config.GetDbConfig()

	containerName := projectName + "-" + config.ProjectConfig.DB.ContainerName

	sourcePath := fmt.Sprintf("%s:/tmp/%s%s", containerName, projectName, db.Filetype)

	return []string{"docker", "cp", sourcePath, "./"}
}

func removeExportGen(out io.Writer) []string {
	var projectName string = config.ProjectConfig.ProjectName
	db, _ := config.GetDbConfig()

	containerName := projectName + "-" + config.ProjectConfig.DB.ContainerName

	targetFile := fmt.Sprintf("/tmp/%s%s", projectName, db.Filetype)

	return []string{"docker", "exec", containerName, "rm", "-rf", targetFile}
}
