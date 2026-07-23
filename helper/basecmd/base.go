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

func ExecuteBackendCommand(out io.Writer, baseCommand string, args []string) {
	// TODO  validation
	container := config.ProjectConfig.Backend.ContainerName
	executeCommand(out, baseCommand, container, args)
}

func ExecuteFrontendCommand(out io.Writer, baseCommand string, args []string) {
	// TODO  validation
	container := config.ProjectConfig.Frontend.ContainerName
	executeCommand(out, baseCommand, container, args)
}

func ExecuteDbCommand(out io.Writer, baseCommand string, args []string) {
	// TODO  validation
	container := config.ProjectConfig.DB.ContainerName
	executeCommandOnlyErrors(out, baseCommand, container, args)
}

func ExecuteMainCommand(out io.Writer, baseCommand string, args []string) {
	// TODO  validation
	container := config.ProjectConfig.Main.ContainerName
	executeCommand(out, baseCommand, container, args)
}

func executeCommand(out io.Writer, baseCommand string, container string, args []string) {
	// todo add valdation for those two
	var projectName string = config.ProjectConfig.ProjectName

	docker.ContainerExists(projectName + "-" + container)
	docker.CommandExistsInContainer(projectName+"-"+container, baseCommand)

	cmdStr := baseCommand + " " + strings.Join(args, " ")

	fmt.Fprintf(out, text.Green("Project %s \n"), projectName)
	fmt.Fprintln(out, "")
	fmt.Printf("Running %s in %s-%s \n", cmdStr, projectName, container)

	// checks if the terminal has a tty
	execArgs := []string{"exec"}
	if file, ok := out.(*os.File); ok && term.IsTerminal(file.Fd()) {
		execArgs = append(execArgs, "-it")
	}

	execArgs = append(execArgs, projectName+"-"+container, "sh", "-c", cmdStr)
	cmd := exec.Command("docker", execArgs...)

	cmd.Stdout = out
	cmd.Stderr = out

	shell.ExecuteCommand(cmd)
}

// only print errors to screen
func executeCommandOnlyErrors(out io.Writer, baseCommand string, container string, args []string) {
	// todo add valdation for those two
	var projectName string = config.ProjectConfig.ProjectName

	docker.ContainerExists(projectName + "-" + container)
	docker.CommandExistsInContainer(projectName+"-"+container, baseCommand)

	cmdStr := baseCommand + " " + strings.Join(args, " ")

	fmt.Fprintf(out, text.Green("Project %s \n"), projectName)
	fmt.Fprintln(out, "")
	fmt.Printf("Running %s in %s-%s \n", cmdStr, projectName, container)

	// checks if the terminal has a tty
	execArgs := []string{"exec"}
	if file, ok := out.(*os.File); ok && term.IsTerminal(file.Fd()) {
		execArgs = append(execArgs, "-it")
	}

	execArgs = append(execArgs, projectName+"-"+container, "sh", "-c", cmdStr)
	cmd := exec.Command("docker", execArgs...)

	cmd.Stdout = out
	cmd.Stderr = out

	shell.ExecuteCommandOnlyErrors(cmd)
}
