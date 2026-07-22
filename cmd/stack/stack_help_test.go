package stack

import (
	"bytes"
	"testing"

	"github.com/mattia37773/mt/cmd"
	_ "github.com/mattia37773/mt/cmd/single"

	"github.com/mattia37773/mt/tests/base"
	"github.com/stretchr/testify/assert"
)

func TestStackCmd(t *testing.T) {
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
	assert.Contains(t, cleanOutput, "build    Build the Docker stack")
	assert.Contains(t, cleanOutput, "destroy  Destroy the docker stack")
	assert.Contains(t, cleanOutput, "logs     Show container logs")
	assert.Contains(t, cleanOutput, "ps       Show stack status")
	assert.Contains(t, cleanOutput, "restart  Restart the Docker Stack")
	assert.Contains(t, cleanOutput, "start    Start the docker stack")
	assert.Contains(t, cleanOutput, "stop     Stop the local stack")
	assert.Contains(t, cleanOutput, "Flags:")
	assert.Contains(t, cleanOutput, "-h, --help   help for stack")
}
