/*
Copyright © 2026 Matze
*/
package single_test

import (
	"bytes"
	"testing"

	"github.com/mattia37773/mt/cmd"
	base "github.com/mattia37773/mt/tests/basetest"

	_ "github.com/mattia37773/mt/cmd/db"
	_ "github.com/mattia37773/mt/cmd/php"
	_ "github.com/mattia37773/mt/cmd/single"
	_ "github.com/mattia37773/mt/cmd/stack"
	"github.com/stretchr/testify/assert"
)

func TestHelpStyle(t *testing.T) {
	buf := new(bytes.Buffer)

	rootCmd := cmd.RootCmd
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"--help"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(buf.String())

	assert.Contains(t, cleanOutput, "Usage: mt [options] [command] [flags]")
	assert.Contains(t, cleanOutput, "A tool for managing projects")
	assert.Contains(t, cleanOutput, "Commands:")

	assert.Contains(t, cleanOutput, "completion  Generate the autocompletion script for the specified shell")

	assert.Contains(t, cleanOutput, "composer    Run Composer inside in a container")
	assert.Contains(t, cleanOutput, "config      Generate the default config file")
	assert.Contains(t, cleanOutput, "db          Manage the local database")
	assert.Contains(t, cleanOutput, "help        Shows the help text for an command")
	assert.Contains(t, cleanOutput, "php         Run some popular PHP tasks")
	assert.Contains(t, cleanOutput, "run         Run a Command inside the Backend Container")
	assert.Contains(t, cleanOutput, "shell       Open a shell in a container")
	assert.Contains(t, cleanOutput, "stack       Manage the local docker stac")

	assert.Contains(t, cleanOutput, "Flags:")
	assert.Contains(t, cleanOutput, "-h, --help      help for mt")
	assert.Contains(t, cleanOutput, "-v, --version   version for mt")
}

func TestCommandNotFound(t *testing.T) {
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)

	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"xyz"})

	err := rootCmd.Execute()
	assert.Error(t, err)

	assert.ErrorContains(t, err, `unknown command "xyz" for "mt"`)
}

func TestConfigHelp(t *testing.T) {
	buf := new(bytes.Buffer)

	rootCmd := cmd.RootCmd
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"config", "--help"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(buf.String())

	assert.Contains(t, cleanOutput, "Usage: mt config")
	assert.Contains(t, cleanOutput, "Generate the default config file")
	assert.Contains(t, cleanOutput, "Flags:")
	assert.Contains(t, cleanOutput, "-h, --help   help for config")
}
