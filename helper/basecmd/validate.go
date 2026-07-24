package basecmd

import (
	"fmt"
	"io"
	"os"

	"github.com/mattia37773/mt/config"
	"github.com/mattia37773/mt/ui/text"
)

func ValidateProjectName(out io.Writer) string {
	CheckConfigFile(out)

	var projectName string = config.ProjectConfig.ProjectName

	if projectName == "" {
		fmt.Fprint(out, text.Red("Error: The Project name hasn't been set\n"))
		fmt.Fprint(out, "Please add it in .mt.yaml \n")
		os.Exit(2)
	}
	return projectName
}

func ValidateContainer(out io.Writer, container string, containerType string) {
	CheckConfigFile(out)

	if container == "" {
		fmt.Fprint(out, text.Red("Error: No "+containerType+" container is set!\n"))
		fmt.Fprint(out, "Please add it in .mt.yaml \n")
		os.Exit(2)
	}
}

func ValidateDockerPath(out io.Writer) string {
	CheckConfigFile(out)

	var dockerPath string = config.ProjectConfig.Paths.DockerCompose

	if dockerPath == "" {
		fmt.Fprint(out, text.Red("Error: No Docker Compose path is set!\n"))
		fmt.Fprint(out, "Please add it in .mt.yaml \n")
		os.Exit(2)
	}

	_, err := os.Stat(dockerPath)
	if err != nil {
		fmt.Fprint(out, text.Red("The docker compose file doesnt exist: ")+dockerPath+"\n")
		fmt.Fprint(out, "Please add it in .mt.yaml \n")
		os.Exit(2)
	}

	return dockerPath
}

func CheckConfigFile(out io.Writer) {
	_, err := os.Stat(".mt.yaml")
	if err != nil {
		fmt.Fprint(out, text.Red("Error: no mt.yaml exist \n"))
		fmt.Fprint(out, "Use this command to create it: mt config \n")
		os.Exit(2)
	}
}
