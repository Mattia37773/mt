package config // TODO add tests

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/joho/godotenv"
	"github.com/mattia37773/mt/ui/text"
)

var exitFunc = os.Exit

type ConfigStruct struct {
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
}

var ProjectConfig ConfigStruct

func ParseConfigFile() {

	data, err := os.ReadFile(".mt.yaml")

	if err != nil {
		fmt.Print(text.Red("Error with opening '.mt.yaml':\n"))
		fmt.Printf("%v\n", err)
		exitFunc(1)
	}

	envFile := getEnvFile(data)

	err = godotenv.Load(envFile)
	if err != nil {
		fmt.Print(text.Red("The defined env file in .mt.yaml doesnt exist: " + envFile + "\n"))
		exitFunc(1)
	}

	// check and resolve env vars
	expandedData, err := expandEnvVarsStrict(string(data))
	if err != nil {
		fmt.Printf(text.Red("%v\n"), err)
		exitFunc(1)
	}

	// resolve to struct
	if err := yaml.Unmarshal([]byte(expandedData), &ProjectConfig); err != nil {
		fmt.Println(text.Red("Error with parsing the config file '.mt.yaml': "))
		fmt.Printf("%v\n", err)
		exitFunc(1)
	}
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

func getEnvFile(data []byte) string {

	path, _ := yaml.PathString("$.paths.env")

	var envFile string
	if err := path.Read(strings.NewReader(string(data)), &envFile); err != nil {
		return ".env"
	}

	if envFile == "" {
		fmt.Println(text.Red("No Envfile is set"))
		fmt.Println("Please set one in .mt.yaml")
	}

	return envFile
}
