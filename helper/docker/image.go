package docker

import (
	"io"
	"os/exec"
	"strings"
)

func ImageExists(imageName string) bool {
	cmd := exec.Command("docker", "image", "inspect", imageName)
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard

	err := cmd.Run()
	return err == nil
}

func GetImagesByProject(projectName string) []string {
	out, _ := exec.Command("docker", "images", "--filter", "reference="+projectName+"*", "--format", "{{.Repository}}").Output()
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) == 1 && lines[0] == "" {
		return []string{}
	}
	return lines
}
