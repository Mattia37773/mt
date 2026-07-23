package docker

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/mattia37773/mt/ui/text"
)

func IsContainerRunning(containerName string) bool {
	cmd := exec.Command("docker", "inspect", "-f", "{{.State.Running}}", containerName)
	out, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(out)) == "true"
}

func GetContainerStatus(containerName string) string {
	cmd := exec.Command("docker", "inspect", "--format", "{{.State.Status}}", containerName)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = io.Discard

	err := cmd.Run()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(out.String())
}

func GetContainerStartedAt(containerName string) string {
	cmd := exec.Command("docker", "inspect", "--format", "{{.State.StartedAt}}", containerName)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = io.Discard

	err := cmd.Run()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(out.String())
}

func CommandExistsInContainer(container string, command string) {
	cmd := exec.Command("docker", "exec", container, "sh", "-c", "command -v "+command)
	err := cmd.Run()
	if err != nil {
		fmt.Printf(text.Red("Error: The command: %s isn't available inside the %s\n"), command, container)
		os.Exit(1)
	}
}

func ContainerExists(name string) {
	cmd := exec.Command("docker", "ps", "-a", "--format", "{{.Names}}")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	containers := strings.Split(string(output), "\n")
	var success bool = false
	for _, c := range containers {
		if c == name {
			success = true
		}
	}

	if !success {
		fmt.Printf(text.Red("Error: The Container: %s doesn't exist\n"), name)
		fmt.Printf("Did you forget to run mt stack start? \n")
		os.Exit(1)
	}

}
