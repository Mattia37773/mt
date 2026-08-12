/*
Copyright © 2026 Matze
*/
package base

import (
	"io"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func SilentCommand(t *testing.T, cmd *cobra.Command, args []string) {

	oldStdout, oldStderr := os.Stdout, os.Stderr
	devNull, _ := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	os.Stdout, os.Stderr = devNull, devNull

	defer func() {
		os.Stdout, os.Stderr = oldStdout, oldStderr
		devNull.Close()
	}()

	cmd.SetOut(io.Discard)
	cmd.SetErr(io.Discard)

	cmd.SetArgs(args)
	err := ExecuteCommand(cmd)

	assert.NoError(t, err)
}

func ExecuteCommand(rootCmd *cobra.Command) error {
	ResetCommand(rootCmd)
	return rootCmd.Execute()
}

func ResetCommand(rootCmd *cobra.Command) {
	var visit func(c *cobra.Command)
	visit = func(c *cobra.Command) {
		c.Flags().VisitAll(func(f *pflag.Flag) {
			if f.Changed {
				_ = f.Value.Set(f.DefValue)
				f.Changed = false
			}
		})
		for _, sub := range c.Commands() {
			visit(sub)
		}
	}
	visit(rootCmd)
}
