package single

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/mattia37773/mt/cmd"

	"github.com/mattia37773/mt/tests/base"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func captureStdout(f func()) string {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		outC <- buf.String()
	}()

	f()

	_ = w.Close()
	os.Stdout = oldStdout
	return <-outC
}

func TestHelpStyle(t *testing.T) {
	base.ChangeDirToDefault(t)

	RootCmd := cmd.RootCmd

	var dummyFlag bool
	RootCmd.Flags().BoolVarP(&dummyFlag, "help", "h", false, "help for mt")

	RootCmd.AddCommand(&cobra.Command{
		Use:   "completion",
		Short: "Generate the autocompletion script for the specified shell",
	})
	RootCmd.AddCommand(configCmd)
	RootCmd.AddCommand(&cobra.Command{
		Use:   "help",
		Short: "Shows the help text for an command",
	})

	helpStyle()

	rawOutput := captureStdout(func() {
		_ = RootCmd.Help()
	})

	cleanOutput := base.StripANSI(rawOutput)

	assert.Contains(t, cleanOutput, "Usage: mt [options] [command] [flags]")
	assert.Contains(t, cleanOutput, "A tool for managing projects")
	assert.Contains(t, cleanOutput, "Commands:")
	assert.Contains(t, cleanOutput, "completion")
	assert.Contains(t, cleanOutput, "Generate the autocompletion script for the specified shell")
	assert.Contains(t, cleanOutput, "config")
	assert.Contains(t, cleanOutput, "This command generates a config file for this project")
	assert.Contains(t, cleanOutput, "help")
	assert.Contains(t, cleanOutput, "Shows the help text for an command")
	assert.Contains(t, cleanOutput, "Flags:")
	assert.Contains(t, cleanOutput, "-h, --help")
}

func TestNotFound(t *testing.T) {
	rootCmd := cmd.RootCmd

	buf := new(bytes.Buffer)

	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"xyz"})

	err := rootCmd.Execute()
	assert.Error(t, err)

	assert.ErrorContains(t, err, `unknown command "xyz" for "mt"`)
}
