/*
Copyright © 2026 Matze
*/
package stack

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/helper/shell"
	"github.com/mattia37773/mt/ui/text"
	"github.com/spf13/cobra"
)

var psCmd = &cobra.Command{
	Use:                   "ps",
	Short:                 "Show stack status",
	GroupID:               "stack",
	DisableFlagsInUseLine: true,
	Run: func(c *cobra.Command, args []string) {
		psStack(c.OutOrStdout())
	},
}

func init() {
	StackCmd.AddCommand(psCmd)
}

func psStack(out io.Writer) {
	// todo add valdation for those two
	var projectName string = config.ProjectConfig.ProjectName
	var dockerPath string = config.ProjectConfig.Paths.DockerCompose

	fmt.Fprintf(out, text.Green("Project %s \n"), projectName)
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "Stack Status")

	cmd := exec.Command("docker", "compose", "-f", dockerPath, "ps")

	// Leite Stdout & Stderr an den 'out' Writer um
	cmd.Stdout = out
	cmd.Stderr = out

	// ExecuteCommand sieht nun: cmd.Stdout ist NICHT nil, und lässt es unberührt!
	err := shell.ExecuteCommand(cmd)
	if err != nil {
		fmt.Fprint(out, text.Red("Something went wrong: "))
		fmt.Fprintln(out, err)
		os.Exit(1)
	}
}
