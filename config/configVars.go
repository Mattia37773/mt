/*
Copyright © 2026 Matze
*/
package config

import (
	"fmt"
	"os"
	"regexp"

	"github.com/mattia37773/mt/functions/env"
	"github.com/mattia37773/mt/functions/file"
	"github.com/mattia37773/mt/functions/ui"
	"github.com/mattia37773/mt/functions/yaml"
)

// .env Config Settings

func GetProjectName() string {
	var projectName string = env.Get("PROJECT_NAME")
	match, _ := regexp.MatchString("^[a-z]+$", projectName)

	if match == false {
		fmt.Println(ui.Red("Error: invalid project name"))
		fmt.Println("Enviroment variable 'PROJECT_NAME' Can Only Contain letters from a-z and no whitespaces")
		os.Exit(1)
	}

	return projectName
}

// .mt.yaml Config Settings

func DockerPath() string {
	var docker string = yaml.GetConfig("projectCompose.paths.dockerCompose", "docker/docker-compose.yaml")

	file.ExistsMany([]string{docker})
	return docker
}

func envPath() string {
	var env string = yaml.GetConfig("projectCompose.paths.env", ".env")
	file.ExistsMany([]string{env})
	return env
}

func DbContainer() string {
	return yaml.GetConfig("projectCompose.db.containerName", "db")
}

func DbType() string {
	return yaml.GetConfig("projectCompose.db.type", "mysql")
}

func DbName() string {
	return yaml.GetConfig("projectCompose.db[*].name", "appDb")
}

func dbUser() string {
	return yaml.GetConfig("projectCompose.db[*].user", "uDb")
}

func dbPassword() string {
	return yaml.GetConfig("projectCompose.db.password", "password")
}

func BackendContainer() string {
	return yaml.GetConfig("projectCompose.backend.containerName", "fpm")
}

func FrontendContainer() string {
	return yaml.GetConfig("projectCompose.frontend.containerName", "fpm")
}
