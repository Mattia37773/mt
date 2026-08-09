/*
Copyright © 2026 Matze
*/
package cmd

import (
	"bytes"
	"testing"

	base "github.com/mattia37773/mt/tests/basetest"

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
	assert.Contains(t, cleanOutput, "mt version v1.0.0")
}
