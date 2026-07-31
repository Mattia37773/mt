package base

import (
	"io"
	"os"
	"testing"

	"github.com/spf13/cobra"
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
	err := cmd.Execute()

	assert.NoError(t, err)
}
