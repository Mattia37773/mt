/*
Copyright © 2026 Matze
*/
package php

import (
	"bytes"
	"testing"

	"github.com/mattia37773/mt/cmd"

	base "github.com/mattia37773/mt/helper/basetest"
	"github.com/stretchr/testify/assert"
)

func TestPhpHelp(t *testing.T) {
	buf := new(bytes.Buffer)

	rootCmd := cmd.RootCmd
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"php"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(buf.String())

	assert.Contains(t, cleanOutput, "  mt php [command] [args]")
	assert.Contains(t, cleanOutput, "Run some popular PHP tasks")
	assert.Contains(t, cleanOutput, "Commands")
	assert.Contains(t, cleanOutput, "console")
	assert.Contains(t, cleanOutput, "Run the Symfony console")
	assert.Contains(t, cleanOutput, "craft")
	assert.Contains(t, cleanOutput, "Run the Craft CMS cli")
	assert.Contains(t, cleanOutput, "Flags:")
	assert.Contains(t, cleanOutput, "-h, --help   help for php")
}
