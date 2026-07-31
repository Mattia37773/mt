/*
Copyright © 2026 Matze
*/
package stack

import (
	"bytes"
	"testing"

	"github.com/mattia37773/mt/cmd"
	_ "github.com/mattia37773/mt/cmd/single"
	base "github.com/mattia37773/mt/helper/basetest"
	"github.com/stretchr/testify/assert"
)

func TestStackHelp(t *testing.T) {
	buf := new(bytes.Buffer)

	rootCmd := cmd.RootCmd
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"stack"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(buf.String())

	assert.Contains(t, cleanOutput, "Usage: mt stack [command] [flags]")
	assert.Contains(t, cleanOutput, "Manage the local docker stack")
	assert.Contains(t, cleanOutput, "Commands:")
	assert.Contains(t, cleanOutput, "destroy  Destroy the docker stack")
	assert.Contains(t, cleanOutput, "logs     Show container logs")
	assert.Contains(t, cleanOutput, "ps       Show stack status")
	assert.Contains(t, cleanOutput, "restart  Restart the Docker Stack")
	assert.Contains(t, cleanOutput, "start    Start the docker stack")
	assert.Contains(t, cleanOutput, "stop     Stop the local stack")
	assert.Contains(t, cleanOutput, "Flags:")
	assert.Contains(t, cleanOutput, "-h, --help   help for stack")
}

func TestDestroyHelp(t *testing.T) {
	buf := new(bytes.Buffer)

	rootCmd := cmd.RootCmd
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"stack", "destroy", "--help"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(buf.String())

	assert.Contains(t, cleanOutput, "mt stack destroy")
	assert.Contains(t, cleanOutput, "Destroy the docker stack")
	assert.Contains(t, cleanOutput, "Flags:")
	assert.Contains(t, cleanOutput, "-h, --help   help for destroy")
}

func TestLogsHelp(t *testing.T) {
	buf := new(bytes.Buffer)

	rootCmd := cmd.RootCmd
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"stack", "logs", "--help"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(buf.String())

	assert.Contains(t, cleanOutput, "mt stack logs [flags] [service]")
	assert.Contains(t, cleanOutput, "Show container logs")
	assert.Contains(t, cleanOutput, "Flags:")
	assert.Contains(t, cleanOutput, "-f, --follow   Follow log output")
	assert.Contains(t, cleanOutput, "-h, --help     help for logs")
}

func TestPsHelp(t *testing.T) {
	buf := new(bytes.Buffer)

	rootCmd := cmd.RootCmd
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"stack", "ps", "--help"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(buf.String())

	assert.Contains(t, cleanOutput, "mt stack ps")
	assert.Contains(t, cleanOutput, "Show stack status")
	assert.Contains(t, cleanOutput, "Flags:")
	assert.Contains(t, cleanOutput, "-h, --help   help for ps")
}

func TestRestartHelp(t *testing.T) {
	buf := new(bytes.Buffer)

	rootCmd := cmd.RootCmd
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"stack", "restart", "--help"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(buf.String())

	assert.Contains(t, cleanOutput, "mt stack restart")
	assert.Contains(t, cleanOutput, "Restart the Docker Stack")
	assert.Contains(t, cleanOutput, "Flags:")
	assert.Contains(t, cleanOutput, "-h, --help   help for restart")
}

func TestStartHelp(t *testing.T) {
	buf := new(bytes.Buffer)

	rootCmd := cmd.RootCmd
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"stack", "start", "--help"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(buf.String())

	assert.Contains(t, cleanOutput, "mt stack start")
	assert.Contains(t, cleanOutput, "Start the docker stack")
	assert.Contains(t, cleanOutput, "Flags:")
	assert.Contains(t, cleanOutput, "-h, --help   help for start")
}

func TestStopHelp(t *testing.T) {
	buf := new(bytes.Buffer)

	rootCmd := cmd.RootCmd
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"stack", "stop", "--help"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(buf.String())

	assert.Contains(t, cleanOutput, "mt stack stop")
	assert.Contains(t, cleanOutput, "Stop the local stack")
	assert.Contains(t, cleanOutput, "Flags:")
	assert.Contains(t, cleanOutput, "-h, --help   help for stop")
}
