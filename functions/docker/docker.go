/*
Copyright © 2026 Matze
*/
package docker

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mattia37773/mt/functions/ui"
)

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
		fmt.Printf(ui.Red("Error: The Container: %s doesn't exist\n"), name)
		fmt.Printf("Did you forget to run mt stack start? \n")
		os.Exit(1)
	}

}

func FileExistsInContainer(container string, filePath string) {
	cmd := exec.Command("docker", "exec", container, "test", "-f", filePath)
	err := cmd.Run()
	if err != nil {
		fmt.Printf(ui.Red("Error: The file: %s doesn't exist inside the %s\n"), filePath, container)
		os.Exit(1)
	}
}

func CommandExistsInContainer(container string, command string) {
	cmd := exec.Command("docker", "exec", container, "sh", "-c", "command -v "+command)
	err := cmd.Run()
	if err != nil {
		fmt.Printf(ui.Red("Error: The command: %s isn't available inside the %s\n"), command, container)
		os.Exit(1)
	}
}
