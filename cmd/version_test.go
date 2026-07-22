package cmd

import (
	"bytes"
	"testing"

	"github.com/mattia37773/mt/tests/base"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	base.ChangeDirToDefault(t)

	rootCmd := RootCmd

	buf := new(bytes.Buffer)

	rootCmd.SetOut(buf)
	rootCmd.SetArgs([]string{"--version"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(buf.String())
	assert.Contains(t, cleanOutput, "mt version 1.0.0")
}
