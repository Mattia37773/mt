/*
Copyright © 2026 Matze
*/
package db

import (
	"bytes"
	"testing"

	"github.com/mattia37773/mt/cmd"
	base "github.com/mattia37773/mt/helper/basetest"
	"github.com/stretchr/testify/assert"
)

func TestDbHelp(t *testing.T) {
	buf := new(bytes.Buffer)

	rootCmd := cmd.RootCmd
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"db"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(buf.String())

	assert.Contains(t, cleanOutput, "  mt db [command] [flags]")
	assert.Contains(t, cleanOutput, "Manage the local database")
	assert.Contains(t, cleanOutput, "Commands")
	assert.Contains(t, cleanOutput, "export")
	assert.Contains(t, cleanOutput, "Export the local database")
	assert.Contains(t, cleanOutput, "import")
	assert.Contains(t, cleanOutput, "Import a database dump")
	assert.Contains(t, cleanOutput, "shell")
	assert.Contains(t, cleanOutput, "Connect to the database")
	assert.Contains(t, cleanOutput, "Flags:")
	assert.Contains(t, cleanOutput, "-h, --help   help for db")
}

func TestExportHelp(t *testing.T) {
	buf := new(bytes.Buffer)

	rootCmd := cmd.RootCmd
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"db", "export", "--help"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(buf.String())

	assert.Contains(t, cleanOutput, "  mt db export")
	assert.Contains(t, cleanOutput, "Export the local database")
	assert.Contains(t, cleanOutput, "Flags:")
	assert.Contains(t, cleanOutput, "-h, --help   help for export")
}

func TestImportHelp(t *testing.T) {
	buf := new(bytes.Buffer)

	rootCmd := cmd.RootCmd
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"db", "import", "--help"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(buf.String())

	assert.Contains(t, cleanOutput, "  mt db import [flags]")
	assert.Contains(t, cleanOutput, "Import a database dump")
	assert.Contains(t, cleanOutput, "Flags:")
	assert.Contains(t, cleanOutput, "-f, --file string   File to import into the Database")
	assert.Contains(t, cleanOutput, "-h, --help          help for import")
}

func TestShellHelp(t *testing.T) {
	buf := new(bytes.Buffer)

	rootCmd := cmd.RootCmd
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"db", "shell", "--help"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	cleanOutput := base.StripANSI(buf.String())

	assert.Contains(t, cleanOutput, "  mt db shell")
	assert.Contains(t, cleanOutput, "Connect to the database")
	assert.Contains(t, cleanOutput, "Flags:")
	assert.Contains(t, cleanOutput, "-h, --help   help for shell")
}
