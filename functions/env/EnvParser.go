/*
Copyright © 2026 Matze
*/
package env

import (
	"fmt"
	"os"

	"github.com/mattia37773/mt/functions/ui"
	"github.com/mattia37773/mt/functions/yaml"

	"github.com/joho/godotenv"
)

func Get(envRequestedValue string) string {
	var envFile string = yaml.GetConfig("projectCompose.paths.env", ".env")

	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Println(ui.Red("Error: .env is missing"))
		os.Exit(1)
	}

	var envValue string = os.Getenv(envRequestedValue)
	if envValue == "" {
		fmt.Printf(ui.Red("Error: %s is not set in the %s file\n"), envRequestedValue, envFile)
		os.Exit(1)
	}

	return envValue
}
