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

	containerErr := basecmd.CheckContainerExits(out, config.ProjectConfig.ProjectName+"-"+config.ProjectConfig.DB.ContainerName)
	if containerErr != nil {
		return containerErr
	}

	cmd, exportGenError := exportGen(out, args)
	if exportGenError != nil {
		return exportGenError
	}

	basecmd.ExecuteCommandOnlyErrors(out, cmd)

	//copy := exec.Command("docker", "cp", projectName+"-"+config.ProjectConfig.DB.ContainerName+":/tmp/"+projectName+db.Filetype, "./")
	copyArgs := copyExportGen(out)
	copyExportCommand := exec.Command(copyArgs[0], copyArgs[1:]...)
	copyExportCommand.Output()

	//rm := exec.Command("docker", "exec", projectName+"-"+config.ProjectConfig.DB.ContainerName, "rm", "-rf", "/tmp/"+projectName+db.Filetype)
	removeArgs := removeExportGen(out)
	removeExportCmd := exec.Command(removeArgs[0], removeArgs[1:]...)
	removeExportCmd.Output()

	fmt.Fprint(out, text.Green("Created Dump Successfully\n"))

	return nil
}

func exportGen(out io.Writer, args []string) ([]string, error) {
	return basecmd.ExecuteDbCommand(out, config.GetDbConfig().Export, args)
}

func copyExportGen(out io.Writer) []string {
	var projectName string = config.ProjectConfig.ProjectName
	var db config.DbConfig = config.GetDbConfig()
	containerName := projectName + "-" + config.ProjectConfig.DB.ContainerName

	sourcePath := fmt.Sprintf("%s:/tmp/%s%s", containerName, projectName, db.Filetype)

	return []string{"docker", "cp", sourcePath, "./"}
}

func removeExportGen(out io.Writer) []string {
	var projectName string = config.ProjectConfig.ProjectName
	var db config.DbConfig = config.GetDbConfig()
	containerName := projectName + "-" + config.ProjectConfig.DB.ContainerName

	targetFile := fmt.Sprintf("/tmp/%s%s", projectName, db.Filetype)

	return []string{"docker", "exec", containerName, "rm", "-rf", targetFile}
}
