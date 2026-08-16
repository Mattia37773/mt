/*
Copyright © 2026 Matze
*/
package basecmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/x/term"
	"github.com/mattia37773/mt/config"

	"github.com/mattia37773/mt/helper/docker"
	"github.com/mattia37773/mt/helper/shell"
	"github.com/mattia37773/mt/ui/text"
)

func ExecuteBackendCommand(out io.Writer, baseCommand string, args []string) ([]string, error) {
	container := config.ProjectConfig.Backend.ContainerName
	ValidateProjectName(out)
	ValidateContainer(out, container, "Backend")

	return BaseCommandGen(out, baseCommand, container, true, args)
}

func ExecuteFrontendCommand(out io.Writer, baseCommand string, args []string) ([]string, error) {
	container := config.ProjectConfig.Frontend.ContainerName
	ValidateProjectName(out)
	ValidateContainer(out, container, "Frontend")
	return BaseCommandGen(out, baseCommand, container, true, args)
}

func ExecuteDbCommand(out io.Writer, baseCommand string, args []string) ([]string, error) {
	container := config.ProjectConfig.DB.ContainerName
	ValidateProjectName(out)
	ValidateContainer(out, container, "Db")
	return BaseCommandGen(out, baseCommand, container, true, args)
}

func ExecuteMainCommand(out io.Writer, baseCommand string, args []string) ([]string, error) {
	container := config.ProjectConfig.Main.ContainerName
	ValidateProjectName(out)
	ValidateContainer(out, container, "Main")
	return BaseCommandGen(out, baseCommand, container, false, args)
}

func BaseCommandGen(out io.Writer, baseCommand string, container string, checkCommandExist bool, args []string) ([]string, error) {
	var projectName string = config.ProjectConfig.ProjectName

	if checkCommandExist != false {
		_, containerErr := docker.ContainerExits(out, projectName+"-"+container)
		if containerErr != nil {
			return nil, containerErr
		}

		commandExists := docker.CommandExistsInContainer(projectName+"-"+container, baseCommand)
		if commandExists == false {
			return nil, fmt.Errorf("The command: %s isn't available inside the %s", baseCommand, container)
		}
	}

	cmdStr := baseCommand + " " + strings.Join(args, " ")

	fmt.Fprintf(out, text.Green("Project %s \n"), projectName)
	fmt.Fprintln(out, "")
	fmt.Fprintf(out, "Running %s in %s-%s \n", cmdStr, projectName, container)

	// checks if the terminal has a tty
	execArgs := []string{"exec"}
	if file, ok := out.(*os.File); ok && term.IsTerminal(file.Fd()) {
		execArgs = append(execArgs, "-it")
	}

	execArgs = append(execArgs, projectName+"-"+container, "sh", "-c", cmdStr)

	return execArgs, nil
}

func ExecuteCommand(out io.Writer, execArgs []string) {
	cmd := exec.Command("docker", execArgs...)

	cmd.Stdout = out
	cmd.Stderr = out

	shell.ExecuteCommand(cmd)
}

func ExecuteHostCommand(out io.Writer, execArgs []string) {
	cmd := exec.Command(execArgs[0], execArgs[1:]...)

	cmd.Stdout = out
	cmd.Stderr = out
	cmd.Stdin = os.Stdin

	shell.ExecuteCommand(cmd)
}

func StackExecute(out io.Writer, execArgs []string, successText string) error {
	cmd := exec.Command(execArgs[0], execArgs[1:]...)

	cmd.Stdout = out
	cmd.Stderr = out
	cmd.Stdin = os.Stdin

	err := shell.ExecuteCommand(cmd)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	if successText != "" {
		fmt.Fprintln(out, text.Green(successText))
	}
	return nil
}
