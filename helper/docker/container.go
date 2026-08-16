/*
Copyright © 2026 Matze
*/
package docker

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
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

func CommandExistsInContainer(container string, command string) bool {
	fields := strings.Fields(command)
	if len(fields) == 0 {
		return false
	}
	binary := fields[0]
	cmd := exec.Command("docker", "exec", container, "sh", "-c", "command -v "+binary)
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}

func ContainerExits(out io.Writer, name string) (bool, error) {
	cmd := exec.Command("docker", "ps", "-a", "--format", "{{.Names}}")
	output, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("Somehting is wrong with the contaner %s \n, %s", name, err)
	}

	containers := strings.Split(string(output), "\n")
	for _, c := range containers {
		if c == name {
			return true, nil
		}
	}

	return false, fmt.Errorf("The Container: %s doesn't exist \nDid you forget to run mt stack start?", name)
}
