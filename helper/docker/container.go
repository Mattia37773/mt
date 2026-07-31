/*
Copyright © 2026 Matze
*/
package docker

import (
	"bytes"
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
	cmd := exec.Command("docker", "exec", container, "sh", "-c", "command -v "+command)
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}

func ContainerExists(name string) (bool, error) {
	cmd := exec.Command("docker", "ps", "-a", "--format", "{{.Names}}")
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	containers := strings.Split(string(output), "\n")
	for _, c := range containers {
		if c == name {
			return true, nil
		}
	}

	return false, err
}
