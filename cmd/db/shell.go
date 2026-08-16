/*
Copyright © 2026 Matze
*/
package db

import (
	"io"

	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/helper/basecmd"
	"github.com/mattia37773/mt/helper/docker"

	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:                   "shell",
	Short:                 "Connect to the database",
	GroupID:               "db",
	DisableFlagsInUseLine: true,
	RunE: func(c *cobra.Command, args []string) error {
		out := c.OutOrStdout()
		projectName, projectNameErr := basecmd.ValidateProjectName(out)
		if projectNameErr != nil {
			return projectNameErr
		}

		_, containerErr := docker.ContainerExits(out, projectName+"-"+config.ProjectConfig.DB.ContainerName)
		if containerErr != nil {
			return containerErr
		}

		cmd, err := shellGen(out, args)
		if err != nil {
			return err
		}

		basecmd.ExecuteCommand(out, cmd)
		return nil
	},
}

func init() {
	DbCmd.AddCommand(shellCmd)
}

func shellGen(out io.Writer, args []string) ([]string, error) {
	db, dbError := config.GetDbConfig()
	if dbError != nil {
		return []string{}, dbError
	}
	return basecmd.ExecuteDbCommand(out, db.Shell, args)
}
