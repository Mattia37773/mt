package basecmd

import (
	"fmt"
	"io"
	"os"

	"github.com/mattia37773/mt/config"
)

func ValidateProjectName(out io.Writer) (string, error) {
	CheckConfigFile(out)

	var projectName string = config.ProjectConfig.ProjectName

	if projectName == "" {
		return "", fmt.Errorf("The Project name hasn't been set!!!\nPlease add it in .mt.yaml")
	}
	return projectName, nil
}

func ValidateContainer(out io.Writer, container string, containerType string) error {
	CheckConfigFile(out)

	if container == "" {
		return fmt.Errorf("No %s container is set!!!\nPlease add it in .mt.yaml", containerType)
	}
	return nil
}

func ValidateDockerPath(out io.Writer) (string, error) {
	CheckConfigFile(out)

	var dockerPath string = config.ProjectConfig.Paths.DockerCompose

	if dockerPath == "" {
		return "", fmt.Errorf("No Docker Compose path is set!!!\nPlease add it in .mt.yaml")
	}

	_, err := os.Stat(dockerPath)
	if err != nil {
		return "", fmt.Errorf("The docker compose file doesnt exist: %s", dockerPath)
	}

	return dockerPath, nil
}

func CheckConfigFile(out io.Writer) error {
	_, err := os.Stat(".mt.yaml")
	if err != nil {
		return fmt.Errorf("No mt.yaml exists!! Use this command to create it: mt config")
	}
	return nil
}
