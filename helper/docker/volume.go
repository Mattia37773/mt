package docker

import (
	"os/exec"
	"strings"
)

func GetVolumesByProject(projectName string) []string {
	out, _ := exec.Command("docker", "volume", "ls", "--filter", "name="+projectName, "--format", "{{.Name}}").Output()
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) == 1 && lines[0] == "" {
		return []string{}
	}
	return lines
}
