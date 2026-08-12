/*
Copyright © 2026 Matze
*/
package single

import (
	"fmt"
	"io"

	"github.com/mattia37773/mt/cmd"
	"github.com/mattia37773/mt/helper/basecmd"

	"github.com/mattia37773/mt/ui/text"

	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:                   "shell [container] [args]",
	Short:                 "Open a shell in a container",
	DisableFlagsInUseLine: true,
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("Missing argument CONTAINER")
		}

		user, _ := c.Flags().GetString("user")

		cmd, err := shellGen(c.OutOrStdout(), user, args)
		if err != nil {
			return err
		}

		basecmd.ExecuteHostCommand(c.OutOrStdout(), cmd)
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(shellCmd)
	shellCmd.Flags().StringP("user", "u", "", "User to Connect to the Container")
}

func shellGen(out io.Writer, user string, args []string) ([]string, error) {
	projectName, err := basecmd.ValidateProjectName(out)
	if err != nil {
		return []string{}, err
	}

	var container string = args[0]

	containerErr := basecmd.CheckContainerExits(out, projectName+"-"+container)
	if containerErr != nil {
		return nil, containerErr
	}

	fmt.Fprintf(out, text.Green("Project %s \n"), projectName)
	fmt.Fprintln(out, "")
	fmt.Fprintf(out, "Opening Shell %s in %s-%s \n", "bash", projectName, container)

	var execArgs string
	if user != "" {
		execArgs = fmt.Sprintf("docker exec -u %s -it %s-%s $(docker exec %s-%s sh -c 'command -v bash || echo /bin/sh')", user, projectName, container, projectName, container)

	} else {
		execArgs = fmt.Sprintf("docker exec -it %s-%s $(docker exec %s-%s sh -c 'command -v bash || echo /bin/sh')", projectName, container, projectName, container)
	}

	return []string{"sh", "-c", execArgs}, nil
}
