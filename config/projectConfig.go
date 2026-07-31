/*
Copyright © 2026 Matze
*/
package config

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/joho/godotenv"
)

var exitFunc = os.Exit

type ConfigStruct struct {
	ProjectName string `yaml:"projectName"`

	Paths struct {
		DockerCompose string `yaml:"dockerCompose"`
		Env           string `yaml:"env"`
	} `yaml:"paths"`

	DB struct {
		ContainerName string `yaml:"containerName"`
		Provider      string `yaml:"provider"`
		Name          string `yaml:"name"`
		Password      string `yaml:"password"`
		User          string `yaml:"user"`
	} `yaml:"db"`

	Backend struct {
		ContainerName string `yaml:"containerName"`
	} `yaml:"backend"`

	Frontend struct {
		ContainerName string `yaml:"containerName"`
	} `yaml:"frontend"`

	Main struct {
		ContainerName string `yaml:"containerName"`
	} `yaml:"main"`
}

var ProjectConfig ConfigStruct

func ParseConfigFile() error {

	data, err := os.ReadFile(".mt.yaml")
	if err != nil {
		return nil
	}

	// for errors with formatting
	if err := yaml.Unmarshal([]byte(data), &ProjectConfig); err != nil {
		return fmt.Errorf("An error osccurred while parsing '.mt.yaml': \n%s", err)
	}

	envFile, err := getEnvFile(data)
	if err != nil {
		return err
	}

	err = godotenv.Load(envFile)
	if err != nil {
		return fmt.Errorf("The defined env file in .mt.yaml doesnt exist: %s", envFile)
	}

	// check and resolve env vars
	expandedData, err := expandEnvVarsStrict(string(data))
	if err != nil {
		return fmt.Errorf("%v\n", err)
	}

	// resolve to struct
	if err := yaml.Unmarshal([]byte(expandedData), &ProjectConfig); err != nil {
		return fmt.Errorf("An error occurred while parsing '.mt.yaml'")
	}
	return nil
}

func expandEnvVarsStrict(input string) (string, error) {
	// search after ${VAR_NAME} or $VAR_NAME
	re := regexp.MustCompile(`\$\{?([a-zA-Z_][a-zA-Z0-9_]*)\}?`)
	var missingErr error

	result := re.ReplaceAllStringFunc(input, func(match string) string {
		varName := re.FindStringSubmatch(match)[1]

		val, exists := os.LookupEnv(varName)
		if !exists {
			missingErr = fmt.Errorf("required environment variable '%s' is not set", varName)
			return match
		}
		return val
	})

	if missingErr != nil {
		return "", missingErr
	}

	return result, nil
}

func getEnvFile(data []byte) (string, error) {

	path, _ := yaml.PathString("$.paths.env")

	var envFile string
	if err := path.Read(strings.NewReader(string(data)), &envFile); err != nil {
		return "", fmt.Errorf("The defined env file in .mt.yaml doesnt exist: %s", envFile)
	}

	if envFile == "" {
		return "", fmt.Errorf("No Envfile is set: %s \n Please set one in .mt.yaml", envFile)
	}

	return envFile, nil
}
